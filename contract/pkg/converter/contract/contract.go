package contract

import (
	"yun.tea/block/bright/contract/pkg/db/ent"
	proto "yun.tea/block/bright/proto/bright/contract"
)

func Ent2Grpc(row *ent.Contract) *proto.Contract {
	if row == nil {
		return nil
	}

	return &proto.Contract{
		ID:        row.ID.String(),
		Address:   row.Address,
		Name:      row.Name,
		Version:   row.Version,
		Remark:    row.Remark,
		CreatedAt: uint64(row.CreatedAt),
		UpdatedAt: uint64(row.UpdatedAt),
	}
}

func Ent2GrpcMany(rows []*ent.Contract) []*proto.Contract {
	infos := []*proto.Contract{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
