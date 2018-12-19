package app

import (
	"github.com/softleader/captain-kube/cmd/cap-ui/app/server"
	"github.com/softleader/captain-kube/pkg/verbose"
	"github.com/spf13/cobra"
	"io"
)

type capuiCmd struct {
	out        io.Writer
	configPath string
}

func NewCapuiCommand() (cmd *cobra.Command) {
	c := capuiCmd{}
	cmd = &cobra.Command{
		Use:  "capui",
		Long: "capui is a web interface for captain",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.run()
		},
	}

	c.out = cmd.OutOrStdout()
	f := cmd.Flags()
	f.BoolVarP(&verbose.Enabled, "verbose", "v", verbose.Enabled, "enable verbose output")
	f.StringVarP(&c.configPath, "config", "c", "configs/default_capui_config.yaml", "path of config file (yaml)")

	return
}

func (cmd *capuiCmd) run() error {
	c, err := server.GetConfig(cmd.configPath)
	if err != nil {
		return err
	} else {
		return server.Ui(c)
	}
}
