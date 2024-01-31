package server

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"test-go/internal/config"
	"test-go/internal/event"
	"test-go/internal/handlers/get"
	"test-go/internal/logger"
	"test-go/internal/storage"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/patrickmn/go-cache"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

type Server struct {
	e   *echo.Echo
	log *slog.Logger
}

func New(cfg config.Config, saver event.Saver, getter get.Getter) *Server {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	log := setupLogger(cfg.Env)

	c, err := getter.GetOrdes(cfg.CacheLimit)
	if err != nil {
		panic(fmt.Sprintf("can't restore cache: %s", err.Error()))
	}

	eH, err := event.New(cfg.ClusterId, cfg.ClientId)
	if err != nil {
		panic(fmt.Sprintf("can't connect to nats: %s", err.Error()))
	}

	err = eH.HandleEvent(log, cfg.SubjectNats, cfg.DurableName, saver, c)
	if err != nil {
		panic(fmt.Sprintf("can't subscribe: %s", err.Error()))
	}

	e.GET("/orders/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")
		res, ok := c.Get(id)
		if ok {
			return ctx.JSON(http.StatusOK, res)
		}
		res, err := get.GetById(ctx, getter)
		if errors.Is(err, storage.ErrEntryNotFound) {
			log.Error("order not found", logger.Err(err))
			return ctx.String(http.StatusNotFound, err.Error())
		}
		if err != nil {
			log.Error("can't get order", logger.Err(err))
			return ctx.String(http.StatusInternalServerError, err.Error())
		}
		c.Set(id, res, cache.DefaultExpiration)
		return ctx.JSON(http.StatusOK, res)
	})

	return &Server{
		e:   e,
		log: log,
	}
}

func (s *Server) Run(port int) error {
	s.log.Info("start server")
	err := s.e.Start(fmt.Sprintf(":%d", port))

	if err != nil {
		return err
	}
	return err
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(
				os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(
				os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(
				os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}
