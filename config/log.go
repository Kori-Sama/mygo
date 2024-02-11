package config

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type MyGoFormatter struct {
	Prefix          string
	TimestampFormat string
}

func (mf *MyGoFormatter) Format(entry *log.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	b.WriteString(fmt.Sprintf("%s: [%s] [%s] - %s\n", mf.Prefix, entry.Time.Format(mf.TimestampFormat), entry.Level.String(), entry.Message))
	return b.Bytes(), nil
}

func InitLog() {
	initSystemLog()
	initGinLog()

	log.Info("Logger initialized successfully")
	fmt.Print(LOGO)
}

func initSystemLog() {
	switch Log.Level {
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	}

	if Server.Mode == "release" && log.GetLevel() > log.WarnLevel {
		log.SetLevel(log.WarnLevel)
	}

	log.SetFormatter(&MyGoFormatter{
		Prefix:          Log.Format.Prefix,
		TimestampFormat: Log.Format.Timestamp,
	})

	var logOutput []io.Writer
	if Log.Console {
		logOutput = append(logOutput, os.Stdout)
	}
	if Log.File {
		dirPath := filepath.Dir(Log.SysPath)

		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			log.Fatalf("Failed to create log directory: %v", err)
			return
		}

		f, err := os.OpenFile(Log.SysPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
			return
		}
		logOutput = append(logOutput, f)
	}

	log.SetOutput(io.MultiWriter(logOutput...))
}

func initGinLog() {
	var logOutput []io.Writer
	if Log.Console {
		logOutput = append(logOutput, os.Stdout)
	}
	if Log.File {
		dirPath := filepath.Dir(Log.GinPath)

		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			log.Fatalf("Failed to create log directory: %v", err)
			return
		}

		f, err := os.OpenFile(Log.GinPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
			return
		}
		logOutput = append(logOutput, f)
	}
	gin.DefaultWriter = io.MultiWriter(logOutput...)
}

const LOGO = `
  __  ____     _______  ____    _   _   _   _   _ 
 |  \/  \ \   / / ____|/ __ \  | | | | | | | | | |
 | \  / |\ \_/ / |  __| |  | | | | | | | | | | | |
 | |\/| | \   /| | |_ | |  | | | | | | | | | | | |
 | |  | |  | | | |__| | |__| | |_| |_| |_| |_| |_|
 |_|  |_|  |_|  \_____|\____/  (_) (_) (_) (_) (_)

`
