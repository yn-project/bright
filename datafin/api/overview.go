//nolint:nolintlint,dupl
package datafin

import (
	"context"

	"yun.tea/block/bright/datafin/pkg/mgr"

	proto "yun.tea/block/bright/proto/bright/overview"
)

func (s *OverviewServer) GetOverview(ctx context.Context, in *proto.GetOverviewRequest) (*proto.GetOverviewResponse, error) {
	return &proto.GetOverviewResponse{
		Info: mgr.GetOverviewData(),
	}, nil
}
