package server

import (
	"github.com/sirupsen/logrus"
	"github.com/softleader/captain-kube/pkg/helm/chart"
	pb "github.com/softleader/captain-kube/pkg/proto"
	"github.com/softleader/captain-kube/pkg/sio"
	"github.com/softleader/captain-kube/pkg/utils"
)

// DeleteChart 將上傳的 chart 從 helm tiller 中刪除
func (s *CaptainServer) DeleteChart(req *pb.DeleteChartRequest, stream pb.Captain_DeleteChartServer) error {
	log := logrus.New()
	log.SetOutput(sio.NewStreamWriter(func(p []byte) error {
		return stream.Send(&pb.ChunkMessage{
			Msg: p,
		})
	}))
	log.SetFormatter(&utils.PlainFormatter{})
	if req.GetVerbose() {
		log.SetLevel(logrus.DebugLevel)
	}
	d, err := chart.NewDeleter(s.K8s, req.GetTiller(), req.GetChartName(), req.GetChartVersion())
	if err != nil {
		return err
	}
	return d.Delete(log)
}
