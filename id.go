// 29 july 2018
package main

import (
	"fmt"
	"strconv"
	"sync"
)

type ID int64

func IDFromString(s string) (ID, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return ID(i), nil
}

func (i ID) String() string {
	return fmt.Sprint(i)
}

type IDPool struct {
	mu		sync.Mutex
	next		ID
}

func NewIDPool() *IDPool {
	return new(IDPool)
}

func (i *IDPool) Mark(visited ID) {
	i.mu.Lock()
	defer i.mu.Unlock()
	if i.next < visited {
		i.next = visited
	}
}

func (i *IDPool) Next() ID {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.next++
	return i.next
}
