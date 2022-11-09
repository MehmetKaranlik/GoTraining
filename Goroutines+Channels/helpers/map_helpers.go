package helpers

type Map[T, V comparable] map[T]V

func (this *Map[T, V]) Append(key T, value V) {
	(*this)[key] = value
}

func (this *Map[T, V]) Remove(key T) {
	delete(*this, key)
}

func (this *Map[T, V]) Has(key T) bool {
	_, ok := (*this)[key]
	return ok
}

func (this *Map[T, V]) Get(key T) V {
	return (*this)[key]
}

func (this *Map[T, V]) Set(key T, value V) {
	(*this)[key] = value
}
