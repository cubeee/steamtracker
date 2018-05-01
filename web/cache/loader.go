package cache

import (
	"sync"
	"time"
)

type Loader struct {
	operations []func()
}

func (l *Loader) Add(fn func()) {
	l.operations = append(l.operations, fn)
}

func (l *Loader) LoadAsync() time.Duration {
	return l.load(true)
}

func (l *Loader) LoadSync() time.Duration {
	return l.load(false)
}

func (l *Loader) load(async bool) time.Duration {
	var wg sync.WaitGroup
	wg.Add(len(l.operations))

	start := time.Now()

	for _, op := range l.operations {
		go func(operation func()) {
			operation()
			wg.Done()
		}(op)
	}

	if !async {
		wg.Wait()
		return time.Since(start)
	}
	return 0
}
