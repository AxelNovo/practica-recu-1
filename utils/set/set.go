package set

import (
	"utils/linkedlist"
	"strings"
)

// Set es un conjunto de elementos comparables
type Set[T comparable] struct {
	elementos linkedlist.LinkedList[T]
}

/*********Métodos*********/
// Contains verifica si el elemento pertenece al conjunto
// O(n)
func (s *Set[T]) Contains(element T) bool {
	return s.elementos.Contains(element)
}

// Add agrega un elemento al conjunto
// Si el elemento ya pertenece al conjunto, no hace nada
// O(n) (en el peor de los casos, ya que debe verificar
// si el elemento ya está en el conjunto)
func (s *Set[T]) Add(element T) {
	if !s.Contains(element) { //O(n)
		s.elementos.InsertAt(element, 0) //O(1)
	}
}

// Remove elimina un elemento del conjunto
// Si el elemento no pertenece al conjunto, no hace nada
// O(n)
func (s *Set[T]) Remove(element T) {
	s.elementos.RemoveAt(s.elementos.Search(element))
}

// Size devuelve la cantidad de elementos del conjunto
// O(1)
func (s *Set[T]) Size() int {
	return s.elementos.Size()
}

// String devuelve una representación en cadena del conjunto
// O(n)
func (s Set[T]) String() string {
	cadena := "Conjunto: {"
	elementos := s.elementos.String()
	inicio := strings.Index(elementos, "[") + 1
	fin := strings.Index(elementos, "]")
	cadena += elementos[inicio:fin]
	cadena += "}"
	return cadena
}

// Values devuelve un arreglo con los elementos del conjunto
func (s *Set[T]) Values() []T {
	var values []T
	for i := 0; i < s.Size(); i++ {
		values = append(values, s.elementos.Get(i))
	}
	return values
}

/*********Operaciones sobre conjuntos *********/

// NewSet crea un nuevo conjunto con los elementos que recibe como parámetros
// Si no recibe ningún elemento, crea un conjunto vacío
// O(n) donde n es la cantidad de elementos
func NewSet[T comparable](elementos ...T) *Set[T] {
	conjunto := &Set[T]{*linkedlist.NewLinkedList[T]()}
	for _, elemento := range elementos {
		conjunto.Add(elemento)
	}
	return conjunto
}

// Union devuelve un nuevo conjunto con la unión de los conjuntos recibidos
// O(n*m), donde n y m son los tamaños de los conjuntos s1 y s2 respectivamente
func Union[T comparable](s1, s2 *Set[T]) *Set[T] {
	result := NewSet[T]()
	for i := 0; i < s1.Size(); i++ {
		result.Add(s1.elementos.Get(i))
	}
	for i := 0; i < s2.Size(); i++ {
		result.Add(s2.elementos.Get(i))
	}
	return result
}

// Intersection devuelve un nuevo conjunto con la intersección de los conjuntos
// recibidos
// O(n*m) donde n y m son los tamaños de los conjuntos s1 y s2 respectivamente
func Intersection[T comparable](s1, s2 *Set[T]) *Set[T] {
	result := NewSet[T]()
	for i := 0; i < s1.Size(); i++ {
		elemento := s1.elementos.Get(i)
		if s2.Contains(elemento) {
			result.Add(elemento)
		}
	}
	return result
}

// Difference devuelve un nuevo conjunto con la diferencia de los conjuntos
// recibidos
// La diferencia entre s1 y s2, son todos los elementos de s1 que no están en s2
// O(n*m) donde n y m son los tamaños de los conjuntos s1 y s2 respectivamente
func Difference[T comparable](s1, s2 *Set[T]) *Set[T] {
	result := NewSet[T]()
	for i := 0; i < s1.Size(); i++ {
		elemento := s1.elementos.Get(i)
		if !s2.Contains(elemento) {
			result.Add(elemento)
		}
	}
	return result
}

// Subset verifica si el conjunto s2 es subconjunto del conjunto s1
// O(n*m) donde n y m son los tamaños de los conjuntos s1 y s2 respectivamente
func Subset[T comparable](s1, s2 *Set[T]) bool {
	for i := 0; i < s2.Size(); i++ {
		if !s1.Contains(s2.elementos.Get(i)) {
			return false
		}
	}
	return true
}

// Equal verifica si los conjuntos recibidos son iguales
// O(n^2) donde n es el tamaño de los conjuntos s1 y s2
func Equal[T comparable](s1, s2 *Set[T]) bool {
	return Subset(s1, s2) && Subset(s2, s1)
}
