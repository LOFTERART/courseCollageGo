package main

import (
	"PC/cache"
	"PC/models"
	"PC/routers"
	"PC/tasks"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"context"
)

func init() {
	cache.Redis()
	models.Initialized()
	// 启动定时任务
	tasks.CronJob()

}

func main() {

	gin.SetMode(gin.ReleaseMode)
	router := routers.InitRouter()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	go func() {

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}
