package pinger

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/McStateHttp/config"
	"github.com/mcstatus-io/mcutil/v4/response"
	"github.com/mcstatus-io/mcutil/v4/status"
)

func Run(ctx context.Context) {
	conf := config.GetConfig()
	ticker := time.NewTicker(time.Second * time.Duration(conf.PingTime))
	
	defer ticker.Stop()

	pinger(conf)
	
	for {
		select {
		case <-ticker.C:
			pinger(conf)
		case <- ctx.Done():
			break
		}
	}
}

func pinger(conf config.Config) {
	for name, server_ip := range conf.Servers {
		log.Printf("[INFO] %s(%s) is pinging now\n", name, server_ip)
		go func(sip string, tout int) {	
			bo, err := ping(sip, tout)
			if err != nil {
				log.Printf("[ERROR][%s] %s\n", name, err)
			}
			Mu.Lock()
			States[name] = bo
			Mu.Unlock()
		}(server_ip, conf.Timeout)
	}
}

func ping(server_ip string, timeout int) (*response.StatusModern, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout) * time.Second)

	defer cancel()

	h := strings.Split(server_ip, ":")
	host := h[0]
	port, err := strconv.Atoi(h[1])
	if err != nil {
		return nil, err
	}
	resp, err := status.Modern(ctx, host, uint16(port))
	if err != nil {
		return nil, err
	}

	
	return resp, nil
}