package contract

import (
	"context"
	"crypto/rand"
	"math/big"
	"os"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"yun.tea/block/bright/contract/pkg/db/ent"
	proto "yun.tea/block/bright/proto/bright/contract"
)

func init() {
	//nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
}

var (
	entContract ent.Contract
	id          string

	contractInfo proto.ContractReq
	info         *ent.Contract
)

func prepareData() {
	entContract = ent.Contract{
		ID:      uuid.New(),
		Address: "1155",
	}

	id = entContract.ID.String()
	contractInfo = proto.ContractReq{
		ID:      &id,
		Address: &entContract.Address,
	}
}

func rowToObject(row *ent.Contract) *ent.Contract {
	return &ent.Contract{
		ID:      row.ID,
		Address: row.Address,
	}
}

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &contractInfo)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			entContract.ID = info.ID
			id := info.ID.String()
			contractInfo.ID = &id
		}
		assert.Equal(t, rowToObject(info), &entContract)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entContract)
	}
}

func deleteT(t *testing.T) {
	info, err := Delete(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entContract)
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
	t.Run("row", row)
	t.Run("delete", deleteT)
}
