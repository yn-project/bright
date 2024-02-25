package filerecord

import (
	"context"
	"errors"
	"time"

	"yun.tea/block/bright/datafin/pkg/db/ent/filerecord"

	"github.com/google/uuid"
	"yun.tea/block/bright/datafin/pkg/db"
	"yun.tea/block/bright/datafin/pkg/db/ent"
	proto "yun.tea/block/bright/proto/bright/filerecord"
)

func Create(ctx context.Context, in *proto.FileRecordReq) (*ent.FileRecord, error) {
	var info *ent.FileRecord
	var err error
	if in == nil {
		return nil, errors.New("input is nil")
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := CreateSet(cli.FileRecord.Create(), in)
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateSet(c *ent.FileRecordCreate, in *proto.FileRecordReq) *ent.FileRecordCreate {
	if in.ID != nil {
		c.SetID(uuid.New())
	}
	if in.FileName != nil {
		c.SetFileName(in.GetFileName())
	}
	if in.TopicID != nil {
		c.SetTopicID(in.GetTopicID())
	}
	if in.RecordNum != nil {
		c.SetRecordNum(in.GetRecordNum())
	}
	if in.Sha1Sum != nil {
		c.SetSha1Sum(in.GetSha1Sum())
	}
	if in.State != nil {
		c.SetState(in.GetState().String())
	}
	if in.Remark != nil {
		c.SetRemark(in.GetRemark())
	}
	return c
}

func CreateBulk(ctx context.Context, in []*proto.FileRecordReq) ([]*ent.FileRecord, error) {
	var err error
	rows := []*ent.FileRecord{}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.FileRecordCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(tx.FileRecord.Create(), info)
		}
		rows, err = tx.FileRecord.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func Update(ctx context.Context, in *proto.FileRecordReq) (*ent.FileRecord, error) {
	var err error
	var info *ent.FileRecord

	if _, err := uuid.Parse(in.GetID()); err != nil {
		return nil, err
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		u := cli.FileRecord.UpdateOneID(uuid.MustParse(in.GetID()))
		u = UpdateSet(u, in)
		info, err = u.Save(ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func UpdateSet(u *ent.FileRecordUpdateOne, in *proto.FileRecordReq) *ent.FileRecordUpdateOne {
	if in.FileName != nil {
		u.SetFileName(in.GetFileName())
	}
	if in.TopicID != nil {
		u.SetTopicID(in.GetTopicID())
	}
	if in.RecordNum != nil {
		u.SetRecordNum(in.GetRecordNum())
	}
	if in.Sha1Sum != nil {
		u.SetSha1Sum(in.GetSha1Sum())
	}
	if in.State != nil {
		u.SetState(in.GetState().String())
	}
	if in.Remark != nil {
		u.SetRemark(in.GetRemark())
	}
	return u
}

func Row(ctx context.Context, id uuid.UUID) (*ent.FileRecord, error) {
	var info *ent.FileRecord
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.FileRecord.Query().Where(filerecord.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint:gocyclo
func setQueryConds(conds *proto.FileRecordConds, cli *ent.Client) (*ent.FileRecordQuery, error) {
	stm := cli.FileRecord.Query()
	if conds == nil {
		return stm, nil
	}
	if _, err := uuid.Parse(conds.GetID()); err == nil {
		id := uuid.MustParse(conds.GetID())
		stm.Where(filerecord.ID(id))
	}
	if conds.IDs != nil && len(conds.IDs) > 0 {
		ids := []uuid.UUID{}
		for _, id := range conds.IDs {
			if _id, err := uuid.Parse(id); err != nil {
				ids = append(ids, _id)
			}
		}
		stm.Where(filerecord.IDIn(ids...))
	}
	if conds.FileName != nil {
		stm.Where(filerecord.FileName(conds.GetFileName()))
	}
	if conds.TopicID != nil {
		stm.Where(filerecord.TopicID(conds.GetTopicID()))
	}
	if conds.RecordNum != nil {
		stm.Where(filerecord.RecordNum(conds.GetRecordNum()))
	}
	if conds.Sha1Sum != nil {
		stm.Where(filerecord.Sha1Sum(conds.GetSha1Sum()))
	}
	if conds.Remark != nil {
		stm.Where(filerecord.Remark(conds.GetRemark()))
	}
	if conds.State != nil {
		stm.Where(filerecord.State(conds.State.String()))
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *proto.FileRecordConds, offset, limit int) ([]*ent.FileRecord, int, error) {
	var err error
	rows := []*ent.FileRecord{}
	var total int

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}
		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}
		rows, err = stm.
			Offset(offset).
			Order(ent.Desc(filerecord.FieldUpdatedAt)).
			Limit(limit).
			All(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return rows, total, nil
}

func Delete(ctx context.Context, id uuid.UUID) (*ent.FileRecord, error) {
	var info *ent.FileRecord
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.FileRecord.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
