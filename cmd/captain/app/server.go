package app

import (
	"fmt"
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"github.com/softleader/captain-kube/cmd/captain/app/server"
	"github.com/softleader/captain-kube/pkg/caplet"
	"github.com/softleader/captain-kube/pkg/captain"
	"github.com/softleader/captain-kube/pkg/env"
	"github.com/softleader/captain-kube/pkg/proto"
	"github.com/softleader/captain-kube/pkg/version"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type captainCmd struct {
	metadata       *version.BuildMetadata
	serve          string
	endpoints      []string
	port           int
	k8sVendor      string
	capletHostname string
	capletPort     int
}

func NewCaptainCommand(metadata *version.BuildMetadata) (cmd *cobra.Command) {
	var verbose bool
	c := captainCmd{
		metadata:       metadata,
		port:           env.LookupInt(captain.EnvPort, captain.DefaultPort),
		k8sVendor:      env.Lookup(captain.EnvK8sVendor, captain.DefaultK8sVendor),
		capletPort:     env.LookupInt(caplet.EnvPort, caplet.DefaultPort),
		capletHostname: env.Lookup(caplet.EnvHostname, caplet.DefaultHostname),
	}
	cmd = &cobra.Command{
		Use:  "captain",
		Long: "captain is the brain of captain-kube system",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			logrus.SetOutput(colorable.NewColorableStdout()) // for windows color output
			if verbose {
				logrus.SetLevel(logrus.DebugLevel)
			}
			return c.Run()
		},
	}

	f := cmd.Flags()
	f.BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")
	f.IntVarP(&c.port, "port", "p", c.port, "specify the port of captain, override "+captain.EnvPort)
	f.StringVar(&c.k8sVendor, "k8s-vendor", c.k8sVendor, "specify the vendor of k8s, override "+captain.EnvK8sVendor)
	f.StringVar(&c.capletHostname, "caplet-hostname", c.capletHostname, "specify the hostname of caplet daemon to lookup, override "+caplet.EnvHostname)
	f.IntVar(&c.capletPort, "caplet-port", c.capletPort, "specify the port of caplet daemon to connect, override "+caplet.EnvPort)
	f.StringArrayVarP(&c.endpoints, "caplet-endpoint", "e", []string{}, "specify the endpoint of caplet daemon to connect, override '--caplet-hostname'")

	return
}

func (c *captainCmd) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", c.port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	tw_com_softleader.RegisterCaptainServer(s, &server.CaptainServer{
		Metadata:  c.metadata,
		Hostname:  c.capletHostname,
		Endpoints: c.endpoints,
		Port:      c.capletPort,
		K8s:       c.k8sVendor,
	})
	reflection.Register(s)
	logrus.Printf("listening and serving GRPC on %v", lis.Addr().String())
	return s.Serve(lis)
}
