package mgr

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"yun.tea/block/bright/account/pkg/crud/txnum"
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
		case <-time.NewTimer(BlockTime).C:
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
	txnum.UpsertAddNum(context.Background(), 1)
	return address.Pub, nil
}

func (aMGR *accountsMGR) SetTreeAccounts(addresses []*AccountKey) error {
	accList := AccountKeyList(addresses)
	return ctredis.Set(treeAccountsStoreKey, &accList, MaxAccountAliveTime)
}

func (aMGR *accountsMGR) GetTreeAccount(ctx context.Context) (address *AccountKey, unlock func(), err error) {
	for {
		select {
		case <-time.NewTimer(BlockTime).C:
			_addresses := &AccountKeyList{}
			err = ctredis.Get(treeAccountsStoreKey, _addresses)
			if err != nil {
				return nil, nil, fmt.Errorf("have no available tree accounts,err: %v", err)
			}
			addresses := *_addresses
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

func (aMGR *accountsMGR) GetAccount(ctx context.Context, address *AccountKey, retries int) (unlock func(), err error) {
	for i := 0; i < retries; i++ {
		lockKey, unlockID, err := aMGR.LockUsingAccount(address, aMGR.RedisExpireTime)
		if err == nil {
			return func() {
				time.Sleep(SafeIntervalBlockTime)
				ctredis.Unlock(lockKey, unlockID)
			}, nil
		}
		select {
		case <-time.NewTimer(BlockTime).C:
			continue
		case <-ctx.Done():
			return nil, fmt.Errorf("failed to get account timeout")
		}
	}

	return nil, fmt.Errorf("failed to get account,err: %v", err)
}

func (aMGR *accountsMGR) LockUsingAccount(address *AccountKey, expire time.Duration) (string, string, error) {
	lockKey := fmt.Sprintf("%v:%v", accountLockKey, address)
	unlockID, err := ctredis.TryLock(lockKey, aMGR.RedisExpireTime)
	if err != nil {
		return "", "", err
	}
	txnum.UpsertAddNum(context.Background(), 1)
	return lockKey, unlockID, err
}

func (aMGR *accountsMGR) GetTreeAccountPub(ctx context.Context) (pubKeys []string, err error) {
	_addresses := &AccountKeyList{}
	err = ctredis.Get(treeAccountsStoreKey, _addresses)
	if err != nil {
		return nil, fmt.Errorf("have no available tree accounts,err: %v", err)
	}
	addresses := *_addresses
	if len(addresses) == 0 {
		return nil, fmt.Errorf("have no available tree accounts")
	}

	pubKeys = []string{}
	for _, v := range addresses {
		pubKeys = append(pubKeys, v.Pub)
	}
	txnum.UpsertAddNum(context.Background(), 1)
	return pubKeys, nil
}
