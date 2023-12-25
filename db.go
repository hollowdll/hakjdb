package kvdb

import "sync"

type Database struct {
	name  string
	data  map[DatabaseKey]DatabaseValue
	mutex sync.RWMutex
}
