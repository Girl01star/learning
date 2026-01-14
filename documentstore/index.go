package documentstore

import "sort"

type indexEntry struct {
	Value string `json:"value"`
	Key   string `json:"key"`
}

type Index struct {
	FieldName string       `json:"fieldName"`
	Entries   []indexEntry `json:"entries"`
}

func NewIndex(fieldName string) *Index {
	return &Index{
		FieldName: fieldName,
		Entries:   make([]indexEntry, 0),
	}
}

func (idx *Index) Add(value, key string) {
	ent := indexEntry{Value: value, Key: key}

	pos := sort.Search(len(idx.Entries), func(i int) bool {
		if idx.Entries[i].Value == ent.Value {
			return idx.Entries[i].Key >= ent.Key
		}
		return idx.Entries[i].Value >= ent.Value
	})

	idx.Entries = append(idx.Entries, indexEntry{})
	copy(idx.Entries[pos+1:], idx.Entries[pos:])
	idx.Entries[pos] = ent
}

func (idx *Index) Remove(value, key string) {

	l := sort.Search(len(idx.Entries), func(i int) bool {
		return idx.Entries[i].Value >= value
	})
	r := sort.Search(len(idx.Entries), func(i int) bool {
		return idx.Entries[i].Value > value
	})

	if l >= r {
		return
	}

	for i := l; i < r; i++ {
		if idx.Entries[i].Key == key {
			copy(idx.Entries[i:], idx.Entries[i+1:])
			idx.Entries = idx.Entries[:len(idx.Entries)-1]
			return
		}
	}
}

func (idx *Index) RangeKeys(min, max *string, desc bool) []string {
	start := 0
	if min != nil {
		start = sort.Search(len(idx.Entries), func(i int) bool {
			return idx.Entries[i].Value >= *min
		})
	}

	end := len(idx.Entries)
	if max != nil {
		end = sort.Search(len(idx.Entries), func(i int) bool {
			return idx.Entries[i].Value > *max
		})
	}

	if start >= end {
		return []string{}
	}

	keys := make([]string, 0, end-start)

	if !desc {
		for i := start; i < end; i++ {
			keys = append(keys, idx.Entries[i].Key)
		}
		return keys
	}

	for i := end - 1; i >= start; i-- {
		keys = append(keys, idx.Entries[i].Key)
	}
	return keys
}
