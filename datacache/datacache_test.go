package datacache

import (
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	interval := time.Minute * 5
	cache := NewCache(interval)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddToCache(t *testing.T) {
	interval := time.Minute * 5
	cache := NewCache(interval)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
	}

	for _, cs := range cases {
		cache.Add(cs.inputKey, cs.inputVal)
		actual, ok := cache.Get(cs.inputKey)
		if !ok {
			t.Errorf("%s not found", cs.inputKey)
		}
		if string(actual) != string(cs.inputVal) {
			t.Errorf("%s doesn't match %s for key %s", string(actual), cs.inputVal, cs.inputKey)
		}
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	key := "test"
	cache.Add(key, []byte("icles"))

	time.Sleep(interval + time.Millisecond*5)

	_, ok := cache.Get(key)
	if ok {
		t.Errorf("%s should've been reaped", key)
	}
}

func TestReapFail(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	key := "test"
	cache.Add(key, []byte("icles"))

	time.Sleep(time.Millisecond)

	_, ok := cache.Get(key)
	if !ok {
		t.Errorf("%s should not have been reaped", key)
	}
}
