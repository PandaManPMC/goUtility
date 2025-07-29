package util

import (
	"fmt"
	"testing"
)

func TestRingQueue(t *testing.T) {
	q := NewRingQueue(3)

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println(q.Values()) // [1 2 3]

	q.Enqueue(4)
	fmt.Println(q.Values()) // [2 3 4]  (1 被覆盖)

	q.Enqueue(5)
	fmt.Println(q.Values()) // [3 4 5]

	oldest, _ := q.Oldest()
	latest, _ := q.Latest()
	fmt.Println("Oldest:", oldest)
	fmt.Println("Latest:", latest)

	val, ok := q.Get(2)
	fmt.Println("Get(1):", val, ok)

	q.Enqueue(6)
	fmt.Println(q.Values())

	oldest, _ = q.Oldest()
	latest, _ = q.Latest()
	fmt.Println("Oldest:", oldest)
	fmt.Println("Latest:", latest)

	q.View(func(i int, v interface{}) {
		t.Log(i, "::", v)
	})

	t.Log(q.Data())
}
