package linkedhashmap

import gcollection "kiexu/go-generic-collection"

var _ gcollection.MapEntry[struct{}, struct{}] = new(LinkedHashEntry[struct{}, struct{}])

type LinkedHashEntry[K, V comparable] struct {
	key   K
	value V
}

func (e *LinkedHashEntry[K, V]) Key() K {
	return e.key
}

func (e *LinkedHashEntry[K, V]) SetKey(key K) {
	e.key = key
}

func (e *LinkedHashEntry[K, V]) Value() V {
	return e.value
}

func (e *LinkedHashEntry[K, V]) SetValue(value V) {
	e.value = value
}
