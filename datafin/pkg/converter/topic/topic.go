package topic

import (
	"yun.tea/block/bright/datafin/pkg/db/ent"
	proto "yun.tea/block/bright/proto/bright/topic"
)

func Ent2Grpc(row *ent.Topic) *proto.TopicInfo {
	if row == nil {
		return nil
	}

	return &proto.TopicInfo{
		TopicID:    row.TopicID,
		Name:       row.Name,
		Type:       proto.TopicType(proto.TopicType_value[row.Type]),
		ChangeAble: row.ChangeAble,
		Remark:     row.Remark,
		CreatedAt:  row.CreatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.Topic) []*proto.TopicInfo {
	infos := []*proto.TopicInfo{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
