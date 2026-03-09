package services

import (
	"context"
	"taskflow/repository"
	"time"
)

type TaskCleaner struct {
	repo     repository.TaskRepositoryInterface
	cache    *TaskCache
	interval time.Duration
	maxAge   time.Duration
	cacel    context.CancelFunc
}

func NewTaskCleaner(repo repository.TaskRepositoryInterface, cache *TaskCache, interval, maxAge time.Duration) *TaskCleaner {
	return &TaskCleaner{
		repo:     repo,
		cache:    cache,
		interval: interval,
		maxAge:   maxAge,
	}
}
