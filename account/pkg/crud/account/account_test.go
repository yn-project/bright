package account

import (
	"context"
	"crypto/rand"
	"math/big"
	"os"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"yun.tea/block/bright/account/pkg/db/ent"
	"yun.tea/block/bright/common/cruder"
	val "yun.tea/block/bright/proto/bright"
	proto "yun.tea/block/bright/proto/bright/account"
)

func init() {
	//nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
}

var (
	entAccount ent.Account
	id         string

	accountInfo proto.AccountReq
	info        *ent.Account
)

func prepareData() {
	entAccount = ent.Account{
		ID:      uuid.New(),
		Address: "1155",
	}

	id = entAccount.ID.String()
	accountInfo = proto.AccountReq{
		ID:      &id,
		Address: &entAccount.Address,
	}
}

func rowToObject(row *ent.Account) *ent.Account {
	return &ent.Account{
		ID:      row.ID,
		Address: row.Address,
	}
}

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &accountInfo)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			entAccount.ID = info.ID
			id := info.ID.String()
			accountInfo.ID = &id
		}
		assert.Equal(t, rowToObject(info), &entAccount)
	}
}

func createBulk(t *testing.T) {
	entAccount := []ent.Account{
		{
			ID:      uuid.New(),
			Address: "1155",
		},
		{
			ID:      uuid.New(),
			Address: "1155",
		},
	}

	accounts := []*proto.AccountReq{}
	for key := range entAccount {
		id := entAccount[key].ID.String()
		accounts = append(accounts, &proto.AccountReq{
			ID:      &id,
			Address: &entAccount[key].Address,
		})
	}
	infos, err := CreateBulk(context.Background(), accounts)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
		assert.NotEqual(t, infos[0].ID, uuid.UUID{}.String())
		assert.NotEqual(t, infos[1].ID, uuid.UUID{}.String())
	}
}

func update(t *testing.T) {
	var err error
	info, err = Update(context.Background(), &accountInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entAccount)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entAccount)
	}
}

func rows(t *testing.T) {
	infos, total, err := Rows(context.Background(),
		&proto.Conds{
			ID: &val.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		}, 0, 0)
	if assert.Nil(t, err) {
		assert.Equal(t, total, 1)
		assert.Equal(t, rowToObject(infos[0]), &entAccount)
	}
}

func deleteT(t *testing.T) {
	info, err := Delete(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entAccount)
	}
}

func RandInt64() int64 {
	MaxUint64 := ^uint64(0)
	MaxInt64 := int64(MaxUint64 >> 1)
	randInt, err := rand.Int(rand.Reader, big.NewInt(MaxInt64))
	if err != nil {
		return 0
	}
	return randInt.Int64()
}

func RandInt() int {
	return int(RandInt64())
}

func RandInt32() int32 {
	return int32(RandInt64())
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	prepareData()
	t.Run("create", create)
	t.Run("createBulk", createBulk)
	t.Run("row", row)
	t.Run("rows", rows)
	t.Run("update", update)
	t.Run("delete", deleteT)
}
