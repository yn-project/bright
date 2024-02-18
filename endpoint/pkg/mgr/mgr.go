package mgr

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/Vigo-Tea/go-ethereum-ant/ethclient"
	"yun.tea/block/bright/common/ctredis"
	"yun.tea/block/bright/common/logger"
	"yun.tea/block/bright/endpoint/pkg/db"
)

func init() {
	err := db.Init()
	if err != nil {
		logger.Sugar().Error(err)
	}
}

type endpointIntervalMGR struct {
	RedisExpireTime time.Duration
}

type EndpointInterval struct {
	Address         string
	MinInterval     time.Duration
	BackoffIndex    int
	MaxBackoffIndex int
	MaxInterval     time.Duration
}

type EndpointList []string

var _eIMGR *endpointIntervalMGR

const (
	lockEndpointWaitTime = time.Millisecond * 100
	eIMGRPrefix          = "eIMGR"
	endpointsListKey     = "endpoint-list"
	endpointsListExpire  = time.Hour * 24
	goaheadLockTime      = time.Second * 30
)

func GetEndpintIntervalMGR() *endpointIntervalMGR {
	if _eIMGR == nil {
		_eIMGR = &endpointIntervalMGR{RedisExpireTime: endpointsListExpire}
	}
	return _eIMGR
}

func (eIMGR *endpointIntervalMGR) putEndpoint(item *EndpointInterval, autoResetBackoffIndex bool) error {
	if autoResetBackoffIndex {
		item.BackoffIndex = 0
		_maxBackoffIndex := math.Log2(float64(item.MaxInterval) / float64(item.MinInterval))
		item.MaxBackoffIndex = int(_maxBackoffIndex)
	}

	return ctredis.Set(eIMGR.getInfoKey(item.Address), item, eIMGR.RedisExpireTime)
}

func (eIMGR *endpointIntervalMGR) SetEndpoinsList(infos []string) error {
	_infos := EndpointList(infos)
	return ctredis.Set(endpointsListKey, &_infos, endpointsListExpire)
}

func (eIMGR *endpointIntervalMGR) GetEndpoinsList() ([]string, error) {
	infos := &EndpointList{}
	err := ctredis.Get(endpointsListKey, infos)
	if err != nil {
		return nil, err
	}
	return *infos, err
}

func (eIMGR *endpointIntervalMGR) GoAheadEndpoint(item *EndpointInterval) error {
	locked, err := ctredis.TryPubLock(eIMGR.getUpdateLockKey(item.Address), goaheadLockTime)
	if !locked || err != nil {
		return nil
	}

	_item := &EndpointInterval{}
	err = ctredis.Get(eIMGR.getInfoKey(item.Address), _item)
	if err != nil {
		return eIMGR.putEndpoint(item, true)
	}

	if _item.BackoffIndex > 0 {
		_item.BackoffIndex--
	}

	return eIMGR.putEndpoint(_item, false)
}

func (eIMGR *endpointIntervalMGR) BackoffEndpoint(address string) error {
	item := &EndpointInterval{}
	err := ctredis.Get(eIMGR.getInfoKey(address), item)
	if err != nil {
		return err
	}

	if item.BackoffIndex < item.MaxBackoffIndex {
		item.BackoffIndex++
	}

	return eIMGR.putEndpoint(item, false)
}

func (eIMGR *endpointIntervalMGR) getEndpointInterval(address string) (time.Duration, error) {
	item := &EndpointInterval{}
	err := ctredis.Get(eIMGR.getInfoKey(address), item)
	if err != nil {
		return 0, err
	}
	interval := item.MinInterval << item.BackoffIndex
	if interval > item.MaxInterval {
		return item.MaxInterval, nil
	}
	return interval, nil
}

func (eIMGR *endpointIntervalMGR) getInfoKey(address string) string {
	return fmt.Sprintf("%v-%v", eIMGRPrefix, address)
}

func (eIMGR *endpointIntervalMGR) getUpdateLockKey(address string) string {
	return fmt.Sprintf("%v-lock-%v", eIMGRPrefix, address)
}

func (e *EndpointInterval) MarshalBinary() (data []byte, err error) {
	data, err = json.Marshal(e)
	return data, err
}

func (e *EndpointInterval) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, e)
}

func (e *EndpointList) MarshalBinary() (data []byte, err error) {
	data, err = json.Marshal(e)
	return data, err
}

func (e *EndpointList) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, e)
}

func WithClient(ctx context.Context, handle func(ctx context.Context, cli *ethclient.Client) error) (err error) {
	eIMGR := GetEndpintIntervalMGR()
	for {
		select {
		case <-time.NewTicker(lockEndpointWaitTime).C:
			endpoints, err := eIMGR.GetEndpoinsList()
			if err != nil {
				return err
			}

			if len(endpoints) == 0 {
				return fmt.Errorf("have no available endpoints")
			}

			_randIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(endpoints))))
			if err != nil {
				return err
			}
			randIndex := int(_randIndex.Int64())
			var interval time.Duration
			endpoint := ""
			for j := 0; j < len(endpoints); j++ {
				endpoint = endpoints[(randIndex+j)%len(endpoints)]
				interval, err = eIMGR.getEndpointInterval(endpoint)
				if err != nil {
					continue
				}
				locked, err := ctredis.TryPubLock(endpoint, interval)
				if err != nil || !locked {
					continue
				}

				break
			}

			if endpoint == "" {
				continue
			}

			cli, err := ethclient.DialContext(ctx, endpoint)
			if err != nil {
				_ = eIMGR.BackoffEndpoint(endpoint)
				return err
			}
			defer cli.Close()

			err = handle(ctx, cli)
			if err != nil {
				checkErr := CheckStateAndChainID(ctx, endpoint)
				if checkErr != nil {
					_ = eIMGR.BackoffEndpoint(endpoint)
				}
			}
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
