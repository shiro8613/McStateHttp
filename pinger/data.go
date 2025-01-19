package pinger

import "sync"

var (
	Mu     sync.Mutex
	States map[string]bool = make(map[string]bool)
)