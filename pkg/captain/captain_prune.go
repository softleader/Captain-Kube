package captain

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	pb "github.com/softleader/captain-kube/pkg/proto"
	"google.golang.org/grpc"
	"io"
	"time"
)

// CallPrune 呼叫 captain Prune gRPC API
func CallPrune(log *logrus.Logger, url string, verbose, color bool, timeout time.Duration) error {
	log.Debugf("dialing %q with insecure", url)
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCaptainClient(conn)
	log.Debugf("setting context with timeout %v", timeout)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req := &pb.PruneRequest{
		Verbose: verbose,
		Timeout: timeout.String(),
		Color:   color,
	}
	stream, err := c.Prune(ctx, req)
	if err != nil {
		return fmt.Errorf("%v.Prune(%v) = _, %v", c, req, err)
	}
	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to receive a chunk msg : %v", err)
		}
		log.Out.Write(recv.GetMsg())
	}
	return nil

}
