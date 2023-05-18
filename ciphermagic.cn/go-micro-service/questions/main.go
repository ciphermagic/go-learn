package main

import (
	"ciphermagic.cn/go-micro-service/questions/pkg/config"
	"ciphermagic.cn/go-micro-service/questions/pkg/logger"
)

func main() {
	logger.Info("http server started, listened on port %v", config.Config.Port)
}
