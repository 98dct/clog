package clog

import "sync"

var std = New()

type logger struct {
	opt       *options
	mu        sync.Mutex
	entryPool *sync.Pool
}

func New(opts ...Option) *logger {
	logger := &logger{
		opt: initOptions(opts...),
	}
	logger.entryPool = &sync.Pool{New: func() interface{} {
		return entry(logger)
	}}
	return logger
}
