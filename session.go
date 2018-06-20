package botframework

import (
	"log"
	"sync"
	"time"
)

const (
	defaultTtl = 400000
)

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

type SessionEntity struct {
	updated  int64
	ValueMap map[string]interface{}
}

type SessionStorage struct {
	sessionMap map[string]*SessionEntity
	mutex      *sync.Mutex
	defaultTtl int64
	started    bool
}

func NewSessionStorage() *SessionStorage {
	storage := &SessionStorage{
		sessionMap: make(map[string]*SessionEntity),
		mutex:      &sync.Mutex{},
		defaultTtl: defaultTtl,
		started:    true,
	}

	go func() {
		for {
			<-time.Tick(time.Second)
			log.Printf("calling gc")
			storage.gc()
			if !storage.started {
				return
			}
		}
	}()
	return storage
}

func (s *SessionStorage) get(user string) *SessionEntity {
	if !s.started {
		panic("session storage is not started yet")
	}
	_, exists := s.sessionMap[user]
	s.mutex.Lock()
	if !exists {
		s.sessionMap[user] = &SessionEntity{
			ValueMap: make(map[string]interface{}),
		}
	}
	s.sessionMap[user].updated = makeTimestamp()
	s.mutex.Unlock()
	return s.sessionMap[user]
}
func (s *SessionStorage) cleanUser(user string) {
	s.mutex.Lock()
	if s.sessionMap[user].updated+s.defaultTtl > makeTimestamp() {
		delete(s.sessionMap, user)
	}
	s.mutex.Unlock()
}

func (s *SessionStorage) gc() {
	for k := range s.sessionMap {
		s.cleanUser(k)
	}
}
