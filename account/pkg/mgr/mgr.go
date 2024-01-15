package mgr

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"time"

	"yun.tea/block/bright/common/ctredis"
)

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

	return ctredis.Set(eIMGR.getKey(item.Address), item, eIMGR.RedisExpireTime)
}

func (eIMGR *endpointIntervalMGR) SetEndpoinsList(infos []string) error {
	return ctredis.Set(endpointsListKey, infos, endpointsListExpire)
}

func (eIMGR *endpointIntervalMGR) GetEndpoinsList() ([]string, error) {
	infos := []string{}
	err := ctredis.Get(endpointsListKey, infos)
	if err != nil {
		return nil, err
	}
	return infos, err
}

func (eIMGR *endpointIntervalMGR) GoAheadEndpoint(item *EndpointInterval) error {
	locked, err := ctredis.TryPubLock(eIMGR.getLockKey(item.Address), goaheadLockTime)
	if !locked || err != nil {
		return nil
	}

	_item := &EndpointInterval{}
	err = ctredis.Get(eIMGR.getKey(item.Address), _item)
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
	err := ctredis.Get(eIMGR.getKey(address), item)
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
	err := ctredis.Get(eIMGR.getKey(address), item)
	if err != nil {
		return 0, err
	}
	interval := item.MinInterval << item.BackoffIndex
	if interval > item.MaxInterval {
		return item.MaxInterval, nil
	}
	return interval, nil
}

func (eIMGR *endpointIntervalMGR) getKey(address string) string {
	return fmt.Sprintf("%v-%v", eIMGRPrefix, address)
}

func (eIMGR *endpointIntervalMGR) getLockKey(address string) string {
	return fmt.Sprintf("%v-lock-%v", eIMGRPrefix, address)
}

func (e *EndpointInterval) MarshalBinary() (data []byte, err error) {
	data, err = json.Marshal(e)
	return data, err
}

func (e *EndpointInterval) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, e)
}

func LockEndpoint(ctx context.Context, endpoints []string, lockTimes uint16) (endpoint string, err error) {
	for {
		select {
		case <-time.NewTicker(lockEndpointWaitTime).C:
			if len(endpoints) == 0 {
				return "", fmt.Errorf("have no available endpoints")
			}

			_randIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(endpoints))))
			if err != nil {
				return "", err
			}

			randIndex := int(_randIndex.Int64())
			var interval time.Duration
			okEndpoints := []string{}

			for j := 0; j < len(endpoints); j++ {
				endpoint := endpoints[(randIndex+j)%len(endpoints)]
				interval, err = GetEndpintIntervalMGR().getEndpointInterval(endpoint)
				if err != nil {
					continue
				}

				okEndpoints = append(okEndpoints, endpoint)

				locked, _ := ctredis.TryPubLock(endpoint, interval*time.Duration(lockTimes))
				if locked {
					return endpoint, err
				}
			}
			endpoints = okEndpoints
		case <-ctx.Done():
			return endpoint, err
		}
	}
}
