package main

//reference https://github.com/afex/hystrix-go.git
import (
	"sync"
	"time"
)

type Number struct {
	Buckets map[int64]*numberBucket
	Mutex   *sync.RWMutex
}

type numberBucket struct {
	Value float64
}


func NewNumber() *Number {
	r := &Number{
		Buckets: make(map[int64]*numberBucket),
		Mutex:   &sync.RWMutex{},
	}
	return r
}

func (r *Number) getCurrentBucket() *numberBucket {
	now := time.Now().Unix()
	var bucket *numberBucket
	var ok bool

	if bucket, ok = r.Buckets[now]; !ok {
		bucket = &numberBucket{}
		r.Buckets[now] = bucket
	}

	return bucket
}

func (r *Number) removeOldBuckets() {
	now := time.Now().Unix() - 10

	for timestamp := range r.Buckets {
		if timestamp <= now {
			delete(r.Buckets, timestamp)
		}
	}
}

func (r *Number) Increment(i float64) {
	if i == 0 {
		return
	}

	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	b := r.getCurrentBucket()
	b.Value += i
	r.removeOldBuckets()
}

func (r *Number) Sum(now time.Time) float64 {
	sum := float64(0)

	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	for timestamp, bucket := range r.Buckets {
		if timestamp >= now.Unix()-10 {
			sum += bucket.Value
		}
	}

	return sum
}

func main() {

	var r Number
	//起一个协程，定时滑动窗口，更新numberBucket区间
	go func() {
		for() {
			time.AfterFunc(1 * time.Second, func() {
				removeOldBuckets()
			}
		}
	}
	for() {
		
		//TODO 获取请求数量
		counter := getRequestNum()
		r.Increment(counter)

		if r.Sum() > requestLimite {
			//TODO 拒绝新的请求
		}
	}

}
