package pokeapi

import (
	"testing"
	"time"
)

func TestReapDeletesEntry(t *testing.T) {
	url := "asdf"
	data := []byte("fdsa")
	cache := NewCache(time.Millisecond * 25, time.Millisecond * 200)
	cache.Add(url, data)

    time.Sleep(time.Millisecond * 350)

	_, ok := cache.Get(&url)
	if ok {
		t.Error("Reap does not delete entry")
	}
}

func TestReapPreservesEntry(t *testing.T) {
	url := "asdf"
	data := []byte("fdsa")
	cache := NewCache(time.Millisecond * 25, time.Millisecond * 200)
	cache.Add(url, data)

    time.Sleep(time.Millisecond * 150)

	_, ok := cache.Get(&url)
	if !ok {
		t.Error("Reap deletes entry")
	}
}