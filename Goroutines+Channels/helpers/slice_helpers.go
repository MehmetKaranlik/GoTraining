package helpers

type Slice[T, V comparable] []T

func (slice *Slice[T, V]) Append(item T) {
	*slice = append(*slice, item)
}

func (slice *Slice[T, V]) Remove(item T) {
	index := slice.IndexOf(item)
	if index == -1 {
		return
	}

	slice.RemoveAt(index)

}

func (slice *Slice[T, V]) RemoveAt(index int) {
	*slice = append((*slice)[:index], (*slice)[index+1:]...)
}

func (slice *Slice[T, V]) IndexOf(item T) int {
	for i, v := range *slice {
		if v == item {
			return i
		}
	}
	return -1
}
