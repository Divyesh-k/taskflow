package services

import (
	"sync"
	"taskflow/models"
)

type TaskCache struct {
	mu    sync.RWMutex
	store map[uint]models.Task
}

func NewTaskCache() *TaskCache {
	return &TaskCache{
		store: make(map[uint]models.Task),
	}
}

func (c *TaskCache) Set(task models.Task) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[task.ID] = task
}

func (c *TaskCache) Get(id uint) (models.Task, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	task, ok := c.store[id]
	return task, ok
}

func (c *TaskCache) Delete(id uint) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.store, id)
}

func (c *TaskCache) Invalidate() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store = make(map[uint]models.Task)
}

func (c *TaskCache) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.store)
}
