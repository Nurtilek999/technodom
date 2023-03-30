package repository

import (
	"time"
)

type Cache interface {
	Add(key string, value string, expireAtTimestamp int64)
	Get(historyLink string) (string, bool)
	Len() int
	StopCleanup()
	CleanupLoop(interval time.Duration)
	Print()
}
