package log

import (
	"github.com/hidalgo27/app-g1/cmd/server"
	"os"

	"github.com/gofiber/fiber/v2"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	logLevel := getLogLevel()

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "time"
	encoderCfg.EncodeTime = zapcore.TimeEncoderOfLayout("02-01-2006 15:04:05.000")
	encoderCfg.EncodeLevel = zapcore.LowercaseLevelEncoder

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(logLevel),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths:       []string{"stderr"},
		ErrorOutputPaths:  []string{"stderr"},
		InitialFields: map[string]interface{}{
			"pid":      os.Getpid(),
			"hostname": getHostname(),
			"system":   "api-vehicular-inspections",
		},
	}

	var err error
	log, err = config.Build()
	if err != nil {
		panic("Falló al inicializar logger: " + err.Error())
	}

	defer log.Sync()
}

func getLogLevel() zapcore.Level {
	logLevelStr := os.Getenv("LOG_LEVEL")
	if logLevelStr == "" {
		return zap.InfoLevel // por defecto
	}

	logLevel, err := zapcore.ParseLevel(logLevelStr)
	if err != nil {
		return zap.InfoLevel // por defecto
	}

	return logLevel
}

// Obtener instancia del logger
func GetLogger() *zap.Logger {
	return log
}

func getInternalLogger() *zap.Logger {
	return log.WithOptions(zap.AddCallerSkip(1))
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		panic("Falló al obtener hostname: " + err.Error())
	}
	return hostname
}

func Info(c *fiber.Ctx, message string) {
	getInternalLogger().Info(message, zap.String("transactionID", server.TransactionID(c)), zap.String("applicationID", server.ApplicationID(c)))
}

func Debug(c *fiber.Ctx, message string) {
	getInternalLogger().Debug(message, zap.String("transactionID", server.TransactionID(c)), zap.String("applicationID", server.ApplicationID(c)))
}

func Warn(c *fiber.Ctx, message string) {
	getInternalLogger().Warn(message, zap.String("transactionID", server.TransactionID(c)), zap.String("applicationID", server.ApplicationID(c)))
}

func Error(c *fiber.Ctx, message string, err error) {
	getInternalLogger().Error(message, zap.String("transactionID", server.TransactionID(c)), zap.String("applicationID", server.ApplicationID(c)), zap.Error(err))
}

func Panic(c *fiber.Ctx, message string, err error) {
	getInternalLogger().Panic(message, zap.String("transactionID", server.TransactionID(c)), zap.String("applicationID", server.ApplicationID(c)), zap.Error(err))
}

func Fatal(c *fiber.Ctx, message string, err error) {
	getInternalLogger().Fatal(message, zap.String("transactionID", server.TransactionID(c)), zap.String("applicationID", server.ApplicationID(c)), zap.Error(err))
}
