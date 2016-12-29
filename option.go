package main

import (
	"flag"
)

var (
	options struct {
		configFile   string
		logFile      string
		logLevel     string
		crashLogFile string
		showVersion  bool
	}
)

func parseFlags() {
	flag.StringVar(&options.configFile, "c", "etc/dist-db.cf", "config file")
	flag.BoolVar(&options.showVersion, "v", false, "show version and exit")
	flag.StringVar(&options.logFile, "log", "stdout", "log file")
	flag.StringVar(&options.logLevel, "level", "info", "log level")
	flag.StringVar(&options.crashLogFile, "crashlog", "panic.dump", "crash log file")

	flag.Parse()
}
