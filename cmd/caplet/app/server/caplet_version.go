package server

import (
	"context"
	"github.com/softleader/captain-kube/pkg/proto"
)

func (s *CapletServer) Version(ctx context.Context, req *tw_com_softleader_captainkube.VersionRequest) (resp *tw_com_softleader_captainkube.VersionResponse, err error) {
	resp = &tw_com_softleader_captainkube.VersionResponse{
		Hostname: s.hostname,
	}
	if req.GetFull() {
		resp.Version = s.metadata.FullString()
	} else {
		resp.Version = s.metadata.String()
	}
	return
}
