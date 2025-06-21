package core

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/prometheus"
)

type Database struct {
	Dsn    string
	Config *gorm.Config
}

var DB *gorm.DB

type GormLogger struct {
	logger.Interface
	log *logrus.Logger
}

func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.log.SetLevel(logrus.Level(level))
	return l
}

func (l *GormLogger) Info(ctx context.Context, msg string, kv ...any) {
	l.log.WithContext(ctx).Infof("%s", msg)
}
func (l *GormLogger) Warn(ctx context.Context, msg string, kv ...any) {
	l.log.WithContext(ctx).Warnf("%s", msg)
}
func (l *GormLogger) Error(ctx context.Context, msg string, kv ...any) {
	l.log.WithContext(ctx).Errorf("%s", msg)
}
func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, row := fc()
	tid := ctx.Value(TraceId)
	l.log.WithContext(ctx).WithField("row", row).WithField("sql", sql).WithField(TraceId, tid)
}

func initDB() (err error) {
	DB, err = gorm.Open(mysql.Open(Cfg.Database.Dsn), &gorm.Config{
		Logger: &GormLogger{
			log: logrus.StandardLogger(),
		},
	})

	// 启用 promethus
	cfg := prometheus.Config{
		MetricsCollector: []prometheus.MetricsCollector{
			&prometheus.MySQL{

				Interval: 5,
			},
		},
	}
	prom := prometheus.New(cfg)
	DB.Use(prom)

	// PromethusReg.MustRegister(prom.Collectors...)
	return
}
