package user

import (
	proto "yun.tea/block/bright/proto/bright/user"
	"yun.tea/block/bright/user/pkg/db/ent"
)

func Ent2Grpc(row *ent.User) *proto.User {
	if row == nil {
		return nil
	}

	return &proto.User{
		ID:     row.ID.String(),
		Name:   row.Name,
		Remark: row.Remark,
	}
}

func Ent2GrpcMany(rows []*ent.User) []*proto.User {
	infos := []*proto.User{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
