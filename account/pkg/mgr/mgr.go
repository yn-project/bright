package mgr

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"yun.tea/block/bright/common/ctredis"
)

const (
	BlockTime             = time.Second * 3
	SafeIntervalBlock     = 1
	SafeIntervalBlockTime = SafeIntervalBlock * BlockTime
	MaxAccountLockTime    = time.Minute
	MaxAccountAliveTime   = time.Hour * 24
	accountLockKey        = "lock-using-account"
	rootAccountStoreKey   = "root-account"
	treeAccountsStoreKey  = "tree-accounts"
)

type accountsMGR struct {
	RedisExpireTime time.Duration
}

var aMGR *accountsMGR

func GetAccountMGR() *accountsMGR {
	if aMGR != nil {
		return aMGR
	}
	aMGR = &accountsMGR{
		RedisExpireTime: MaxAccountLockTime,
	}
	return aMGR
}

type AccountKey struct {
	Pri string
	Pub string
}

type AccountKeyList []*AccountKey

func (aMGR *accountsMGR) SetRootAccount(address *AccountKey) error {
	return ctredis.Set(rootAccountStoreKey, address, MaxAccountAliveTime)
}

func (aMGR *accountsMGR) GetRootAccount(ctx context.Context) (acc *AccountKey, unlock func(), err error) {
	acc = &AccountKey{}
	for {
		select {
		case <-time.NewTicker(BlockTime).C:
			err = ctredis.Get(rootAccountStoreKey, acc)
			if err != nil {
				return nil, nil, fmt.Errorf("have no available tree accounts,err: %v", err)
			}

			lockKey, unlockID, err := aMGR.LockUsingAccount(acc, aMGR.RedisExpireTime)
			if err != nil {
				continue
			}

			return acc, func() {
				time.Sleep(SafeIntervalBlockTime)
				ctredis.Unlock(lockKey, unlockID)
			}, nil
		case <-ctx.Done():
			return nil, nil, fmt.Errorf("failed to get root account timeout")
		}
	}
}

func (aMGR *accountsMGR) GetRootAccountPub(ctx context.Context) (pubKey string, err error) {
	address := &AccountKey{}
	err = ctredis.Get(rootAccountStoreKey, address)
	if err != nil {
		return "", fmt.Errorf("have no available root account")
	}
	return address.Pub, nil
}

func (aMGR *accountsMGR) SetTreeAccounts(addresses []*AccountKey) error {
	accList := AccountKeyList(addresses)
	return ctredis.Set(treeAccountsStoreKey, &accList, MaxAccountAliveTime)
}

func (aMGR *accountsMGR) GetTreeAccount(ctx context.Context) (address *AccountKey, unlock func(), err error) {
	for {
		select {
		case <-time.NewTicker(BlockTime).C:
			addresses := AccountKeyList{}
			err = ctredis.Get(treeAccountsStoreKey, addresses)
			if err != nil {
				return nil, nil, fmt.Errorf("have no available tree accounts,err: %v", err)
			}

			if len(addresses) == 0 {
				return nil, nil, fmt.Errorf("have no available tree accounts")
			}

			randN := rand.Intn(len(addresses))
			lockKey := ""
			unlockID := ""
			for i := 0; i < len(addresses); i++ {
				address = addresses[(randN+i)%len(addresses)]
				lockKey, unlockID, err = aMGR.LockUsingAccount(address, aMGR.RedisExpireTime)
			}
			if err != nil {
				continue
			}
			return address, func() {
				time.Sleep(SafeIntervalBlockTime)
				ctredis.Unlock(lockKey, unlockID)
			}, nil
		case <-ctx.Done():
			return nil, nil, fmt.Errorf("failed to get tree account timeout")
		}
	}
}

func (aMGR *accountsMGR) LockUsingAccount(address *AccountKey, expire time.Duration) (string, string, error) {
	lockKey := fmt.Sprintf("%v:%v", accountLockKey, address)
	unlockID, err := ctredis.TryLock(lockKey, aMGR.RedisExpireTime)
	if err != nil {
		return "", "", err
	}
	return lockKey, unlockID, err
}

func (aMGR *accountsMGR) GetTreeAccountPub(ctx context.Context) (pubKeys []string, err error) {
	addresses := []AccountKey{}
	err = ctredis.Get(treeAccountsStoreKey, addresses)
	if err != nil {
		return nil, fmt.Errorf("have no available tree accounts,err: %v", err)
	}

	if len(addresses) == 0 {
		return nil, fmt.Errorf("have no available tree accounts")
	}

	pubKeys = []string{}
	for _, v := range addresses {
		pubKeys = append(pubKeys, v.Pub)
	}
	return pubKeys, nil
}

func (e *AccountKey) MarshalBinary() (data []byte, err error) {
	data, err = json.Marshal(e)
	return data, err
}

func (e *AccountKey) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, e)
}

func (e *AccountKeyList) MarshalBinary() (data []byte, err error) {
	data, err = json.Marshal(e)
	return data, err
}

func (e *AccountKeyList) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, e)
}
