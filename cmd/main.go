package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/fideltak/oneview-event-logger/pkg/oneview"
)

const (
	defaultOvVer             = 1200
	defaultOvSslVerification = false
	defaultLogPath           = "/tmp/oneview_evnet.log"
	defaultMaxLogSize        = 50 //MB
	defaultMaxBackups        = 5
	defaultMaxAge            = 365 //Day
	defaultCompress          = true
)

var (
	version    = "test"
	ovAddr     string
	ovUser     string
	ovPassword string
	ovVer      int
	path       string
	size       int
	backups    int
	age        int
	compress   bool
)

func init() {
	//Get Envronment Values
	ovAddr = os.Getenv("OV_ADDR")
	if ovAddr == "" {
		fmt.Printf("Error: Not set OneView address or hostname\n")
		os.Exit(1)
	}

	ovUser = os.Getenv("OV_USER")
	if ovUser == "" {
		fmt.Printf("Error: Not set OneView user\n")
		os.Exit(1)
	}

	ovPassword = os.Getenv("OV_PASSWORD")
	if ovPassword == "" {
		fmt.Printf("Error: Not set OneView password\n")
		os.Exit(1)
	}

	ovVerStr := os.Getenv("OV_VERSION")
	if ovVerStr == "" {
		ovVer = defaultOvVer
	} else {
		v, err := strconv.Atoi(ovVerStr)
		if err != nil {
			fmt.Printf("Error: OneView API version parse error: %v\n", err)
			os.Exit(1)
		}
		ovVer = v
	}

	path = os.Getenv("OV_LOG_PATH")
	if path == "" {
		path = defaultLogPath
	}

	sizeStr := os.Getenv("OV_LOG_MAX_SIZE_MB")
	if sizeStr == "" {
		size = defaultMaxLogSize
	} else {
		s, err := strconv.Atoi(sizeStr)
		if err != nil {
			fmt.Printf("Error: Log max size parse error: %v\n", err)
			os.Exit(1)
		}
		size = s
	}

	backupsStr := os.Getenv("OV_LOG_MAX_BACKUPS")
	if backupsStr == "" {
		backups = defaultMaxBackups
	} else {
		b, err := strconv.Atoi(backupsStr)
		if err != nil {
			fmt.Printf("Error: Log max backups parse error: %v\n", err)
			os.Exit(1)
		}
		backups = b
	}

	ageStr := os.Getenv("OV_LOG_MAX_AGE")
	if ageStr == "" {
		age = defaultMaxAge
	} else {
		a, err := strconv.Atoi(ageStr)
		if err != nil {
			fmt.Printf("Error: Log max age parse error: %v\n", err)
			os.Exit(1)
		}
		age = a
	}

	compressStr := os.Getenv("OV_LOG_COMPRESS")
	if compressStr == "" {
		compress = defaultCompress
	} else {
		c, err := strconv.ParseBool(compressStr)
		if err != nil {
			fmt.Printf("Error: Log compress parse error: %v\n", err)
			os.Exit(1)
		}
		compress = c
	}

}

func main() {
	fmt.Printf("HPE OneView Event Logger Version %s\n", version)
	fmt.Printf("OneView Address: %s\n", ovAddr)
	fmt.Printf("OneView User: %s\n", ovUser)
	fmt.Printf("OneView API Version: %v\n", ovVer)
	fmt.Printf("OneView Event Log Path: %s\n", path)
	fmt.Printf("Log Max Size: %v\n", size)
	fmt.Printf("Log Buckup Days: %v\n", backups)
	fmt.Printf("Log Buckup Ages: %v\n", age)
	fmt.Printf("Log Compression: %v\n", compress)

	ovUrl := "https://" + ovAddr
	var ovClient *ov.OVClient
	c := ovClient.NewOVClient(
		ovUser,
		ovPassword,
		"",
		ovUrl,
		defaultOvSslVerification, //ssl verificcation
		ovVer,
		"*")

	fmt.Printf("%v Trying to connect HPE OneView %s\n", time.Now(), ovUrl)
	events, err := oneview.GetEventList(c)
	if err != nil {
		fmt.Printf("Error: Could not get events: %v", err)
	}

	l := &oneview.Logging{
		OvAddr:     ovAddr,
		Path:       path,
		MaxSize:    size,
		MaxBackups: backups,
		MaxAge:     age,
		Compress:   compress,
		Events:     events,
	}

	if err := l.Write(); err != nil {
		fmt.Printf("Error: Could not logging events: %v", err)
	}
	fmt.Printf("%v Logged events into %s\n", time.Now(), path)
}
