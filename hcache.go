package hcache

// Key represents a single key used in the hirearchical cache.
type Key interface{}

// Value represents a value stored in the hirearchical cache.
type Value interface{}

type cacheBox struct {
	next  map[Key]*cacheBox
	value Value
}

func newCacheBox() *cacheBox {
	return &cacheBox{next: map[Key]*cacheBox{}}
}

// Cache is a store that keeps values in a hierarchical fashion.
type Cache struct {
	root *cacheBox
}

// New creates a hierarchical cache object.
func New() *Cache {
	return &Cache{root: newCacheBox()}
}

func (c *Cache) getBox(box *cacheBox, keys []Key) *cacheBox {
	if box == nil || len(keys) == 0 {
		return box
	}

	nextBox := box.next[keys[0]]
	return c.getBox(nextBox, keys[1:])
}

func (c *Cache) getOrInsertBox(box *cacheBox, keys []Key) *cacheBox {
	if len(keys) == 0 {
		return box
	}

	nextBox, ok := box.next[keys[0]]
	if !ok {
		nextBox = newCacheBox()
		box.next[keys[0]] = nextBox
	}

	return c.getOrInsertBox(nextBox, keys[1:])
}

// Insert adds an entry in the cache following keys.
//
// If there's already a value in the specific location it will be replaced
// and returned as a result. If the value is new the result of Insert will
// be nil.
func (c *Cache) Insert(value Value, keys ...Key) Value {
	box := c.getOrInsertBox(c.root, keys)

	oldValue := box.value
	box.value = value

	return oldValue
}

// GetOrInsert retrieves an entry from the cache following keys.
//
// If the entry is not present it is created with the specified value.
func (c *Cache) GetOrInsert(value Value, keys ...Key) Value {
	box := c.getOrInsertBox(c.root, keys)

	if box.value == nil {
		box.value = value
	}

	return box.value
}

// Has returns true if an entry is present in the cache following keys
// and false otherwise.
func (c *Cache) Has(keys ...Key) bool {
	box := c.getBox(c.root, keys)

	return box != nil && box.value != nil
}

// Erase removes an entry from the cache following keys and returns it.
//
// If the entry is not present in the cache Erase returns nil.
func (c *Cache) Erase(keys ...Key) Value {
	box := c.getBox(c.root, keys)

	var oldValue Value
	if box != nil {
		oldValue = box.value
		box.value = nil
	}

	return oldValue
}

// Get retrieves an entry, if present, from the cache following keys.
func (c *Cache) Get(keys ...Key) (Value, bool) {
	box := c.getBox(c.root, keys)

	if box == nil {
		return nil, false
	}

	return box.value, true
}
