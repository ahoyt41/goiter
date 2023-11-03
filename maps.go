package goiter

type MapPair[A comparable, B any] struct {
	Key   A
	Value B
}

func ToMap[A comparable, B any](iter Iterator[MapPair[A, B]]) map[A]B {
	out := make(map[A]B)
	for {
		p, ok := iter.Next()
		if !ok {
			break
		}
		out[p.Key] = p.Value
	}
	return out
}

func FromMap[A comparable, B any](mapping map[A]B) Iterator[MapPair[A, B]] {
	itemChan := make(chan MapPair[A, B])
	go func() {
		defer close(itemChan)
		for k, v := range mapping {
			itemChan <- MapPair[A, B]{k, v}
		}
	}()
	return &channelIter[MapPair[A, B]]{itemChan}
}
