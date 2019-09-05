package urlManager

import (
	"fmt"
	"github.com/willf/bloom"
	"sync"
)

type UrlManager interface {
	Put(url string) error
	Get() (string, error)
}

type bloomUrlManager struct {
	bloomFilter *bloom.BloomFilter
	urlQueue    []string
	mu          sync.Mutex
}

func (b *bloomUrlManager) Put(url string) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	if !b.bloomFilter.Test([]byte(url)) {
		b.bloomFilter.Add([]byte(url))
		b.urlQueue = append(b.urlQueue, url)
	}
	return nil
}

func (b *bloomUrlManager) Get() (string, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if len(b.urlQueue) == 0 {
		return "", fmt.Errorf("No Elem")
	}
	result := b.urlQueue[0]
	b.urlQueue = b.urlQueue[1:]
	return result, nil
}

func NewBloomUrlManager(m uint, k uint) *bloomUrlManager {
	return &bloomUrlManager{bloomFilter: bloom.New(m, k)}
}
