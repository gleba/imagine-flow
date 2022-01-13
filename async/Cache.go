package async

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type TimedValue[T any] struct {
	value T
	time  time.Time
}
type NSyncMap[T any] struct {
	mx       sync.RWMutex
	values   map[string]TimedValue[T]
	aliases  map[string]string
	aliveDur time.Duration
}

func (n *NSyncMap[T]) Length() int {
	n.mx.RLock()
	defer n.mx.RUnlock()
	return len(n.values)
}

func (n NSyncMap[T]) ForEach(f func(value T, key any)) {
	n.mx.RLock()
	defer n.mx.RUnlock()
	for key, v := range n.values {
		f(v.value, key)
	}
}
func (n *NSyncMap[T]) Alias(key string, filename string) {
	n.mx.Lock()
	defer n.mx.Unlock()
	n.aliases[key] = filename
}
func (n *NSyncMap[T]) Set(key string, value T) {
	n.mx.Lock()
	defer n.mx.Unlock()
	fmt.Println("add cache:", key)
	n.values[key] = TimedValue[T]{
		value: value,
		time:  time.Now(),
	}
}
func (n *NSyncMap[T]) Drop(filename any) {
	n.mx.Lock()
	defer n.mx.Unlock()
	for key, f := range n.aliases {
		if f == filename {
			fmt.Println("drop cache:", key)
			delete(n.values, key)
			delete(n.aliases, key)
		}
	}
}

func (n *NSyncMap[T]) Get(key string) (T, bool) {
	n.mx.Lock()
	defer n.mx.Unlock()
	v, ok := n.values[key]
	v.time = time.Now()
	n.values[key] = v
	fmt.Println("get cache:", key, ok)
	return v.value, ok
}
func (n *NSyncMap[T]) syncAliveTime(key any) {
}
func (n *NSyncMap[T]) clean() {
	n.mx.RLock()
	defer n.mx.RUnlock()
	now := time.Now()
	for key, v := range n.values {
		d := now.Sub(v.time)
		if d > n.aliveDur {
			fmt.Println("remove cache", key)
			delete(n.values, key)
			delete(n.aliases, key)
		}
	}
}

func getId() int {
	t := time.Now().UTC()
	return rand.New(rand.NewSource(t.UnixNano())).Int()
}

func Cache[T any](aliveTime string, refreshSecond time.Duration) *NSyncMap[T] {
	fmt.Println(time.ParseDuration(aliveTime))
	aliveDur, _ := time.ParseDuration(aliveTime)
	sMap := &NSyncMap[T]{
		mx:       sync.RWMutex{},
		values:   map[string]TimedValue[T]{},
		aliases:  map[string]string{},
		aliveDur: aliveDur,
		//id:       getId(),
	}
	var latestTicker = time.NewTicker(time.Second * refreshSecond)
	go func() {
		for {
			select {
			case _ = <-latestTicker.C:
				go sMap.clean()
			}
		}
	}()

	return sMap
}
