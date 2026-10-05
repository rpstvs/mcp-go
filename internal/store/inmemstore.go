package store

import (
	"github.com/rpstvs/mcp-go/internal/workflow"
)

type InMemStore struct {
	runs map[string]workflow.Run
}

func NewInMemStore() *InMemStore {
	return &InMemStore{
		runs: make(map[string]workflow.Run),
	}
}

func (s *InMemStore) CreateRun(run workflow.Run) {
	s.runs[string(run.ID)] = run
}

func (s *InMemStore) GetRun(runId string) workflow.Run {
	val, ok := s.runs[runId]

	if !ok {
		return nil
	}
	return val
}

func (s *InMemStore) UpdateRun(run workflow.Run) {
	_, ok := s.runs[run.ID]

	if ok {
		s.runs[run.ID] = run
	}
}

func (s *InMemStore) DeleteRun(runid string) {
	delete(s.runs, runid)
}
