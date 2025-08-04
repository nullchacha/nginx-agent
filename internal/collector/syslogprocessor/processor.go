package syslogprocessor

import (
	"context"
	"time"

	rfc3164 "github.com/leodido/go-syslog/v4/rfc3164"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
)

// processLogs parses syslog messages from log bodies and adds structured attributes.
func processLogs(ctx context.Context, ld plog.Logs) (plog.Logs, error) { // nolint:revive // ctx is part of interface
	parser := rfc3164.NewParser(rfc3164.WithBestEffort())

	rl := ld.ResourceLogs()
	for i := 0; i < rl.Len(); i++ {
		sl := rl.At(i).ScopeLogs()
		for j := 0; j < sl.Len(); j++ {
			lrs := sl.At(j).LogRecords()
			for k := 0; k < lrs.Len(); k++ {
				lr := lrs.At(k)
				if lr.Body().Type() != pcommon.ValueTypeStr {
					continue
				}
				line := lr.Body().Str()
				msg, err := parser.Parse([]byte(line))
				if err != nil {
					continue
				}
				m, ok := msg.(*rfc3164.SyslogMessage)
				if !ok || !m.Valid() {
					continue
				}
				attrs := lr.Attributes()
				if m.Timestamp != nil {
					attrs.PutStr("syslog.timestamp", m.Timestamp.Format(time.RFC3339))
				}
				if m.Hostname != nil {
					attrs.PutStr("syslog.hostname", *m.Hostname)
				}
				if m.Appname != nil {
					attrs.PutStr("syslog.appname", *m.Appname)
				}
				if m.ProcID != nil {
					attrs.PutStr("syslog.procid", *m.ProcID)
				}
				if sev := m.SeverityLevel(); sev != nil {
					attrs.PutStr("syslog.severity", *sev)
				}
				if fac := m.FacilityLevel(); fac != nil {
					attrs.PutStr("syslog.facility", *fac)
				}
				if m.Message != nil {
					lr.Body().SetStr(*m.Message)
				}
			}
		}
	}
	return ld, nil
}
