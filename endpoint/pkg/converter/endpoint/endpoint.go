package endpoint

import (
	"yun.tea/block/bright/endpoint/pkg/db/ent"
	"yun.tea/block/bright/proto/bright/basetype"
	proto "yun.tea/block/bright/proto/bright/endpoint"
)

func Ent2Grpc(row *ent.Endpoint) *proto.Endpoint {
	if row == nil {
		return nil
	}

	return &proto.Endpoint{
		ID:      row.ID.String(),
		Address: row.Address,
		Name:    row.Name,
		State:   basetype.EndpointState(basetype.EndpointState_value[row.State]),
		RPS:     row.Rps,
		Remark:  row.Remark,
	}
}

func Ent2GrpcMany(rows []*ent.Endpoint) []*proto.Endpoint {
	infos := []*proto.Endpoint{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
