package cache_test

import (
	"testing"
	"time"

	"github.com/zs5460/cache"
)

var defaultTTL = 20 * time.Millisecond

func TestGetSet(t *testing.T) {
	c := cache.New(defaultTTL)
	defer c.Close()

	c.Set("foo", "bar")

	foo, found := c.Get("foo")
	if !found {
		t.FailNow()
	}

	if foo.(string) != "bar" {
		t.FailNow()
	}

	time.Sleep(defaultTTL)

	_, found = c.Get("foo")
	if found {
		t.FailNow()
	}

	_, found = c.Get("404")
	if found {
		t.FailNow()
	}

}

func TestRange(t *testing.T) {
	c := cache.New(20 * time.Second)
	defer c.Close()
	c.Set("foo", "bar")
	time.Sleep(32 * time.Second)
	_, found := c.Get("foo")
	if found {
		t.FailNow()
	}
}

func TestDelete(t *testing.T) {
	c := cache.New(defaultTTL)
	c.Set("foo", "bar")
	_, found := c.Get("foo")
	if !found {
		t.FailNow()
	}

	c.Delete("foo")

	_, found = c.Get("foo")
	if found {
		t.FailNow()
	}
}

func BenchmarkNew(b *testing.B) {
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cache.New(defaultTTL).Close()
		}
	})
}

func BenchmarkGet(b *testing.B) {
	c := cache.New(defaultTTL)
	defer c.Close()
	c.Set("foo", "bar")

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Get("foo")
		}
	})
}

func BenchmarkSet(b *testing.B) {
	c := cache.New(defaultTTL)
	defer c.Close()

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Set("foo", "bar")
		}
	})
}

func BenchmarkDelete(b *testing.B) {
	c := cache.New(defaultTTL)
	defer c.Close()

	b.ResetTimer()
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Delete("foo")
		}
	})
}
