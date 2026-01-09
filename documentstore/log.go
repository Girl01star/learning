package documentstore

import "log/slog"

type LogType string

const (
	LogCollectionCreate LogType = "collection.create"
	LogCollectionDelete LogType = "collection.delete"

	LogDocumentCreate LogType = "document.create"
	LogDocumentUpdate LogType = "document.update"
	LogDocumentDelete LogType = "document.delete"
)

type LogEntry struct {
	Type       LogType
	Collection string
	Key        string
	Action     any
}

func (s *Store) addLog(t LogType, collection, key string) {
	entry := LogEntry{
		Type:       t,
		Collection: collection,
		Key:        key,
	}
	s.logs = append(s.logs, entry)

	if s.logger == nil {
		s.logger = slog.Default()
	}

	s.logger.Info(
		"documentstore event",
		slog.String("type", string(entry.Type)),
		slog.String("collection", entry.Collection),
		slog.String("key", entry.Key),
	)
}
