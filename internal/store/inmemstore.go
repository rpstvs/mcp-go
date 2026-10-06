package store

import (
	"github.com/rpstvs/mcp-go/types"
)

type InMemStore struct {
	runs map[string]types.Run
}

func NewInMemStore() *InMemStore {
	return &InMemStore{
		runs: make(map[string]types.Run),
	}
}

func (s *InMemStore) CreateRun(run *types.Run) {
	s.runs[string(run.ID)] = *run
}

func (s *InMemStore) GetRun(runId string) *types.Run {
	val, ok := s.runs[runId]

	if !ok {
		return nil
	}
	return &val
}

func (s *InMemStore) UpdateRun(run *types.Run) {
	_, ok := s.runs[run.ID]

	if ok {
		s.runs[run.ID] = *run
	}
}

func (s *InMemStore) DeleteRun(runid string) {
	delete(s.runs, runid)
}
