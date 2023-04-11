package tests

import (
	"guia6/ejercicios"
	"guia6/linkedlist"
	"testing"
)

func TestInterseccion(t *testing.T) {
	l1:= linkedlist.NewLinkedList[string]()
	l1.InsertAt("A",0)
	l1.InsertAt("B",1)
	l1.InsertAt("C",2)
	l2:= linkedlist.NewLinkedList[string]()
	l2.InsertAt("A",0)
	l2.InsertAt("E",1)
	l2.InsertAt("F",2)
	lr := ejercicios.Interseccion(*l1, *l2)
	if lr.Size() != 1 {
		t.Errorf("Tamaño %d, debería ser %d", lr.Size(), 1)
	}
	if !lr.Contains("A"){
		t.Errorf("Contains(A) = %t, debería ser %t", lr.Contains("A"), true)
	}
}
