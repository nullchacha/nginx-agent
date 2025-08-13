# Syslog Parser Processor

The `syslog_parser` processor converts RFC3164 Syslog messages into structured
attributes on a log record. It expects the log body to contain a raw Syslog
line. When parsing succeeds, fields such as timestamp, hostname, app name,
facility, and severity are added as `syslog.*` attributes while the body is
replaced with the message contents.

## Configuration

Add the processor to the `processors` section of your OpenTelemetry Collector
configuration and reference it from a logs pipeline. The example below uses the
built-in `tcplog` receiver and `debug` exporter:

```yaml
receivers:
  tcplog:
    listen_address: "0.0.0.0:514"

processors:
  syslog_parser: {}

exporters:
  debug: {}

service:
  pipelines:
    logs:
      receivers: [tcplog]
      processors: [syslog_parser]
      exporters: [debug]
```

## Running

1. Build and run the agent or collector with a configuration that points to the
   above Collector YAML. For example:
   ```bash
   go run ./cmd/agent -c /path/to/nginx-agent.conf
   ```
2. Send a Syslog message to the receiver:
   ```bash
   logger -n localhost -P 514 "<13>Oct 11 22:14:15 host app[123]: hello"
   ```
3. Observe the structured attributes through the configured exporter (the
   `debug` exporter prints them to stdout).
