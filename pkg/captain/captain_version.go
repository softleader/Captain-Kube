package captain

import (
	"context"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/softleader/captain-kube/pkg/dur"
	"github.com/softleader/captain-kube/pkg/proto"
	"google.golang.org/grpc"
	"io"
)

func Version(log *logrus.Logger, url string, short bool, timeout int64) error {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("did not connect: %v\n", err)
	}
	defer conn.Close()
	c := proto.NewCaptainClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), dur.Deadline(timeout))
	defer cancel()
	stream, err := c.Version(ctx, &proto.VersionRequest{
		Short:   short,
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
			fmt.Errorf("%v.Version(_) = _, %v", c, err)
		}
		log.Writer().Write(recv.GetMsg())
	}
	return nil

}
