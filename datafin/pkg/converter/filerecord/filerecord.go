package filerecord

import (
	"yun.tea/block/bright/datafin/pkg/db/ent"
	proto "yun.tea/block/bright/proto/bright/filerecord"
)

func Ent2Grpc(row *ent.FileRecord) *proto.FileRecordInfo {
	if row == nil {
		return nil
	}

	return &proto.FileRecordInfo{
		ID:          row.ID.String(),
		PackageName: row.PackageName,
		FileName:    row.FileName,
		TopicID:     row.TopicID,
		RecordNum:   row.RecordNum,
		Sha1Sum:     row.Sha1Sum,
		State:       proto.FileRecordState(proto.FileRecordState_value[row.State]),
		Remark:      row.Remark,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.FileRecord) []*proto.FileRecordInfo {
	infos := []*proto.FileRecordInfo{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
