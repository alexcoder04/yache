package yache

import "time"

type Item struct {
	Expires int64
	Value   any
}

type Store struct {
	Timeout int64
	Data    map[string]Item
}

func NewStore(timeout int64) Store {
	return Store{
		Timeout: timeout,
		Data:    map[string]Item{},
	}
}

func (s *Store) Set(key string, val any) {
	s.Data[key] = Item{
		Expires: time.Now().Unix() + s.Timeout,
		Value:   val,
	}
}

func (s *Store) Get(key string) any {
	if val, hasKey := s.Data[key]; hasKey {
		if time.Now().Unix() > s.Data[key].Expires {
			delete(s.Data, key)
			return nil
		}
		return val
	}
	return nil
}
