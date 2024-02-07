package account

import (
	"yun.tea/block/bright/account/pkg/db/ent"
	proto "yun.tea/block/bright/proto/bright/account"
	"yun.tea/block/bright/proto/bright/basetype"
)

func Ent2Grpc(row *ent.Account) *proto.Account {
	if row == nil {
		return nil
	}

	return &proto.Account{
		ID:      row.ID.String(),
		Address: row.Address,
		Balance: row.Balance,
		Nonce:   row.Nonce,
		State:   basetype.AccountState(basetype.AccountState_value[row.State]),
		IsRoot:  row.IsRoot,
		Remark:  row.Remark,
	}
}

func Ent2GrpcMany(rows []*ent.Account) []*proto.Account {
	infos := []*proto.Account{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
