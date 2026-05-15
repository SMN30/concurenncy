package initonce

import (
	"sync"
)

var (
	once        sync.Once
	initialized bool
	mu          sync.RWMutex
)

// Init выполняет однократную инициализацию ресурса.
func Init() {
	// TODO: инициализировать ресурс через sync.Once
	once.Do(func() {
		mu.Lock()
		initialized = true
		mu.Unlock()
	})
}

// Initialized возвращает, был ли инициализирован ресурс.
func Initialized() bool {
	// TODO: вернуть признак инициализации
	mu.RLock()
	defer mu.RUnlock()
	return initialized
}
