package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/McStateHttp/config"
	"github.com/McStateHttp/pinger"
	"github.com/gin-gonic/gin"
)

func Run(ctx context.Context) error {
	conf := config.GetConfig()
	r := gin.Default()
	r.GET("/ping/:server_name", handle)

	srv := &http.Server {
		Addr: conf.Bind,
		Handler: r,
	}

	go func() {
		<- ctx.Done()
		ctx1, cancel := context.WithTimeout(ctx, 5 * time.Second)

		defer cancel()

		if err := srv.Shutdown(ctx1); err != nil {
			log.Printf("[ERROR] %s" , err)
		}
	}()

	err := srv.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func handle(c *gin.Context) {
	name := c.Param("server_name")

	pinger.Mu.Lock()
	state, ok := pinger.States[name]
	pinger.Mu.Unlock()

	if !ok {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.AbortWithStatusJSON(http.StatusOK, state)
}