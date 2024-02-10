package mgr

import (
	"context"
	"time"

	"yun.tea/block/bright/common/logger"
	crud "yun.tea/block/bright/endpoint/pkg/crud/endpoint"
	"yun.tea/block/bright/endpoint/pkg/db/ent"

	"yun.tea/block/bright/proto/bright/basetype"
	proto "yun.tea/block/bright/proto/bright/endpoint"
)

const (
	RefreshTime     = time.Minute
	MaxUseEndpoints = 100
)

func Maintain(ctx context.Context) {
	for {
		select {
		case <-time.NewTicker(RefreshTime).C:
			conds := &proto.Conds{}
			infos, total, err := crud.Rows(ctx, conds, 0, MaxUseEndpoints)
			if err != nil {
				logger.Sugar().Error(err)
				continue
			}

			if total == 0 {
				continue
			}

			okEndpoints := []string{}
			for _, info := range infos {
				_, _ = CheckAndUpdateEndpoint(ctx, info)

				err = CheckStateAndChainID(ctx, info.Address)

				if err == nil {
					GetEndpintIntervalMGR().GoAheadEndpoint(&EndpointInterval{
						Address:     info.Address,
						MinInterval: time.Second / time.Duration(info.Rps),
						MaxInterval: time.Minute,
					})
					okEndpoints = append(okEndpoints, info.Address)
				}
			}

			err = GetEndpintIntervalMGR().SetEndpoinsList(okEndpoints)
			if err != nil {
				logger.Sugar().Error(err)
			}

			logger.Sugar().Infof("available endpoints: %v", okEndpoints)
		case <-ctx.Done():
			return
		}
	}
}

func CheckAndUpdateEndpoint(ctx context.Context, info *ent.Endpoint) (*ent.Endpoint, error) {
	err := CheckStateAndChainID(ctx, info.Address)
	if err != nil {
		info.State = basetype.EndpointState_EndpointError.String()
		info.Remark = err.Error()
		logger.Sugar().Warnf("endpoint:%v is not available,err: %v", info.Address, err)
	} else {
		info.State = basetype.EndpointState_EndpointAvailable.String()
		info.Remark = ""
	}

	id := info.ID.String()
	state := basetype.EndpointState(basetype.EndpointState_value[info.State])
	return crud.Update(ctx, &proto.EndpointReq{
		ID:     &id,
		State:  &state,
		Remark: &info.Remark,
	})
}
