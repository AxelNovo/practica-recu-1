package main

import (
	"fmt"
	"utils/set"
	"utils/dictionary"
)

func main() {

	s1 := set.NewSet(1, 2, 3, 4)
	s2 := set.NewSet(5, 6, 7, 8)
	s3 := set.NewSet(9, 10, 11, 12)

	union := Union(s1, s2, s3)

	fmt.Println(union)

	d := dictionary.NewDictionary[string, []string]()

	d.Put("Ana", []string{"Mie 10", "Vie 12"})
	d.Put("Luz", []string{"Vie 12", "Mie 17"})
	d.Put("Pedro", []string{"Mie 10", "Mie 17"})

	invertido := Invertir(d)

	fmt.Println(invertido)
	
}

// Ejercicio 1

func Union[T comparable](conjuntos ...*set.Set[T]) *set.Set[T] {

	newSet := set.NewSet[T]()

	for i := 0; i < len(conjuntos); i++ {
		
		for _, v := range conjuntos[i].Values() {
			
			newSet.Add(v)
		}
	}

	return newSet
}

// Ejercicio 2

func Invertir(diccionario dictionary.Dictionary[string, []string]) dictionary.Dictionary[string, []string]{

	newDictionary := dictionary.NewDictionary[string, []string]()
	
	for _, key := range diccionario.GetKeys() {
		
		for _, value := range diccionario.Get(key) {
			
			if !newDictionary.Contains(value){

				newDictionary.Put(value, []string{key})
			} else {

				aux := newDictionary.Get(value)
				aux = append(aux, key)

				newDictionary.Put(value, aux)
			}
		}
	}

	return newDictionary
}
