package dictionary

import (
	"fmt"
)

// Diccionario es un conjunto de entradas formadas por pares únicos (clave: valor)
type Dictionary[K comparable, V any] struct {
	mapa map[K]V
}

/*********Métodos*********/
// Crea un diccionario
// O(1)
func NewDictionary[K comparable, V any]() Dictionary[K, V] {
	dict := Dictionary[K, V]{mapa: make(map[K]V)}
	return dict
}

// Agrega una nueva entrada al diccionario, si existe pisa el valor para esa entrada
// O(1)
func (dict *Dictionary[K, V]) Put(key K, val V) {
	dict.mapa[key] = val
}

// Remueve una entrada del diccionario
// O(1)
func (dict *Dictionary[K, V]) Remove(key K) bool {
	var exists bool
	_, exists = dict.mapa[key]
	if exists {
		delete(dict.mapa, key)
	}
	return exists
}

// Verifica si existe una entrada para ese valor de clave
// O(1)
func (dict *Dictionary[K, V]) Contains(key K) bool {
	var exists bool
	_, exists = dict.mapa[key]
	return exists
}

// Devuelve el valor para esa clave
// O(1)
func (dict *Dictionary[K, V]) Get(key K) V {
	return dict.mapa[key]
}

// Devuelve la cantidad de entradas del diccionario
// O(1)
func (dict *Dictionary[K, V]) Size() int {
	return len(dict.mapa)
}

// Devuelve un arreglo con las claves del diccionario
// O(n)
func (dict *Dictionary[K, V]) GetKeys() []K {
	var dictKeys []K
	dictKeys = []K{}
	var key K
	for key = range dict.mapa {
		dictKeys = append(dictKeys, key)
	}
	return dictKeys
}

// Devuelve un arreglo con las valores del diccionario
// O(n)
func (dict *Dictionary[K, V]) GetValues() []V {
	var dictValues []V
	dictValues = []V{}
	var key K
	for key = range dict.mapa {
		dictValues = append(dictValues, dict.mapa[key])
	}
	return dictValues
}

// Devuelve una representación en cadena del diccionario
// O(n)
func (dict Dictionary[K, V]) String() string {
	var str string
	str = "Dictionary {\n"
	var key K
	for key = range dict.mapa {
		str += fmt.Sprintf("\t%v: %v\n", key, dict.mapa[key])
	}
	str += "}"
	return str
}
