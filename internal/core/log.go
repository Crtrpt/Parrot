package core

import (
	"errors"
	"io"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

type LogCfg struct {
	Dir       string
	Level     int
	Caller    bool
	Formatter string
	Out       string
	Writer    map[string]map[string]any
}

func initLog() (err error) {
	outs := strings.Split(Cfg.Log.Out, ",")
	writers := []io.Writer{}

	for _, out := range outs {

		if writer, ok := Cfg.Log.Writer[out]; ok {
			//	输出到文件
			if out == "file" {
				dirCfg, ok := writer["dir"]
				if !ok {
					logrus.Info("need dir")
					return
				}
				dir := cast.ToString(dirCfg)
				_, err = os.Open(dir)
				if err != nil {
					if errors.Is(err, os.ErrNotExist) {
						err = os.MkdirAll(dir, os.ModePerm)
						if err != nil {
							panic(err)
						}
					} else {
						panic(err)
					}
				}
				logFile := dir + "/" + Cfg.App.Name + ".log"
				// f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
				// if err != nil {
				// 	fmt.Println("Failed to create logfile" + logFile)
				// 	panic(err)
				// }

				writer, _ := rotatelogs.New(
					logFile+".%Y%m%d",
					rotatelogs.WithLinkName(logFile),
					rotatelogs.WithRotationCount(10),
					rotatelogs.WithRotationTime(time.Hour*24),
				)

				writers = append(writers, writer)
				continue
			}
			// 输出到标准输出

			if out == "stdout" {
				writers = append(writers, os.Stdout)
			}

		}
	}
	logrus.SetOutput(io.MultiWriter(writers...))

	if Cfg.Log.Formatter == "text" {
		logrus.SetFormatter(&logrus.TextFormatter{
			TimestampFormat:  "2006-01-02 15:04:05",
			DisableTimestamp: false,
			ForceColors:      true,
		})
	}

	if Cfg.Log.Formatter == "json" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	// logrus.SetReportCaller(Cfg.Log.Caller)

	// Only log the warning severity or above.
	logrus.SetLevel(logrus.Level(Cfg.Log.Level))
	return nil
}

func LogMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()
		guid := xid.New().String()
		ctx.Set(TraceId, guid)
		ctx.Header(TraceId, guid)
		ctx.Next()
		// 结束时间
		endTime := time.Now()

		latencyTime := endTime.Sub(startTime)

		reqMethod := ctx.Request.Method

		reqUri := ctx.Request.RequestURI

		statusCode := ctx.Writer.Status()

		clientIP := ctx.ClientIP()

		logrus.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"start_time":   startTime,
			"end_time":     endTime,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
			TraceId:        guid,
		}).Info()
	}
}
