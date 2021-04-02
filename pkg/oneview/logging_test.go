package oneview

import (
	"os"
	"testing"
)

const (
	defaultPath = "/tmp/oneview.log"
)

func TestLoggingWrite(t *testing.T) {
	c := createOvClient()
	events, err := GetEventList(c)
	if err != nil {
		t.Fatal(err)
	}

	l := &Logging{
		OvAddr:     "192.168.2.6",
		Path:       "/tmp/oneview_evnet.log",
		MaxSize:    50,
		MaxBackups: 3,
		MaxAge:     7,
		Compress:   true,
		Events:     events,
	}

	if err := l.Write(); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(defaultPath); os.IsNotExist(err) {
		t.Fatalf("Log file is not created: %s", defaultPath)
	}
}
