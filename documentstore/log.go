package documentstore

import "time"

type LogAction string

const (
	LogCollectionCreate LogAction = "collection.create"
	LogCollectionDelete LogAction = "collection.delete"

	LogDocumentCreate LogAction = "document.create"
	LogDocumentUpdate LogAction = "document.update"
	LogDocumentDelete LogAction = "document.delete"
)

type LogEntry struct {
	At         time.Time `json:"at"`
	Action     LogAction `json:"action"`
	Collection string    `json:"collection"`
	Key        string    `json:"key,omitempty"`
}

func (s *Store) addLog(e LogEntry) {
	s.logs = append(s.logs, e)
}

func (s *Store) Logs() []LogEntry {
	out := make([]LogEntry, len(s.logs))
	copy(out, s.logs)
	return out
}
