package server

import (
	"github.com/Sirupsen/logrus"
	"github.com/docker/docker/cli/command"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/softleader/captain-kube/pkg/dockerctl"
	"github.com/softleader/captain-kube/pkg/helm/chart"
	"github.com/softleader/captain-kube/pkg/proto"
	"github.com/softleader/captain-kube/pkg/sio"
	"os"
)

type capletServer struct {
	log    *logrus.Logger
	fields logrus.Fields
}

func NewCapletServer(log *logrus.Logger) (s *capletServer) {
	s = &capletServer{
		log: log,
	}
	s.fields = make(logrus.Fields)
	s.fields["hostname"], _ = os.Hostname()
	return
}

func (s *capletServer) PullImage(req *proto.PullImageRequest, stream proto.Caplet_PullImageServer) error {
	log := logrus.New()
	log.SetOutput(sio.NewStreamWriter(func(p []byte) error {
		return stream.Send(&proto.ChunkMessage{
			Msg: p,
		})
	}))
	log.WithFields(s.fields)
	if req.GetVerbose() {
		log.SetLevel(logrus.DebugLevel)
	}
	for _, image := range req.Images {
		if err := pull(log, image, req.GetRegistryAuth()); err != nil {
			return err
		}
	}
	return nil
}

func pull(log *logrus.Logger, image *proto.Image, auth *proto.RegistryAuth) error {
	if tag := image.GetTag(); len(tag) == 0 {
		image.Tag = "latest"
	}
	rc, err := dockerctl.Pull(log, chart.Image{
		Host: image.Host,
		Repo: image.Repo,
		Tag:  image.Tag,
	}, auth)
	if err != nil {
		return err
	}
	defer rc.Close()
	return jsonmessage.DisplayJSONMessagesToStream(rc, command.NewOutStream(log.Writer()), nil)
}
