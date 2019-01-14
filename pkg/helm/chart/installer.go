package chart

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/softleader/captain-kube/pkg/proto"
)

type Installer interface {
	Install(log *logrus.Logger) error
}

func NewInstaller(k8s string, tiller *captainkube_v2.Tiller, chart string) (Installer, error) {
	if tiller.GetEndpoint() == "" {
		return nil, fmt.Errorf("tiller endpoint is required")
	}
	switch k8s {
	case "icp":
		return &icpInstaller{
			endpoint:          tiller.GetEndpoint(),
			chart:             chart,
			username:          tiller.GetUsername(),
			password:          tiller.GetPassword(),
			account:           tiller.GetAccount(),
			skipSslValidation: tiller.GetSkipSslValidation(),
		}, nil
	case "gcp":
		return &gcpInstaller{
			endpoint: tiller.GetEndpoint(),
			chart:    chart,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported kubernetes vendor: %v", k8s)
	}
}
