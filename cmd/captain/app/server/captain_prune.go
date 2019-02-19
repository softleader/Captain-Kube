package server

import (
	"github.com/sirupsen/logrus"
	"github.com/softleader/captain-kube/pkg/caplet"
	"github.com/softleader/captain-kube/pkg/dur"
	"github.com/softleader/captain-kube/pkg/proto"
	"github.com/softleader/captain-kube/pkg/sio"
	"github.com/softleader/captain-kube/pkg/utils"
)

func (s *CaptainServer) Prune(req *captainkube_v2.PruneRequest, stream captainkube_v2.Captain_PruneServer) error {
	log := logrus.New()
	log.SetOutput(sio.NewStreamWriter(func(p []byte) error {
		return stream.Send(&captainkube_v2.ChunkMessage{
			Msg: p,
		})
	}))
	log.SetFormatter(&utils.PlainFormatter{})
	if req.GetVerbose() {
		log.SetLevel(logrus.DebugLevel)
	}

	endpoints, err := s.lookupCaplet(log, req.GetColor())
	if err != nil {
		return err
	}

	log.SetNoLock()
	timeout := dur.Parse(req.GetTimeout())
	endpoints.Each(func(e *caplet.Endpoint) {
		if err := e.Prune(log, req, timeout); err != nil {
			log.Error(err)
		}
	})
	return nil
}
