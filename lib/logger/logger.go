package logger

import (
	"fmt"
	logpkg "github.com/ckeyer/go-log"
	"io"
	"os"
	"strings"
)

var (
	LOG_DIR  = "./tmp/test/"
	LOG_NAME = "test"
)

var log *logpkg.Logger

// SetLogger 设置及检查日志路径
func SetLogger(name, dir string) error {
	stat, err := os.Stat(dir)
	if err != nil {
		if strings.HasSuffix(err.Error(), "no such file or directory") {
			err := os.MkdirAll(dir, 0755)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	} else if !stat.IsDir() {
		return fmt.Errorf("%s is not a directory", dir)
	}
	LOG_DIR = dir
	LOG_NAME = name
	return nil
}

func GetLogger() *logpkg.Logger {
	if log != nil {
		return log
	}
	format := logpkg.MustStringFormatter(
		"%{time:15:04:05} [%{color}%{level:.4s}%{color:reset}] %{shortfile} %{color}▶%{color:reset} %{message}")
	backend := logpkg.NewLogBackend(os.Stdout, "", 2)
	backend1Leveled := logpkg.AddModuleLevel(backend)
	backend1Leveled.SetLevel(logpkg.DEBUG, "")
	backendFormatter := logpkg.NewBackendFormatter(backend, format)

	err_format := logpkg.MustStringFormatter("%{time:15:04:05} [%{level:.4s}] %{shortfile} : %{message}")
	err_writer, err := getLogFile("err")
	if err != nil {
		fmt.Println("get log file err, ", err)
	}
	err_backend := logpkg.NewLogBackend(err_writer, "", 0)
	// err_backend1Leveled := logpkg.AddModuleLevel(err_backend)
	// err_backend1Leveled.SetLevel(logpkg.ERROR, "")

	_ = err_format
	err_backendFormatter := logpkg.NewBackendFormatter(err_backend, err_format)

	logpkg.SetBackend(backendFormatter, err_backendFormatter)
	// logpkg.SetBackend(err_backendFormatter)

	log = logpkg.MustGetLogger("/")

	return log
}

func getLogFile(level string) (io.Writer, error) {
	name := LOG_DIR + "/" + LOG_NAME + "_" + level + ".log"
	stat, err := os.Stat(name)
	if err != nil {
		if strings.HasSuffix(err.Error(), "no such file or directory") {
			f, err := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0600)
			if err != nil {
				return os.Stdout, err
			}
			return f, nil
		}
	}
	if stat.IsDir() {
		return os.Stdout, fmt.Errorf("%s is a directory", name)
	}

	f, err := os.OpenFile(name, os.O_APPEND|os.O_RDWR, 0600)
	if err != nil {
		return os.Stdout, err
	}
	return f, nil
}
