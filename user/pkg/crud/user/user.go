package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"yun.tea/block/bright/user/pkg/db/ent/user"

	"github.com/google/uuid"
	"yun.tea/block/bright/common/cruder"
	proto "yun.tea/block/bright/proto/bright/user"
	"yun.tea/block/bright/user/pkg/db"
	"yun.tea/block/bright/user/pkg/db/ent"
)

type UserReq struct {
	ID       *string
	Name     *string
	Password *string
	Salt     *string
	Remark   *string
}

func Create(ctx context.Context, in *UserReq) (*ent.User, error) {
	var info *ent.User
	var err error
	if in == nil {
		return nil, errors.New("input is nil")
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := CreateSet(cli.User.Create(), in)
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	if info != nil {
		info.Password = ""
	}

	return info, nil
}

func CreateSet(c *ent.UserCreate, in *UserReq) *ent.UserCreate {
	if in.ID != nil {
		c.SetID(uuid.New())
	}
	if in.Name != nil {
		c.SetName(*in.Name)
	}
	if in.Password != nil {
		c.SetPassword(*in.Password)
	}
	if in.Salt != nil {
		c.SetSalt(*in.Salt)
	}
	if in.Remark != nil {
		c.SetRemark(*in.Remark)
	}
	return c
}

func Update(ctx context.Context, in *UserReq) (*ent.User, error) {
	var err error
	var info *ent.User

	if _, err := uuid.Parse(*in.ID); err != nil {
		return nil, err
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		u := cli.User.UpdateOneID(uuid.MustParse(*in.ID))
		u = UpdateSet(u, in)
		info, err = u.Save(ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	if info != nil {
		info.Password = ""
	}

	return info, nil
}

func UpdateSet(u *ent.UserUpdateOne, in *UserReq) *ent.UserUpdateOne {
	if in.Password != nil {
		u.SetPassword(*in.Password)
	}
	if in.Salt != nil {
		u.SetSalt(*in.Salt)
	}
	if in.Remark != nil {
		u.SetRemark(*in.Remark)
	}
	return u
}

func Row(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	var info *ent.User
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.User.Query().Where(user.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	if info != nil {
		info.Password = ""
	}

	return info, nil
}

func RowByName(ctx context.Context, name string) (*ent.User, error) {
	var info *ent.User
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.User.Query().Where(user.Name(name)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	// if info != nil {
	// 	info.Password = ""
	// }

	return info, nil
}

//nolint:gocyclo
func setQueryConds(conds *proto.Conds, cli *ent.Client) (*ent.UserQuery, error) {
	stm := cli.User.Query()
	if conds == nil {
		return stm, nil
	}
	if _, err := uuid.Parse(conds.GetID().GetValue()); err == nil {
		id := uuid.MustParse(conds.GetID().GetValue())
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(user.ID(id))
		default:
			return nil, fmt.Errorf("invalid user id field")
		}
	}
	if conds.IDs != nil {
		if conds.GetIDs().GetOp() == cruder.IN {
			var ids []uuid.UUID
			for _, val := range conds.GetIDs().GetValue() {
				id, err := uuid.Parse(val)
				if err != nil {
					return nil, err
				}
				ids = append(ids, id)
			}
			stm.Where(user.IDIn(ids...))
		}
	}
	if conds.Name != nil {
		switch conds.GetName().GetOp() {
		case cruder.EQ:
			stm.Where(user.Name(conds.GetName().GetValue()))
		default:
			return nil, fmt.Errorf("invalid endpoint name field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *proto.Conds, offset, limit int) ([]*ent.User, int, error) {
	var err error
	rows := []*ent.User{}
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
			Order(ent.Desc(user.FieldUpdatedAt)).
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	var info *ent.User
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.User.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	if info != nil {
		info.Password = ""
	}

	return info, nil
}
