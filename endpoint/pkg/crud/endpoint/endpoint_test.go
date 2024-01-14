package endpoint

import (
	"context"
	"crypto/rand"
	"math/big"
	"os"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"yun.tea/block/bright/common/cruder"
	"yun.tea/block/bright/endpoint/pkg/db/ent"
	val "yun.tea/block/bright/proto/bright"
	proto "yun.tea/block/bright/proto/bright/endpoint"
)

func init() {
	//nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
}

var (
	entEndpoint ent.Endpoint
	id          string

	endpointInfo proto.EndpointReq
	info         *ent.Endpoint
)

func prepareData() {
	entEndpoint = ent.Endpoint{
		ID:      uuid.New(),
		Address: "1155",
	}

	id = entEndpoint.ID.String()
	endpointInfo = proto.EndpointReq{
		ID:      &id,
		Address: &entEndpoint.Address,
	}
}

func rowToObject(row *ent.Endpoint) *ent.Endpoint {
	return &ent.Endpoint{
		ID:      row.ID,
		Address: row.Address,
	}
}

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &endpointInfo)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			entEndpoint.ID = info.ID
			id := info.ID.String()
			endpointInfo.ID = &id
		}
		assert.Equal(t, rowToObject(info), &entEndpoint)
	}
}

func createBulk(t *testing.T) {
	entEndpoint := []ent.Endpoint{
		{
			ID:      uuid.New(),
			Address: "1155",
		},
		{
			ID:      uuid.New(),
			Address: "1155",
		},
	}

	endpoints := []*proto.EndpointReq{}
	for key := range entEndpoint {
		id := entEndpoint[key].ID.String()
		endpoints = append(endpoints, &proto.EndpointReq{
			ID:      &id,
			Address: &entEndpoint[key].Address,
		})
	}
	infos, err := CreateBulk(context.Background(), endpoints)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
		assert.NotEqual(t, infos[0].ID, uuid.UUID{}.String())
		assert.NotEqual(t, infos[1].ID, uuid.UUID{}.String())
	}
}

func update(t *testing.T) {
	var err error
	info, err = Update(context.Background(), &endpointInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entEndpoint)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entEndpoint)
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
		assert.Equal(t, rowToObject(infos[0]), &entEndpoint)
	}
}

func deleteT(t *testing.T) {
	info, err := Delete(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entEndpoint)
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
