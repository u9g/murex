//go:build !windows
// +build !windows

package man

import "sync"

type summaryCacheT struct {
	mutex   sync.Mutex
	summary map[string]string
}

func NewSummaryCache() *summaryCacheT {
	sc := new(summaryCacheT)
	sc.summary = make(map[string]string)
	return sc
}

func (sc *summaryCacheT) Get(path string) string {
	sc.mutex.Lock()
	defer sc.mutex.Unlock()

	return sc.summary[path]
}

func (sc *summaryCacheT) Set(path, summary string) {
	sc.mutex.Lock()
	sc.summary[path] = summary
	sc.mutex.Unlock()
}

var SummaryCache = NewSummaryCache()
