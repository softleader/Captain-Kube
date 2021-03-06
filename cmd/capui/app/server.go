package app

import (
	"fmt"
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"github.com/softleader/captain-kube/pkg/release"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

type capUICmd struct {
	Metadata   *release.Metadata
	port       int
	StartupCtx string         // 啟動 server 的 context
	Context    *ActiveContext // 當前啟動的 context
	BaseURL    string
}

// NewCapUICommand 建立 capui root command
func NewCapUICommand(metadata *release.Metadata) (cmd *cobra.Command) {
	var verbose bool
	c := capUICmd{
		Metadata: metadata,
	}

	cmd = &cobra.Command{
		Use:  "capui",
		Long: "capui is a web interface for captain-kube",
		RunE: func(cmd *cobra.Command, args []string) error {
			logrus.SetOutput(colorable.NewColorableStdout()) // for windows color output
			if verbose {
				logrus.SetLevel(logrus.DebugLevel)
			}
			if !strings.HasSuffix(c.BaseURL, "/") {
				c.BaseURL += "/"
			}
			return c.run()
		},
	}

	f := cmd.Flags()
	f.BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")
	f.IntVarP(&c.port, "port", "p", 8080, "port of web ui serve port")
	f.StringVar(&c.StartupCtx, "active-ctx", "", "active ctx on server startup")
	f.StringVar(&c.BaseURL, "base-url", "/", "specify base url, more details: https://www.w3schools.com/tags/tag_base.asp")

	cmd.MarkFlagRequired("active-ctx")

	return
}

func (c *capUICmd) run() (err error) {
	if err = initContext(os.Environ()); err != nil {
		return
	}
	if c.Context, err = activateContext(logrus.StandardLogger(), c.StartupCtx); err != nil {
		return
	}
	server := NewCapUIServer(c)
	addr := fmt.Sprintf(":%v", c.port)
	logrus.Printf("listening and serving CapUI on %v", addr)
	return server.Run(addr)
}
