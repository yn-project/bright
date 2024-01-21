package mgr

import (
	"time"
)

const (
	RefreshTime     = time.Minute
	MaxUseEndpoints = 100
)

// func Maintain(ctx context.Context) {
// 	for {
// 		select {
// 		case <-time.NewTicker(RefreshTime).C:
// 			conds := &proto.Conds{}
// 			infos, total, err := crud.Rows(ctx, conds, 0, MaxUseEndpoints)
// 			if err != nil {
// 				logger.Sugar().Error(err)
// 				continue
// 			}

// 			if total == 0 {
// 				continue
// 			}

// 			okEndpoints := []string{}
// 			for _, info := range infos {
// 				err = CheckStateAndChainID(ctx, info.Address)
// 				if err != nil {
// 					info.State = basetype.EndpointState_EndpointError.String()
// 					info.Remark = err.Error()
// 					logger.Sugar().Warnf("endpoint:%v is not available,err: %v", info.Address, err)
// 				} else {
// 					okEndpoints = append(okEndpoints, info.Address)
// 					info.State = basetype.EndpointState_EndpointAvaliable.String()
// 					info.Remark = ""
// 				}

// 				id := info.ID.String()
// 				state := basetype.EndpointState(basetype.EndpointState_value[info.State])
// 				crud.Update(ctx, &proto.EndpointReq{
// 					ID:     &id,
// 					State:  &state,
// 					Remark: &info.Remark,
// 				})
// 			}

// 			err = GetEndpintIntervalMGR().SetEndpoinsList(okEndpoints)
// 			if err != nil {
// 				logger.Sugar().Error(err)
// 			}

// 			logger.Sugar().Infof("available endpoints: %v", okEndpoints)
// 		case <-ctx.Done():
// 			return
// 		}
// 	}
// }
