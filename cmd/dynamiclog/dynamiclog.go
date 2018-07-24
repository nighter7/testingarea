package main

import (
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {

	atom := zap.NewAtomicLevel()

	// To keep the example deterministic, disable timestamps in the output.
	encoderCfg := zap.NewProductionEncoderConfig()

	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		atom,
	))
	defer logger.Sync()

	viper.SetConfigName("agent-config")
	viper.AddConfigPath("./../../config")

	if err := viper.ReadInConfig(); err != nil {
		logger.Panic("Read file not read")
	}

	ll := viper.Get("logger.level")

	var level zapcore.Level

	switch ll {
	case "error":
		level = zapcore.ErrorLevel
	case "debug":
		level = zapcore.DebugLevel
	case "fatal":
		level = zapcore.FatalLevel
	default:
		level = zapcore.InfoLevel
	}

	atom.SetLevel(level)
	logger.Info("logging")
}
