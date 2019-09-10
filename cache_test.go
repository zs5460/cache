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

	c.Set("Hello", "World")
	hello, found := c.Get("Hello")

	if !found {
		t.Error("19")
		t.FailNow()
	}

	if hello.(string) != "World" {
		t.Error("24")
		t.FailNow()
	}

	time.Sleep(defaultTTL)

	_, found = c.Get("Hello")

	if found {
		t.Error("33")
		t.FailNow()
	}

	time.Sleep(defaultTTL)

	_, found = c.Get("404")

	if found {
		t.FailNow()
	}
}

func TestDelete(t *testing.T) {
	c := cache.New(defaultTTL)
	c.Set("Hello", "World")
	_, found := c.Get("Hello")

	if !found {
		t.FailNow()
	}

	c.Delete("Hello")

	_, found = c.Get("Hello")

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
	c.Set("Hello", "World")

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Get("Hello")
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
			c.Set("Hello", "World")
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
			c.Delete("Hello")
		}
	})
}
