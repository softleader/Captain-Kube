package ctx

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

const (
	ContextsFile   = "contexts.yaml"
	EnvMountVolume = "SL_PLUGIN_MOUNT"
)

var (
	ErrMountVolumeNotExist = errors.New(`mount volume not found
It looks like you are running the command outside slctl (https://github.com/softleader/slctl)
Please set SL_PLUGIN_MOUNT variable to manually specify the location for the command to store data 
For more details: https://github.com/softleader/slctl/wiki/Plugins-Guide#mount-volume
`)
	ErrNoActiveContextPresent = errors.New("no active context present") // 代表當前沒有 active 的 context
	PlainContexts             = new(Contexts)
	contextNameRegexp         = regexp.MustCompile(`^(.|-)$`)
	contextNameContainsRegexp = regexp.MustCompile(`(=|\s)+`)
)

type Contexts struct {
	log      *logrus.Logger
	path     string
	Contexts map[string][]string
	Active   string // 當前
	Previous string // 上一個
}

func LoadContextsFromEnv(log *logrus.Logger) (*Contexts, error) {
	mount, found := os.LookupEnv(EnvMountVolume)
	if !found {
		return nil, ErrMountVolumeNotExist
	}
	mount, err := homedir.Expand(mount)
	if err != nil {
		return nil, err
	}
	return LoadContexts(log, filepath.Join(mount, ContextsFile))
}

func LoadContexts(log *logrus.Logger, path string) (*Contexts, error) {
	log.Debugf("loading ctx from: %s\n", path)
	ctx := &Contexts{
		log:  log,
		path: path,
	}
	data, err := ioutil.ReadFile(path)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	} else if os.IsNotExist(err) {
		ctx.Contexts = make(map[string][]string)
		return ctx, nil
	}
	return ctx, json.Unmarshal(data, ctx)
}

func (c *Contexts) GetActiveExpandEnv() (*Context, error) {
	if c.Active == "" {
		return nil, ErrNoActiveContextPresent
	}
	if args, found := c.Contexts[c.Active]; !found {
		return nil, fmt.Errorf("no active context exists with name %q", c.Active)
	} else {
		ctx, err := newContext(args...)
		if err != nil {
			return nil, err
		}
		return ctx, ctx.expandEnv()
	}
}

func (c *Contexts) Add(name string, args []string, force bool) error {
	if contextNameRegexp.MatchString(name) {
		return fmt.Errorf("context name must not match regexp: %s", contextNameRegexp.String())
	}
	if contextNameContainsRegexp.MatchString(name) {
		return fmt.Errorf("context name must not match regexp: %s", contextNameContainsRegexp.String())
	}
	if _, found := c.Contexts[name]; found {
		if !force {
			return fmt.Errorf("context %q already exists", name)
		}
		delete(c.Contexts, name)
	}
	// make sure every args is fine
	if _, err := newContext(args...); err != nil {
		return err
	}
	c.Contexts[name] = args
	if err := c.save(); err != nil {
		return err
	}
	c.log.Printf("Context %q added.\n", name)
	return nil
}

func (c *Contexts) Rename(from, to string) error {
	if from == "." {
		from = c.Active
	}
	args, found := c.Contexts[from]
	if !found {
		return fmt.Errorf("no context exists with name %q", from)
	}
	if _, found := c.Contexts[to]; found {
		return fmt.Errorf("context %q already exists", to)
	}
	c.Contexts[to] = args
	delete(c.Contexts, from)
	if c.Active == from {
		c.Active = to
	}
	if c.Previous == from {
		c.Previous = to
	}
	if err := c.save(); err != nil {
		return err
	}
	c.log.Printf("Renamed context %q to %q.\n", from, to)
	return nil
}

func (c *Contexts) Delete(names ...string) (err error) {
	for _, name := range names {
		if name == "." {
			name = c.Active
		}
		if _, found := c.Contexts[name]; !found {
			return fmt.Errorf("no context exists with name %q", name)
		}
		delete(c.Contexts, name)
		if c.Active == name {
			c.Active = ""
		}
		if err = c.save(); err == nil {
			c.log.Printf("Context %q deleted.\n", name)
		}
	}
	return
}

func (c *Contexts) Switch(name string) (err error) {
	if name == "-" {
		return c.switchToPrevious()
	}
	if _, found := c.Contexts[name]; !found {
		return fmt.Errorf("no context exists with name %q", name)
	}
	c.Previous = c.Active
	c.Active = name
	if err = c.save(); err == nil {
		c.log.Printf("Switched to context %q.\n", c.Active)
	}
	return
}

func (c *Contexts) SwitchOff() (err error) {
	c.Previous = c.Active
	c.Active = ""
	if err = c.save(); err == nil {
		c.log.Print("Switched off the context.\n")
	}
	return
}

func (c *Contexts) switchToPrevious() (err error) {
	last := c.Previous
	c.Previous = c.Active
	c.Active = last
	if err = c.save(); err == nil {
		c.log.Printf("Switched to context %q.\n", c.Active)
	}
	return
}

func (c *Contexts) save() error {
	if c == PlainContexts {
		return errors.New("plain contexts is not able to save")
	}
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(c.path, data, 0644)
}
