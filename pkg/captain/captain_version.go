package captain

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/softleader/captain-kube/pkg/dur"
	"github.com/softleader/captain-kube/pkg/proto"
	"google.golang.org/grpc"
	"io"
)

func Version(log *logrus.Logger, url string, full, color bool, timeout int64) error {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("did not connect: %v\n", err)
	}
	defer conn.Close()
	c := tw_com_softleader_captainkube.NewCaptainClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), dur.Deadline(timeout))
	defer cancel()
	stream, err := c.Version(ctx, &tw_com_softleader_captainkube.VersionRequest{
		Full:    full,
		Color:   color,
		Timeout: timeout,
	})
	if err != nil {
		return fmt.Errorf("could not version: %v\n", err)
	}
	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("%v.Version(_) = _, %v", c, err)
		}
		log.Out.Write(recv.GetMsg())
	}
	return nil

}
