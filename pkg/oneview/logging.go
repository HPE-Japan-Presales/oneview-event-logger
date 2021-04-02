package oneview

import (
	"log"

	"gopkg.in/natefinch/lumberjack.v2"
)

type Logging struct {
	OvAddr     string
	Path       string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
	Events     EventList
}

func (l *Logging) Write() error {
	log.SetOutput(&lumberjack.Logger{
		Filename:   l.Path,       // log file path
		MaxSize:    l.MaxSize,    // megabytes
		MaxBackups: l.MaxBackups, // Maximum log file generations
		MaxAge:     l.MaxAge,     // how long keep old log?
		LocalTime:  true,
		Compress:   l.Compress, // gzip
	})
	//log.SetFlags(0)
	for _, event := range l.Events.Members {
		log.Printf("OneView:%s Created:%s Severity:%s Category:%s Desc:\"%s\"\n", l.OvAddr, event.Created, event.Severity, event.HealthCategory, event.Description)
	}

	return nil
}
