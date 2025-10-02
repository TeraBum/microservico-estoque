package utils

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type appNameHook struct {
	appName string
}

func (h *appNameHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *appNameHook) Fire(entry *logrus.Entry) error {
	entry.Data["appName"] = h.appName
	return nil
}

type customFormatter struct {
	TimestampFormat string
}

func (f *customFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := entry.Time.Format(f.TimestampFormat)
	var level string
	switch entry.Level {
	case logrus.WarnLevel:
		level = "WARN"
	default:
		level = strings.ToUpper(entry.Level.String())
	}
	level = fmt.Sprintf("%-5s", level)

	appName, _ := entry.Data["appName"].(string)

	b := &bytes.Buffer{}
	fmt.Fprintf(b, "%s - %s : [%s] %s\n", timestamp, appName, level, entry.Message)

	return b.Bytes(), nil
}

func SetupLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&customFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.AddHook(&appNameHook{appName: "ApiEstoque"})

	return logger
}
