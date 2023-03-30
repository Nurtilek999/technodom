package cache

import (
	"fmt"
	"sync"
	"technodom/internal/repository"
	"time"
)

// Структура ссылки с TTL
type cachedUrl struct {
	activeLink        string
	expireAtTimestamp int64
}

type LocalCache struct {
	stop chan struct{}

	wg   sync.WaitGroup
	mu   sync.RWMutex
	urls map[string]cachedUrl
}

// Инициализация кэша
func NewLocalCache() repository.Cache {
	lc := &LocalCache{
		urls: make(map[string]cachedUrl),
		stop: make(chan struct{}),
	}
	lc.wg.Add(1)
	go func(cleanupInterval time.Duration) {
		defer lc.wg.Done()
		lc.CleanupLoop(cleanupInterval)
	}(time.Minute * 60)
	return lc
}

// Очистка неактивных ссылок
func (lc *LocalCache) CleanupLoop(interval time.Duration) {
	t := time.NewTicker(interval)
	defer t.Stop()

	for {
		select {
		case <-lc.stop:
			return
		case <-t.C:
			lc.mu.Lock()
			for uid, cu := range lc.urls {
				if cu.expireAtTimestamp <= time.Now().Unix() {
					delete(lc.urls, uid)
				}
			}
			lc.mu.Unlock()
		}
	}
}

func (lc *LocalCache) StopCleanup() {
	close(lc.stop)
	lc.wg.Wait()
}

func (lc *LocalCache) Add(key string, value string, expireAtTimestamp int64) {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	lc.urls[key] = cachedUrl{
		activeLink:        value,
		expireAtTimestamp: expireAtTimestamp,
	}
}

func (lc *LocalCache) Get(historyLink string) (string, bool) {
	lc.mu.RLock()
	defer lc.mu.RUnlock()

	cu, ok := lc.urls[historyLink]
	return cu.activeLink, ok
}

func (lc *LocalCache) Len() int {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	return len(lc.urls)
}

func (lc *LocalCache) Print() {
	fmt.Println(lc.urls)
}
