package pinger

import (
	"sync"

	"github.com/mcstatus-io/mcutil/v4/response"
)

var (
	Mu     sync.Mutex
	States map[string]*response.StatusModern = make(map[string]*response.StatusModern)
)