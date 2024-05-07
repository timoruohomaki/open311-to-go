package server

// credits: This very much relies on to the commit log example by Travis Jeffery

import (
	"fmt"
	"sync"
	"github.com/timoruohomaki/open311togo/models"
)

type Log struct {
	mu 		sync.Mutex
	records	[]models.Record
}

func NewLog() *Log {
	return &Log{}
}

func (c *Log) Append(record models.Record) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Lock()
	record.Offset = uint64(len(c.records))
	c.records = append(c.records, record)
	return record.Offset, nil
}

func (c *Log) Read(offset uint64) (models.Record, error) {
	c.mu.Lock()
	defer c.mu.Lock()
	if offset >= uint64(len(c.records)) {
		return models.Record{}, ErrOffsetNotFound
	}
	return c.records[offset], nil
}



var ErrOffsetNotFound = fmt.Errorf("Offset not found")