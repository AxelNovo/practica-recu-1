package main

import (
	"errors"
	"fmt"
	"utils/dictionary"
	"utils/linkedlist"
	"utils/stack"
)

func main() {

	// Ejercicio 2

	l := linkedlist.NewLinkedList[int]()
	l.InsertAt(1, l.Size())
	l.InsertAt(2, l.Size())
	l.InsertAt(3, l.Size())
	l.InsertAt(4, l.Size())
	l.InsertAt(5, l.Size())

	fmt.Println(l)

	l.Revertir()

	fmt.Println(l)

	// Ejercicio 1

	c := Cola{}

	c.Encolar(1)
	c.Encolar(2)
	c.Encolar(3)

	fmt.Println(c)

	v, _ := c.Desencolar()

	fmt.Println(v)

	fmt.Println(c)

	v, _ = c.Desencolar()

	fmt.Println(v)

	fmt.Println(c)

	// Ejercicio 3

	g := NewGlossary()

	g.Add("Hola", "Expresión con que se saluda.")
	g.Add("Mundo", "Conjunto de todas las cosas que existen y de la humanidad.")

	fmt.Println(g)

	g.Search("Hola")
	g.Search("Mundo")
	g.Search("Foo")

	g.Add("Hola", "Expresión con la que se atiende a una llamada telefónica")

	g.Search("Hola")

	// Ejercicio 4

	list := []int{10, 12, 23, 4, 56, 71, 89, 2, 3, 6}
	fmt.Println(BuscarMaximo(list))

}

// Ejercicio 1

type Cola struct {
	a, b stack.Stack
}

func (c *Cola) Encolar(v any) {

	c.a.Push(v)
}

func (c *Cola) Desencolar() (any, error) {

	if c.EstaVacia() {

		return nil, errors.New("La cola está vacía.")
	}

	if c.b.IsEmpty() {

		for !c.a.IsEmpty() {

			aux, _ := c.a.Pop()

			c.b.Push(aux)
		}
	}

	v, _ := c.b.Pop()

	return v, nil
}

func (c *Cola) EstaVacia() bool {

	return c.a.IsEmpty() && c.b.IsEmpty()
}

func (c *Cola) Frente() (any, error) {

	if c.EstaVacia() {

		return nil, errors.New("La cola está vacía.")
	}

	if c.b.IsEmpty() {

		for !c.a.IsEmpty() {

			aux, _ := c.a.Pop()

			c.b.Push(aux)
		}
	}

	v, _ := c.b.Top()

	return v, nil
}

// Ejercicio 2 en utils/linkedlist.go

// Ejercicio 3

type Glossary struct {
	dict dictionary.Dictionary[string, []string]
}

func NewGlossary() Glossary {

	glos := Glossary{dict: dictionary.NewDictionary[string, []string]()}

	return glos
}

func (g *Glossary) Add(word string, definition string) {

	if !g.dict.Contains(word) {

		g.dict.Put(word, []string{definition})

		return
	}

	v := g.dict.Get(word)
	v = append(v, definition)

	g.dict.Put(word, v)
}

func (g *Glossary) Search(word string) {

	if !g.dict.Contains(word) {

		fmt.Println("Palabra no encontrada.")

		return
	}

	definitions := g.dict.Get(word)

	for _, def := range definitions {

		fmt.Println(def)
	}
}

// Ejercicio 4

func BuscarMaximo(lista []int) int {

	if len(lista) == 1 {

		return lista[0]
	}

	mitad := len(lista) / 2

	max1 := BuscarMaximo(lista[:mitad])
	max2 := BuscarMaximo(lista[mitad:])

	if max1 > max2 {

		return max1
	} else {

		return max2
	}

}
