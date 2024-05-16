package mqueue

import (
	"yun.tea/block/bright/datafin/pkg/db/ent"
	proto "yun.tea/block/bright/proto/bright/mqueue"
)

func Ent2Grpc(row *ent.Mqueue) *proto.Mqueue {
	if row == nil {
		return nil
	}

	return &proto.Mqueue{
		ID:        row.ID.String(),
		Name:      row.Name,
		Remark:    row.Remark,
		TopicName: row.TopicName,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.Mqueue) []*proto.Mqueue {
	infos := []*proto.Mqueue{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
