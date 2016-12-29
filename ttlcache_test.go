package ttlcache

import (
	"testing"
	"time"
)

func TestGetExisting(t *testing.T) {
	c := New(1 * time.Minute)
	key, value := "key", "value"
	c.Set(key, value)

	gotValue, ok := c.Get(key)

	if !ok {
		t.Error("ok = false, want true")
	}

	if got, want := gotValue, value; got != want {
		t.Errorf("c.Get(%q) = %v, want %q", key, got, want)
	}
}

func TestGetNonExistent(t *testing.T) {
	c := New(1 * time.Minute)
	key := "key"
	c.Set(key, "value")

	gotValue, ok := c.Get("no-key")

	if ok {
		t.Error("ok = true, want false")
	}

	if gotValue != nil {
		t.Errorf("c.Get(%q) = %v, want nil", key, gotValue)
	}
}

func TestGetTTL(t *testing.T) {
	c := New(10 * time.Millisecond)
	key, value := "key", "value"

	c.Set(key, value)
	time.Sleep(20 * time.Millisecond)

	gotValue, ok := c.Get(key)

	if gotValue != nil || ok {
		t.Errorf("c.Get(%q) = %v, %v; want <nil>, false", key, gotValue, ok)
	}
}

func TestExpire(t *testing.T) {
	c := New(1 * time.Minute)
	key, value := "key", "value"

	c.Set(key, value)
	c.Expire(key)

	gotValue, ok := c.Get(key)

	if gotValue != nil || ok {
		t.Errorf("c.Get(%q) = %v, %v; want <nil>, false", key, gotValue, ok)
	}
}

func TestExpireAll(t *testing.T) {
	c := New(1 * time.Minute)

	for i := 0; i < 10; i++ {
		c.Set(i, "value")
	}

	c.ExpireAll()

	key := 0
	gotValue, ok := c.Get(key)

	if gotValue != nil || ok {
		t.Errorf("c.Get(%d) = %q, %t; want <nil>, false", key, gotValue, ok)
	}

	if got, want := len(c.(*cache).items), 0; got != want {
		t.Errorf("cache has %d items, want %d", got, want)
	}
}

func TestSetTTLReset(t *testing.T) {
	c := New(20 * time.Millisecond)
	key, value := "key", "value"

	for i := 0; i < 10; i++ {
		c.Set(key, value)
		time.Sleep(10 * time.Millisecond)
	}

	gotValue, ok := c.Get(key)

	if gotValue != value || !ok {
		t.Errorf("c.Get(%q) = %v, %v; want %q, true", key, gotValue, ok, value)
	}
}
