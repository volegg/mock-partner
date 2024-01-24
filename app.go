package main

import (
	"bitbucket.org/1-pixel-games/mock-partner/internal/health"
	"bitbucket.org/1-pixel-games/mock-partner/internal/partner"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type application struct {
	httpSrv *fiber.App
	logic   *partner.Service
}

func newApplication(config *config) (*application, error) {
	var app = new(application)
	app.logic = partner.New(config.CoreURL)

	greeting(config)

	app.httpSrv = app.initHTTPSrv()

	return app, nil
}

// @title                      1 Pixel Games WebSocket API
// @version                    1.0
// @schemes                    http ws
// @securityDefinitions.apikey Bearer
// @in                         header
// @name                       Authorization
// @description                Type "Bearer" followed by a space and JWT token.
func (a *application) initHTTPSrv() *fiber.App {
	srv := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	//srv.Use(recover.New())

	health.RegisterHandler(srv)
	a.logic.RegisterHandler(srv)

	return srv
}

func greeting(config *config) {
	log.Infof("Listening on port %s", config.HTTPPort)
}