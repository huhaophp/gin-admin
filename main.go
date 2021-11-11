package main

import (
	"context"
	"fmt"
	"github.com/huhaophp/gin-admin/config"
	"github.com/huhaophp/gin-admin/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())

	server := &http.Server{
		Addr:           fmt.Sprintf("%s:%s", config.C.Server.Host, config.C.Server.Port),
		Handler:        routes.Register(engine),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 1 seconds.")
	}
}
