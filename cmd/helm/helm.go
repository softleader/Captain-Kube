package helm

import (
	"github.com/spf13/cobra"
	"os/exec"
	"fmt"
	"github.com/softleader/captain-kube/logs"
)

/** 用 helm 指令 install chart */
func Install(name string) (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "helm <Chart directory>",
		Short: "Install charts to pure Kubernetes",
		Long:  ``,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			execCmd := exec.Command("sh", "-c", "helm install "+args[0]+" -n "+name)
			stdoutStderr := log.Output(execCmd.CombinedOutput())
			fmt.Printf("Finish install  %s\n", stdoutStderr)

		},
	}
	cmd.Flags().StringVarP(&name, "name", "n", "", "helm charm nickname (required)")
	cmd.MarkFlagRequired("name")
	return
}

/** 用 helm 指令移除 chart */
func Uninstall() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "helm <helm name>",
		Short: "Uninstall charts from pure Kubernetes",
		Long:  ``,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			execCmd := exec.Command("sh", "-c", "helm delete --purge "+args[0])
			stdoutStderr := log.Output(execCmd.CombinedOutput())
			fmt.Printf("Finish uninstall  %s\n", stdoutStderr)

		},
	}
	return
}