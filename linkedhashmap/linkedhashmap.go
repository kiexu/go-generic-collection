package linkedhashmap

import (
	"container/list"
	"github.com/kiexu/go-generic-collection"
)

var _ gcollection.Map[struct{}, struct{}] = new(LinkedHashMap[struct{}, struct{}])

// LinkedHashMap
// not thread-safe
type LinkedHashMap[K, V comparable] struct {
	entryList *list.List          // list.Element contains entries
	entryMap  map[K]*list.Element // list.Element map
}

func NewLinkedHashMap[K, V comparable]() *LinkedHashMap[K, V] {
	return &LinkedHashMap[K, V]{
		entryList: list.New(),
		entryMap:  make(map[K]*list.Element),
	}
}

// Size get current size
func (m *LinkedHashMap[K, V]) Size() int {
	return m.entryList.Len()
}

// IsEmpty check is empty
func (m *LinkedHashMap[K, V]) IsEmpty() bool {
	return m.Size() == 0
}

// ContainsKey check if key exists
func (m *LinkedHashMap[K, V]) ContainsKey(key K) bool {
	_, exist := m.entryMap[key]
	return exist
}

// ContainsValue check if value exists, it takes O(n)
func (m *LinkedHashMap[K, V]) ContainsValue(val V) bool {

	for e := m.entryList.Front(); e != nil; e = e.Next() {
		if val == e.Value.(gcollection.MapEntry[K, V]).Value() {
			return true
		}
	}

	return false
}

// Put a key-value and try return previous value of the given key
// return: previous: previous value
// 		   exist: if previous value exists
func (m *LinkedHashMap[K, V]) Put(key K, value V) (previous V, exist bool) {

	previousElement, exist := m.entryMap[key]
	if !exist {
		m.entryMap[key] = m.entryList.PushBack(&LinkedHashEntry[K, V]{
			key:   key,
			value: value,
		})
		return
	}
	previousEntry := previousElement.Value.(gcollection.MapEntry[K, V])
	previous = previousEntry.Value()
	previousEntry.SetValue(value)

	return
}

// Get value by key
// return: value: value of given key
// 		   exist: if value exists
func (m *LinkedHashMap[K, V]) Get(key K) (value V, exist bool) {

	element, exist := m.entryMap[key]
	if !exist {
		return
	}
	return element.Value.(gcollection.MapEntry[K, V]).Value(), true

}

// Remove key-value
// return: success or not exist
func (m *LinkedHashMap[K, V]) Remove(key K) bool {

	element, exist := m.entryMap[key]
	if !exist {
		return false
	} else {
		m.entryList.Remove(element)
		delete(m.entryMap, key)
		return true
	}
}

// Clear all entries
func (m *LinkedHashMap[K, V]) Clear() {
	m.entryList = list.New()
	m.entryMap = make(map[K]*list.Element)
}

// Keys get slice of keys in put-order
func (m *LinkedHashMap[K, V]) Keys() []K {

	res := make([]K, 0, m.entryList.Len())

	for e := m.entryList.Front(); e != nil; e = e.Next() {
		res = append(res, e.Value.(gcollection.MapEntry[K, V]).Key())
	}

	return res
}

// Values get slice of values in put-order
func (m *LinkedHashMap[K, V]) Values() []V {

	res := make([]V, 0, m.entryList.Len())

	for e := m.entryList.Front(); e != nil; e = e.Next() {
		res = append(res, e.Value.(gcollection.MapEntry[K, V]).Value())
	}

	return res
}

// Entries get slice of entries in put-order
func (m *LinkedHashMap[K, V]) Entries() []gcollection.MapEntry[K, V] {

	res := make([]gcollection.MapEntry[K, V], 0, m.entryList.Len())

	for ele := m.entryList.Front(); ele != nil; ele = ele.Next() {
		res = append(res, ele.Value.(gcollection.MapEntry[K, V]))
	}

	return res
}
