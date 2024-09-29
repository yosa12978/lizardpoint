package router

import (
	"os"

	"github.com/yosa12978/lizardpoint/internal/logging"
)

type options struct {
	logger logging.Logger
}

type optionFunc func(opt *options)

func newOptions(opts ...optionFunc) options {
	o := newDefaultOptions()
	for _, opt := range opts {
		opt(&o)
	}
	return o
}

func newDefaultOptions() options {
	return options{
		logger: logging.NewJsonLogger(os.Stdout, "INFO"),
	}
}

func WithLogger(logger logging.Logger) optionFunc {
	return func(opt *options) {
		opt.logger = logger
	}
}
