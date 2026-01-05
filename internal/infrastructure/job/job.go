package job

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/robfig/cron/v3"
)

type Job struct {
	cron *cron.Cron
	log  *log.Helper
}

// JobOption is a Job server option.
type JobOption func(*Job)

type slog struct {
	log *log.Helper
}

// 实现cron.Logger接口所需的方法
func (s *slog) Info(msg string, keysAndValues ...interface{}) {
	s.log.Info(msg)
}

func (s *slog) Error(err error, msg string, keysAndValues ...interface{}) {
	s.log.Errorf("Cron Error: %v, msg: %s, keysAndValues: %v", err, msg, keysAndValues)
}

// NewJob creates a Job server by options.
func NewJob(opts ...JobOption) *Job {
	srv := &Job{
		log: log.NewHelper(log.DefaultLogger),
	}
	for _, o := range opts {
		o(srv)
	}

	srv.cron = cron.New(
		cron.WithSeconds(),
		cron.WithChain(cron.Recover(newLogger(srv.log))))
	return srv
}

// Logger with server logger.
func Logger(logger log.Logger) JobOption {
	return func(s *Job) {
		s.log = log.NewHelper(logger)
	}
}

func newLogger(log *log.Helper) *slog {
	return &slog{log: log}
}
