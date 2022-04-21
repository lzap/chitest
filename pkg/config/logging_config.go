package config

import (
	"os"
	"sync"
)

type LoggingConfig struct {
	Cloudwatch bool
	CWGroup    string
	CWStream   string
	AWSRegion  string
	AWSKey     string
	AWSSecret  string
	AWSSession string

	initialized bool
}

var loggingConfig LoggingConfig
var initMutex sync.Mutex

func initializeConfig() {
	loggingConfig = LoggingConfig{
		Cloudwatch:  os.Getenv("CLOUDWATCH") == "1",
		CWGroup:     os.Getenv("CLOUDWATCH_GROUP"),
		CWStream:    os.Getenv("CLOUDWATCH_STREAM"),
		AWSRegion:   os.Getenv("AWS_REGION"),
		AWSKey:      os.Getenv("AWS_KEY"),
		AWSSecret:   os.Getenv("AWS_SECRET"),
		AWSSession:  os.Getenv("AWS_SESSION"),
		initialized: true,
	}
}

func GetLoggingConfig() *LoggingConfig {
	initMutex.Lock()
	defer initMutex.Unlock()

	if !loggingConfig.initialized {
		initializeConfig()
	}

	return &loggingConfig
}
