package server

import (
	"context"
	"fmt"
	pb "github.com/softleader/captain-kube/pkg/proto"
)

// ConsoleURL 回傳 kubernetes console 的 url
func (s *CaptainServer) ConsoleURL(ctx context.Context, req *pb.ConsoleURLRequest) (*pb.ConsoleURLResponse, error) {
	resp := &pb.ConsoleURLResponse{
		Vendor: s.K8s.Server.GitVersion,
	}
	if s.K8s.Server.IsICP() {
		resp.Url = fmt.Sprintf("https://%s:%v", req.GetHost(), 8443)
		return resp, nil
	}
	if s.K8s.Server.IsGCP() {
		return nil, fmt.Errorf("GCP is not supported yet")
	}
	return nil, fmt.Errorf("unsupported kubernetes vendor: %v", s.K8s.Server.GitVersion)
}
