package app

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/softleader/captain-kube/pkg/captain"
	"github.com/softleader/captain-kube/pkg/ctx"
	"github.com/softleader/captain-kube/pkg/proto"
	"github.com/spf13/cobra"
	"io/ioutil"
	"path/filepath"
)

const (
	syncHelp = `同步 helm-chart 中的 image 到每個 worker node

	$ {{.}} sync CHART... -e CAPTAIN_ENDPOINT

亦可結合 '--retag-from' 及 '--retag-to', 同步 re-tag 之後的 image

	$ {{.}} sync CHART... -e CAPTAIN_ENDPOINT -s -f hub.softleader.com.tw -t client-registry:5000

如果 registry 需要登入, 可以傳入 '--reg-*' 開頭的 flags 指定 docker registry 的認證資訊

	$ {{.}} sync CHART... -e CAPTAIN_ENDPOINT --reg-pass SECRET
`
)

type syncCmd struct {
	hex          bool
	charts       []string
	registryAuth *ctx.RegistryAuth // docker registry auth
	retag        *ctx.ReTag
	endpoint     *ctx.Endpoint // captain 的 endpoint ip
}

func newSyncCmd() *cobra.Command {
	c := syncCmd{
		registryAuth: activeCtx.RegistryAuth,
		retag:        activeCtx.ReTag,
		endpoint:     activeCtx.Endpoint,
	}

	cmd := &cobra.Command{
		Use:   "sync CHART...",
		Short: "同步 helm-chart 中的 image 到每個 worker node",
		Long:  usage(syncHelp),
		RunE: func(cmd *cobra.Command, args []string) error {
			if c.charts = args; len(c.charts) == 0 {
				return errors.New("chart path is required")
			}
			if err := c.endpoint.Validate(); err != nil {
				return err
			}
			return c.run()
		},
	}

	f := cmd.Flags()
	f.BoolVar(&c.hex, "hex", false, "upload chart via hex string instead of bytes")
	c.registryAuth.AddFlags(f)
	c.retag.AddFlags(f)

	return cmd
}

func (c *syncCmd) run() error {
	for _, chart := range c.charts {
		if err := c.sync(chart); err != nil {
			return err
		}
	}
	return nil
}

func (c *syncCmd) sync(path string) error {
	expanded, err := homedir.Expand(path)
	if err != nil {
		expanded = path
	}
	abs, err := filepath.Abs(expanded)
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadFile(abs)
	if err != nil {
		return err
	}

	req := &captainkube_v2.SyncChartRequest{
		Color:   settings.Color,
		Timeout: settings.Timeout,
		Verbose: settings.Verbose,
		Chart: &captainkube_v2.Chart{
			FileName: filepath.Base(abs),
			FileSize: int64(len(bytes)),
		},
		RegistryAuth: &captainkube_v2.RegistryAuth{
			Username: c.registryAuth.Username,
			Password: c.registryAuth.Password,
		},
		Retag: &captainkube_v2.ReTag{
			From: c.retag.From,
			To:   c.retag.To,
		},
	}

	if c.hex {
		req.Chart.ContentHex = hex.EncodeToString(bytes)
		v, _ := json.Marshal(req)
		logrus.Debugln(string(v)) // 如果是 hex string 印出來才有意義
	} else {
		req.Chart.Content = bytes
	}

	return captain.SyncChart(logrus.StandardLogger(), c.endpoint.String(), req, settings.Timeout)
}
