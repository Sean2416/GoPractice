package rpc

import (
	"context"
	"shop-payment/shop-payment/client"
	"shop-payment/shop-payment/sutil"
	"scm.tutorabc.com/tgo-framework/go-log"
	"scm.tutorabc.com/tgo-framework/go-redis"
	"time"
)

const (
	lockKeyPrefix = "lock_"
)

type DLock struct {
	key    string
	val    string
	expire time.Duration
}

// NewDLock used for new DLock instance
func NewDLock(key string, expire time.Duration) *DLock {
	return &DLock{
		key:    lockKeyPrefix + key,
		val:    "1",
		expire: expire,
	}
}

func (d *DLock) MyKey() string {
	return d.key
}

// Lock lock the key
func (d *DLock) SetLock(ctx context.Context) (locked bool, err error) {
	lb := sutil.NewLogBuild(ctx, "SetLock")
	r := client.RedisClient(ctx)
	defer r.Close()

	defer func() {
		if lb.IsError() {
			log.Error(lb)
			return
		}

		log.Info(lb)
	}()

	var option []redis.Option
	option = append(option, redis.NX())
	option = append(option, redis.EX(d.expire))
	res, err := r.SetWithOptions(d.key, d.val, option...)
	if err != nil {
		lb.SetError(err).SetDescp("redis_error_happens")
		return
	}
	if res == false {
		lb.SetError(err).SetDescp("key_already_exists")
		return
	}

	lb.SetDescp("lock_ok")
	return true, nil

}

// UnLock unlock the key
func (d *DLock) UnLock(ctx context.Context) (err error) {
	lb := sutil.NewLogBuild(ctx, "UnLock")
	r := client.RedisClient(ctx)
	defer r.Close()
	defer func() {
		if lb.IsError() {
			log.Error(lb)
			return
		}

		log.Info(lb)
	}()

	_, err = r.Del(d.key)
	if err != nil {
		lb.SetError(err).SetDescp("unlock_failed")

		return
	}
	lb.SetDescp("unlock_succ")
	return
}

// Block the key
func (d *DLock) Block(ctx context.Context, expireAt time.Time) (locked bool, err error) {
	lb := sutil.NewLogBuild(ctx, "Block")
	r := client.RedisClient(ctx)
	defer r.Close()
	defer func() {
		if lb.IsError() {
			log.Error(lb)
			return
		}

		log.Info(lb)
	}()

	reply, err := r.ExpireAt(d.key, expireAt)
	if err != nil {
		lb.SetError(err).SetDescp("redis_error_happens")
		return
	}
	lb.SetError(err).SetDescp("lock_ok").Set("reply", reply)
	return true, nil
}
