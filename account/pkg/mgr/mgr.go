package mgr

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"yun.tea/block/bright/common/ctredis"
)

const (
	BlockTime             = time.Second * 3
	SafeIntervalBlock     = 5
	SafeIntervalBlockTime = SafeIntervalBlock * BlockTime
	MaxAccountLockTime    = time.Minute
	MaxAccountAliveTime   = time.Hour * 24
	rootAccountLockKey    = "lock-root-account"
	rootAccountStoreKey   = "root-account"
	treeAccountLockKey    = "lock-tree-account"
	treeAccountStoreKey   = "tree-account"
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

func (aMGR *accountsMGR) SetSafeRootAccount(address string) error {
	return ctredis.Set(rootAccountStoreKey, address, MaxAccountAliveTime)
}

func (aMGR *accountsMGR) GetSafeRootAccount(ctx context.Context) (address string, unlock func(), err error) {
	for {
		select {
		case <-time.NewTicker(BlockTime).C:
			err = ctredis.Get(rootAccountStoreKey, address)
			if err != nil {
				return "", nil, fmt.Errorf("have no avaliable root account")
			}
			unlockID, err := ctredis.TryLock(rootAccountLockKey, aMGR.RedisExpireTime)
			if err != nil {
				continue
			}

			return address, func() {
				time.Sleep(SafeIntervalBlockTime)
				ctredis.Unlock(rootAccountLockKey, unlockID)
			}, nil
		case <-ctx.Done():
			return "", nil, fmt.Errorf("failed to get root account timeout")
		}
	}
}

func (aMGR *accountsMGR) SetTreeAccounts(addresses []string) error {
	return ctredis.Set(treeAccountStoreKey, addresses, MaxAccountAliveTime)
}

func (aMGR *accountsMGR) GetTreeAccount(ctx context.Context) (address string, unlock func(), err error) {
	for {
		select {
		case <-time.NewTicker(BlockTime).C:
			addresses := []string{}
			err = ctredis.Get(treeAccountStoreKey, addresses)
			if err != nil {
				return "", nil, fmt.Errorf("have no avaliable tree accounts")
			}

			if len(addresses) == 0 {
				return "", nil, fmt.Errorf("have no avaliable tree accounts")
			}

			randN := rand.Intn(len(addresses))
			lockKey := ""
			unlockID := ""
			for i := 0; i < len(addresses); i++ {
				address = addresses[(randN+i)%len(addresses)]
				lockKey = fmt.Sprintf("%v:%v", treeAccountLockKey, address)
				unlockID, err = ctredis.TryLock(lockKey, aMGR.RedisExpireTime)
				if err == nil {
					break
				}
			}
			if err != nil {
				continue
			}
			return address, func() {
				time.Sleep(SafeIntervalBlockTime)
				ctredis.Unlock(lockKey, unlockID)
			}, nil
		case <-ctx.Done():
			return "", nil, fmt.Errorf("failed to get tree account timeout")
		}
	}
}
