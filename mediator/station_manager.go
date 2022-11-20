package mediator

import "sync"

type StationManager struct {
	IsPlatformFree bool
	lock           *sync.Mutex
	TrainQueue     []Train
}

func NewStationManager() *StationManager {
	return &StationManager{
		IsPlatformFree: true,
		lock:           &sync.Mutex{},
	}
}

func (s *StationManager) CanLand(t Train) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.IsPlatformFree {
		s.IsPlatformFree = false
		return true
	}
	s.TrainQueue = append(s.TrainQueue, t)
	return false
}

func (s *StationManager) NotifyFree() {
	s.lock.Lock()
	defer s.lock.Unlock()
	if !s.IsPlatformFree {
		s.IsPlatformFree = true
	}
	if len(s.TrainQueue) > 0 {
		firstTrainInQueue := s.TrainQueue[0]
		s.TrainQueue = s.TrainQueue[1:]
		firstTrainInQueue.PermitArrival()
	}
}
