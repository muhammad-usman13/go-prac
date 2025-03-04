package main

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	// Create a list of integers.
	var s *List[int]
	for i := 10; i >0; i-- {
		s = &List[int]{val: i, next: s}
	}

	// Iterate over the list and print its elements.
	for s != nil {
		println(s.val)
		s = s.next
	}

	

}