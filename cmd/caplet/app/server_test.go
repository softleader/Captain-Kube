package app

import (
	"github.com/sirupsen/logrus"
	"github.com/softleader/captain-kube/pkg/caplet"
	"github.com/softleader/captain-kube/pkg/proto"
	"github.com/softleader/captain-kube/pkg/utils/tcp"
	"testing"
)

func TestPullImage(t *testing.T) {
	endpoint := "localhost"
	port := caplet.DefaultPort
	if !tcp.IsReachable(endpoint, port, 3) {
		t.Skipf("endpoint %s:%v is not reachable", endpoint, port)
	}

	log := logrus.New()
	req := &proto.PullImageRequest{}
	req.Images = append(req.Images, &proto.Image{
		Host: "softleader",
		Repo: "helm",
	})

	ep := &caplet.Endpoint{
		Target: endpoint,
		Port:   port,
	}

	if err := ep.PullImage(log, req, 10); err != nil {
		t.Error(err)
	}

}
