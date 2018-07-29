// 29 july 2018
package main

import (
	"sync"
)

type ID int

type IDPool struct {
	next		ID
	mu		sync.Mutex
}

func (i *IDPool) Generate() ID {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.next++
	return i.next
}
