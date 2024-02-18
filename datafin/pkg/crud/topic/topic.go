package topic

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"yun.tea/block/bright/datafin/pkg/db"
	"yun.tea/block/bright/datafin/pkg/db/ent"
	"yun.tea/block/bright/datafin/pkg/db/ent/topic"
	proto "yun.tea/block/bright/proto/bright/topic"
)

func Create(ctx context.Context, in *proto.TopicReq) (*ent.Topic, error) {
	var info *ent.Topic
	var err error
	if in == nil {
		return nil, errors.New("input is nil")
	}
	db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		c := CreateSet(tx.Topic.Create(), in)
		info, err = c.Save(ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateSet(c *ent.TopicCreate, in *proto.TopicReq) *ent.TopicCreate {
	if in.ID != nil {
		c.SetID(uuid.New())
	}
	if in.TopicID != nil {
		c.SetTopicID(*in.TopicID)
	}
	if in.Name != nil {
		c.SetName(*in.Name)
	}
	if in.Type != nil {
		c.SetType(in.Type.String())
	}
	if in.ChangeAble != nil {
		c.SetChangeAble(*in.ChangeAble)
	}
	if in.Contract != nil {
		c.SetContract(*in.Contract)
	}
	if in.Remark != nil {
		c.SetRemark(in.GetRemark())
	}
	return c
}

func Row(ctx context.Context, topicID, contractAddr string) (*ent.Topic, error) {
	var info *ent.Topic
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Topic.Query().Where(topic.TopicID(topicID), topic.Contract(contractAddr)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Rows(ctx context.Context, offset, limit int, contractAddr string) ([]*ent.Topic, int, error) {
	var err error
	rows := []*ent.Topic{}
	var total int

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.Topic.Query().Where(topic.Contract(contractAddr))
		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}
		rows, err = stm.
			Offset(offset).
			Order(ent.Desc(topic.FieldUpdatedAt)).
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

func Delete(ctx context.Context, topicID string) (*ent.Topic, error) {
	var info *ent.Topic
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Topic.Query().Where(topic.TopicID(topicID)).Only(ctx)
		if err != nil {
			return err
		}

		info, err = cli.Topic.UpdateOne(info).SetDeletedAt(uint32(time.Now().Unix())).Save(ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
