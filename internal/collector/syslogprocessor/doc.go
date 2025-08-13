package syslogprocessor

// Package syslogprocessor provides an OpenTelemetry processor that parses
// Syslog formatted log records and annotates them with structured attributes.
//
// # Usage
//
// Enable the processor in an OpenTelemetry Collector configuration by adding
// it to the `processors` section and referencing it from a logs pipeline:
//
//  processors:
//    syslog_parser:
//
//  service:
//    pipelines:
//      logs:
//        receivers: [tcplog]
//        processors: [syslog_parser]
//        exporters: [debug]
//
// Incoming log bodies must contain RFC3164 syslog lines. Parsed fields such as
// timestamp, hostname, app name, facility, severity, and message are written to
// the log record attributes.
