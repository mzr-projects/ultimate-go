package semantics

import "testing"

var fa int

func BenchmarkLinkedListTraverse(b *testing.B) {

	var a int

	for i := 0; i < b.N; i++ {
		a = LinkedListTraverse()
	}

	fa = a
}

func BenchmarkLinkedListRowTraverse(b *testing.B) {

	var a int

	for i := 0; i < b.N; i++ {
		a = RowTraverse()
	}

	fa = a
}

func BenchmarkLinkedListColumnTraverse(b *testing.B) {

	var a int

	for i := 0; i < b.N; i++ {
		a = ColumnTraverse()
	}

	fa = a
}
