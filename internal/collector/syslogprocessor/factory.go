package syslogprocessor

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/processor/processorhelper"
)

// TypeStr is the identifier of the processor.
const TypeStr = "syslog_parser"

var processorType = component.MustNewType(TypeStr)

// Config defines configuration for the processor. No custom options for now.
type Config struct{}

// createDefaultConfig returns an empty config.
func createDefaultConfig() component.Config {
	return &Config{}
}

// NewFactory creates a factory for the syslog processor.
func NewFactory() processor.Factory {
	return processor.NewFactory(
		processorType,
		createDefaultConfig,
		processor.WithLogs(createLogsProcessor, component.StabilityLevelAlpha),
	)
}

// createLogsProcessor instantiates the logs processor.
func createLogsProcessor(ctx context.Context, set processor.Settings, cfg component.Config, next consumer.Logs) (processor.Logs, error) {
	c, ok := cfg.(*Config)
	if !ok {
		return nil, errors.New("invalid config type")
	}
	_ = c // currently unused
	return processorhelper.NewLogs(ctx, set, cfg, next, processLogs)
}
