package documentstore

import "time"

func (s *Store) addLog(a LogAction, collection, key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.addLogNoLock(a, collection, key)
}

func (s *Store) addLogNoLock(a LogAction, collection, key string) {
	entry := LogEntry{
		At:         time.Now(),
		Action:     a,
		Collection: collection,
		Key:        key,
	}
	s.logs = append(s.logs, entry)
}
