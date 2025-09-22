package config

import (
	"log"
	"os"
	"strconv"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// setupLogger настраивает логгер с ротацией
func setupLogger() *zap.Logger {
	logsDir := "logs"
	logPath := getEnv("LOG_PATH", "logs/app.log")

	if err := os.MkdirAll(logsDir, 0755); err != nil {
		log.Printf("Failed to create logs directory: %v", err)
	}

	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		if file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY, 0644); err == nil {
			file.Close()
		}
	} else {
		os.Chmod(logPath, 0644)
	}

	// Настройка ротации логов
	logWriter := &lumberjack.Logger{
		Filename:   getEnv("LOG_PATH", "logs/app.log"),
		MaxSize:    getEnvInt("LOG_MAX_SIZE", 20),    // megabytes
		MaxBackups: getEnvInt("LOG_MAX_BACKUPS", 10), // количество файлов
		MaxAge:     getEnvInt("LOG_MAX_AGE", 30),     // days
		Compress:   getEnvBool("LOG_COMPRESS", true), // gzip
	}

	// Уровень логирования
	level := parseLogLevel(getEnv("LOG_LEVEL", "info"))

	// Конфиг энкодера для JSON
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// Создаем cores
	cores := []zapcore.Core{}

	// File core (основной)
	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(logWriter),
		level,
	)
	cores = append(cores, fileCore)

	// Console core (только для development или если включено)
	if getEnvBool("LOG_CONSOLE_OUTPUT", true) || os.Getenv("APP_ENV") == "development" {
		consoleEncoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		})

		consoleCore := zapcore.NewCore(
			consoleEncoder,
			zapcore.AddSync(os.Stdout),
			level,
		)
		cores = append(cores, consoleCore)
	}

	core := zapcore.NewTee(cores...)

	logger := zap.New(core,
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel), // stacktrace только для ошибок
	)

	return logger
}

// parseLogLevel парсит уровень логирования
func parseLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}

// Простые вспомогательные функции
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}
