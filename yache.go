package yache

import "time"

type Item struct {
	Expires int64
	Value   any
}

type Store struct {
	Timeout          int64
	Autoflush        bool
	AutoflushStarted bool
	Data             map[string]Item
}

func NewStore(timeout int64, autoflush bool) Store {
	return Store{
		Timeout:   timeout,
		Autoflush: autoflush,
		Data:      map[string]Item{},
	}
}

func (s *Store) Set(key string, val any) {
	s.Data[key] = Item{
		Expires: time.Now().Unix() + s.Timeout,
		Value:   val,
	}
	if !s.AutoflushStarted {
		go func() {
			for {
				time.Sleep(time.Second * time.Duration(s.Timeout+1))
				s.Flush()
			}
		}()
		s.AutoflushStarted = true
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

func (s *Store) Flush() {
	for key := range s.Data {
		if time.Now().Unix() > s.Data[key].Expires {
			delete(s.Data, key)
		}
	}
}
