package datafin

import (
	"yun.tea/block/bright/datafin/pkg/db/ent"
	proto "yun.tea/block/bright/proto/bright/datafin"
)

func Ent2Grpc(row *ent.DataFin) *proto.DataFinInfo {
	if row == nil {
		return nil
	}

	return &proto.DataFinInfo{
		DataID:      row.DataID,
		DataFinID:   row.ID.String(),
		TopicID:     row.TopicID,
		DataFin:     row.Datafin,
		TxTime:      row.TxTime,
		TxHash:      row.TxHash,
		BlockHeight: row.BlockHeight,
		State:       proto.DataFinState(proto.DataFinState_value[row.State]),
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.DataFin) []*proto.DataFinInfo {
	infos := []*proto.DataFinInfo{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
