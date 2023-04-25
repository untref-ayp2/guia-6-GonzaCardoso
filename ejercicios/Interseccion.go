package ejercicios

import (
	"guia6/dictionary"
	"guia6/linkedlist"
)

func Interseccion(l1 linkedlist.LinkedList[string], l2 linkedlist.LinkedList[string]) linkedlist.LinkedList[string] {
	dict := dictionary.NewDictionary[string, int]()
	l3 := linkedlist.NewLinkedList[string]()
	repetidos := 0
	for i := 0; i < l1.Size(); i++ {
		dict.Put(l1.Get(i), 1)
	}
	for i := 0; i < l2.Size(); i++ {
		if dict.Contains(l2.Get(i)) {
			l3.InsertAt(l2.Get(i), repetidos)
			repetidos++
		}
	}

	return *l3
}
