package app

import "go-microservice/internal/config"

func Init() {
	config.Init()
	//database.InitDB()
	//logger.SetupLogger(&configs.Log)
	//
	//err := metrics.InitiateStatsDMetrics(configs.StatsD)
	//if err != nil {
	//	logger.Errorf("error initiating statsD %+v", err)
	//}
	//
	//metrics.InitNewrelic(configs.NewRelic)
}

func ShutDown() {
	//database.CloseDB()
}
