package logger

import (
	"io"
	"os"
	"path/filepath"

	"github.com/alecthomas/units"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/xhit/go-str2duration/v2"
)

type Logconfig struct {
	Path, Pattern, MaxAge, RotationTime, RotationSize, Mode string
}

type logBundle struct {
	logger *logrus.Logger
}

func New(lc *Logconfig) (Logger, error) {
	l := &logBundle{logger: logrus.New()}
	if lc.Mode == "debug" {
		l.logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
		l.logger.SetOutput(os.Stdout)
	}
	if lc.Mode == "product" {
		writer, err := LoggerWriter(lc)
		if err != nil {
			return nil, err
		}
		l.logger.SetOutput(writer)
		l.logger.SetFormatter(&logrus.JSONFormatter{})
	}
	return l, nil
}

func LoggerWriter(lc *Logconfig) (io.Writer, error) {
	maxAge, err := str2duration.ParseDuration(lc.MaxAge)
	if err != nil {
		return nil, err
	}

	rotationTime, err := str2duration.ParseDuration(lc.RotationTime)
	if err != nil {
		return nil, err
	}

	rotationSize, err := units.ParseBase2Bytes(lc.RotationSize)
	if err != nil {
		return nil, err
	}

	return rotatelogs.New(
		filepath.Join(lc.Path, lc.Pattern),
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
		rotatelogs.WithRotationSize(int64(rotationSize)),
	)
}

// Info is log with level Info
func (l *logBundle) Info(msg string, extras ...map[string]interface{}) {
	l.logger.WithFields(logrus.Fields{
		"extras": extras,
	}).Info(msg)
}

// Warning is log with level warning
func (l *logBundle) Warning(msg string, extras ...map[string]interface{}) {
	l.logger.WithFields(logrus.Fields{
		"extras": extras,
	}).Warning(msg)
}

// Error is log with level error
func (l *logBundle) Error(msg string, extras ...map[string]interface{}) {
	l.logger.WithFields(logrus.Fields{
		"extras": extras,
	}).Error(msg)
}

// Panic is log with level panic
func (l *logBundle) Panic(msg string, extras ...map[string]interface{}) {
	l.logger.WithFields(logrus.Fields{
		"extras": extras,
	}).Panic(msg)
}
