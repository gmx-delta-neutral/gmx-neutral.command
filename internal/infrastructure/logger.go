package infrastructure

import (
	log "github.com/sirupsen/logrus"
	"github.com/uptrace/opentelemetry-go-extra/otellogrus"
)

func NewLogger(serviceName string) *log.Entry {
	log.SetFormatter(&log.JSONFormatter{})

	logger := log.WithFields(log.Fields{
		"service": serviceName,
	})

	log.AddHook(otellogrus.NewHook(otellogrus.WithLevels(
		log.AllLevels...,
	)))

	return logger
}
