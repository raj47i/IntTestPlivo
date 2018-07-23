package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/raj47i/IntTestPlivo/config"

	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := config.Get()
	if !cfg.Debug {
		log.Info("Debug mode is off")
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	InitRoutes(r)

	if !cfg.Debug {
		log.Infof("starting http server on %d ..", cfg.Port)
	}
	r.Run(fmt.Sprintf(":%d", cfg.Port))
}

// InitRoutes defines all the routes, exported so its convenient to test
func InitRoutes(r *gin.Engine) {
	authOnlyRoute := r.Group("/")
	authOnlyRoute.Use(basicAuthMiddleware)
	authOnlyRoute.POST("/inbound/sms/", smsInbound)
	authOnlyRoute.POST("/outbound/sms/", smsOutbound)

	r.NoRoute(http405)
}
