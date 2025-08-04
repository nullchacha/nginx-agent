package syslogprocessor

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/pdata/plog"
)

func TestProcessLogs(t *testing.T) {
	ld := plog.NewLogs()
	rl := ld.ResourceLogs().AppendEmpty()
	sl := rl.ScopeLogs().AppendEmpty()
	lr := sl.LogRecords().AppendEmpty()
	lr.Body().SetStr("<34>Oct 11 22:14:15 mymachine su: 'su root' failed for lonvick on /dev/pts/8")

	out, err := processLogs(context.Background(), ld)
	require.NoError(t, err)

	lrOut := out.ResourceLogs().At(0).ScopeLogs().At(0).LogRecords().At(0)
	attrs := lrOut.Attributes()
	val, ok := attrs.Get("syslog.hostname")
	require.True(t, ok)
	require.Equal(t, "mymachine", val.Str())
}
