package main

import (
	//"runtime"

	"github.com/nicholaskh/dist-db/config"
	"github.com/nicholaskh/dist-db/network"
	"github.com/nicholaskh/golib/server"
)

func init() {
	//runtime.GOMAXPROCS(runtime.NumCPU())
	parseFlags()

	if options.showVersion {
		server.ShowVersionAndExit()
	}

	server.SetupLogging(options.logFile, options.logLevel, options.crashLogFile)

	conf := server.LoadConfig(options.configFile)
	config.DistDb = new(config.DistDbConfig)
	config.DistDb.LoadConfig(conf)
}

func main() {
	distDbServer := server.NewTcpServer("dist-db")
	distDbServer.SetProtoType(server.SIMPLE)
	processor := network.NewProcessor(distDbServer)
	distDbServer.LaunchTcpServer(config.DistDb.ListenAddr, processor, config.DistDb.SessionTimeout, config.DistDb.ServInitialGoroutineNum)
	<-make(chan int)
}
