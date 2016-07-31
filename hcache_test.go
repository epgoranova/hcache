package hcache_test

import (
	"testing"

	. "github.com/s2gatev/hcache"
)

func testCache(test func(*Cache)) {
	cache := New()
	test(cache)
}

func TestInsertValue(t *testing.T) {
	testCache(func(cache *Cache) {
		cache.Insert(42, 'a', 'b', 'c')
		a := cache.GetOrInsert(13, 'a', 'b', 'c').(int)
		if a != 42 {
			t.Error()
		}
	})
}

func TestGetOrInsertSameValue(t *testing.T) {
	testCache(func(cache *Cache) {
		cache.GetOrInsert(42, 'a', 'b', 'c')
		b := cache.GetOrInsert(13, 'a', 'b', 'c').(int)
		if b != 42 {
			t.Error()
		}
	})
}

func TestGetOrInsertDifferentValues(t *testing.T) {
	testCache(func(cache *Cache) {
		cache.GetOrInsert(42, 'a', 'b', 'c')
		b := cache.GetOrInsert(13, 'a', 'c').(int)
		if b != 13 {
			t.Error()
		}
	})
}

func TestUsingHeterogeneousKeys(t *testing.T) {
	testCache(func(cache *Cache) {
		cache.GetOrInsert(42, 'a', 13, 3+2i)
		b := cache.GetOrInsert(13, 'a', 13, 3+2i).(int)
		if b != 42 {
			t.Error()
		}
	})
}

func TestHasVerifiesValueIsInserted(t *testing.T) {
	testCache(func(cache *Cache) {
		cache.GetOrInsert(42, 'a', 'b', 'c')

		if !cache.Has('a', 'b', 'c') {
			t.Error()
		}
	})
}

func TestHasFailsIfValueIsNotInserted(t *testing.T) {
	testCache(func(cache *Cache) {
		if cache.Has('a', 'b', 'c') {
			t.Error()
		}
	})
}

func TestErasingValue(t *testing.T) {
	testCache(func(cache *Cache) {
		cache.GetOrInsert(42, 'a', 'b', 'c')
		b := cache.Erase('a', 'b', 'c').(int)
		if b != 42 {
			t.Error()
		}
		if cache.Has('a', 'b', 'c') {
			t.Error()
		}
	})
}

func TestErasingNonPresentValue(t *testing.T) {
	testCache(func(cache *Cache) {
		a := cache.Erase("something", 'a', 'b', 'c')
		if a != nil {
			t.Error()
		}
	})
}

func TestGetValue(t *testing.T) {
	testCache(func(cache *Cache) {
		cache.GetOrInsert(42, 'a', 'b', 'c')
		a, ok := cache.Get('a', 'b', 'c')
		if a != 42 || !ok {
			t.Error()
		}
	})
}

func TestGetNonPresentValue(t *testing.T) {
	testCache(func(cache *Cache) {
		a, ok := cache.Get('a', 'b', 'c')
		if a != nil || ok {
			t.Error()
		}
	})
}
