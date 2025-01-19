package pinger

import (
	"sync"
)

var (
	Mu     sync.Mutex
	States map[string]map[string]interface{} = make(map[string]map[string]interface{})

)