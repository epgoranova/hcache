// hcache is a library providing hierarchical value store.
//
// New value stores are created using the New function of the
// package, each keeping separate hierarchy of values.
//
//	cache := New()
//
// A new value is inserted using the Insert function.
//
//	cache.Insert(42, "key1", "key2", "key3")
//
// The value could be retrieved using the Get function.
//
//	value := cache.Get("key1", "key2", "key3")
//
// A value could be removed from the store using the Erase function.
//
//	cache.Erase("key1", "key2", "key3")
//
package hcache
