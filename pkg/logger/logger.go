package logger

import (
	"os"

	"github.com/op/go-logging"
)

func Logger(appName string) *logging.Logger {
	var logger logging.LeveledBackend
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	formatter := logging.MustStringFormatter(
		`%{shortfile} %{color:bold} ▶ [%{module}] [%{level:.6s}] %{message}%{color:reset}`,
	)

	backendFormatter := logging.NewBackendFormatter(backend, formatter)
	logger = logging.SetBackend(backendFormatter)

	logger.SetLevel(logging.DEBUG, "")
	return logging.MustGetLogger(appName)
}
