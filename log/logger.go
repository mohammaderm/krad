package log

import (
	"io"
	"path/filepath"

	"github.com/alecthomas/units"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/xhit/go-str2duration/v2"
)

type Logconfig struct {
	Path, Pattern, MaxAge, RotationTime, RotationSize string
}

type logBundle struct {
	logger *logrus.Logger
}

func New(lc *Logconfig) (Logger, error) {
	l := &logBundle{logger: logrus.New()}
	writer, err := LoggerWriter(lc)
	if err != nil {
		return nil, err
	}
	l.logger.SetOutput(writer)
	l.logger.SetFormatter(&logrus.JSONFormatter{})
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

//Info is log with level Info
func (l *logBundle) Info(field *Field) {
	l.logger.WithFields(logrus.Fields{
		"package":  field.Package,
		"function": field.Function,
		"params":   field.Params,
	}).Info(field.Message)
}

//Warning is log with level warning
func (l *logBundle) Warning(field *Field) {
	l.logger.WithFields(logrus.Fields{
		"package":  field.Package,
		"function": field.Function,
		"params":   field.Params,
	}).Warning(field.Message)
}

//Error is log with level error
func (l *logBundle) Error(field *Field) {
	l.logger.WithFields(logrus.Fields{
		"package":  field.Package,
		"function": field.Function,
		"params":   field.Params,
	}).Error(field.Message)
}

//Panic is log with level panic
func (l *logBundle) Panic(field *Field) {
	l.logger.WithFields(logrus.Fields{
		"package":  field.Package,
		"function": field.Function,
		"params":   field.Params,
	}).Panic(field.Message)
}
