package linkedlist

import (
	"fmt"
)

// node es el nodo de la lista enlazada
// contiene un valor y un puntero al siguiente nodo
// el valor es de tipo genérico, comparable
type node[T any] struct {
	value T
	next  *node[T]
}

// newNode crea un nuevo nodo, con el valor recibido
// y el puntero al siguiente nodo en nil
func newNode[T comparable](value T) *node[T] {
	return &node[T]{value: value, next: nil}
}

/************************************************************/

// LinkedList es la lista enlazada simple
// contiene punteros al primer nodo y al último
type LinkedList[T comparable] struct {
	head *node[T] // puntero al primer nodo
	tail *node[T] // puntero al último nodo
	size int
}

// NewLinkedList crea una nueva lista enlazada, vacía
// En nuestra implementación representamos la lista vacía
// con un puntero al primer nodo en nil
// y un puntero al último nodo en nil
// O(1)
func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{head: nil, tail: nil, size: 0}
}

// InsertAt agrega un nuevo nodo, con el valor recibido,
// en la posición recibida.
// Si la posición es inválida, no hace nada
// La posición 0 agrega el nodo al principio de la lista O(1)
// La posición size agrega el nodo al final de la lista O(1)
// Insertar en otra posición. O(n)
func (l *LinkedList[T]) InsertAt(value T, position int) {
	if position < 0 || position > l.size {
		return
	}
	newNode := newNode(value)
	// Insertar al principio
	// O(1)
	if position == 0 {
		newNode.next = l.head
		l.head = newNode
		if l.tail == nil {
			l.tail = newNode
		}
		l.size++
		return
	}
	// Insertar al final
	// O(1)
	if position == l.size {
		l.tail.next = newNode
		l.tail = newNode
		l.size++
		return
	}
	// Insertar en position
	current := l.head
	for current != nil && position > 1 {
		current = current.next
		position--
	}

	newNode.next = current.next
	current.next = newNode
	l.size++
}

// RemoveAt elimina el nodo en la posición recibida
// Si la posición es inválida, no hace nada
// La posición 0 elimina el primer nodo de la lista O(1)
// Eliminar en otra posición. O(n)
func (l *LinkedList[T]) RemoveAt(position int) {
	if position < 0 || position >= l.size {
		return
	}
	// Eliminar el primer nodo
	// O(1)
	if position == 0 {
		l.head = l.head.next
		l.size--
		return
	}

	current := l.head
	for current != nil && position > 1 {
		current = current.next
		position--
	}

	current.next = current.next.next
	l.size--
}

// Search busca el primer nodo que contenga el valor recibido
// y devuelve su posición en la lista o -1 si no lo encuentra
// O(n)
func (l *LinkedList[T]) Search(value T) int {
	if l.head == nil {
		return -1
	}
	current := l.head
	position := 0
	for current != nil {
		if current.value == value {
			return position
		}
		current = current.next
		position++
	}
	return -1
}

// Get devuelve el valor del nodo en la posición recibida
// o un valor nulo si la posición es inválida
// O(n)
func (l *LinkedList[T]) Get(position int) T {
	if position < 0 || position >= l.size {
		var t T // la variable t se inicializ con un valor nulo
		return t
	}

	current := l.head
	for current != nil && position > 0 {
		current = current.next
		position--
	}

	return current.value
}

// Size devuelve la cantidad de nodos en la lista
// O(n)
func (l *LinkedList[T]) Size() int {
	return l.size
}

// Contains verifica si el valor recibido está en la lista
// O(n)
func (l *LinkedList[T]) Contains(value T) bool {
	return l.Search(value) != -1
}

// IsEmpty verifica si la lista está vacía
// O(1)
func (l *LinkedList[T]) IsEmpty() bool {
	return l.head == nil
}

// String devuelve una representación en cadena de la lista
// en el formato [1 2 3].
// Se puede usar para imprimir la lista con fmt.Println
// O(n)
func (l LinkedList[T]) String() string {
	if l.head == nil {
		return "[]"
	}
	current := l.head
	result := "lista: ["
	for current != nil {
		result += fmt.Sprintf("%v", current.value)
		if current.next != nil {
			result += " "
		}
		current = current.next
	}
	result += "]"
	return result
}


// Ejercicio 3 Ma-Mie

func (l *LinkedList[T]) Revertir() {

	size := l.Size()

	for i := 0 ; i < size-1; i++ {
		
		aux, pos := l.Get(0), l.Size() - i

		l.InsertAt(aux, pos)
		l.RemoveAt(0)
	}
}
