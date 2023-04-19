package config

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitLogger() {
	// Initialize the logrus logger
	logger := logrus.New()

	// ----------------------------------------
	// Set logrus logger level
	// ----------------------------------------
	level, err := logrus.ParseLevel(strings.ToUpper(viper.GetString("LOG_LEVEL")))
	if err != nil {
		level = logrus.InfoLevel // Default to "info" if the level is invalid
	}
	logger.SetLevel(level)

	// ----------------------------------------
	// Configure the logrus logger formatter
	// ----------------------------------------
	// &logrus.TextFormatter{}: A formatter that outputs log messages in a human-readable text format.
	// &logrus.JSONFormatter{}: A formatter that outputs log messages in JSON format.
	// &logrus.LogstashFormatter{}: A formatter that outputs log messages in the Logstash JSON format.
	// &logrus.HumanReadableFormatter{}: A formatter that outputs log messages in a more readable text format than the TextFormatter.
	// &logrus.FullTimestampFormatter{}: A formatter that includes the full timestamp in the log message.
	// &logrus.TimestampFormatter{}: A formatter that includes a customizable timestamp in the log message.

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(level)
}
