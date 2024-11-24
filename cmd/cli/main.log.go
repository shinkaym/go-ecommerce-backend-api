package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	//1.
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("hello name:%s, age:%d, ", "Khoa", 22)

	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "Khoa dep trai"), zap.Int("age", 22))

	//2.
	// logger := zap.NewExample()
	// logger.Info("Hello NewExample")
	//DEV
	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello NewDevelopment")
	//PROD
	// logger, _ = zap.NewProduction()
	// logger.Info("Hello NewProduction")

	//3.
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))
}

// format log
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	// :1732181582.9600563 -> 2024-11-21T16:33:02.959+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// ts -> Time
	encodeConfig.TimeKey = "time"
	// from info INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// "caller":"cli/main.log.go:22"
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}

func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
