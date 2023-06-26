# Recuperatorio Parcial 1 - Martes y Miércoles

1. Implementar un TAD Cola que internamente utiliza dos pilas para contener los datos de la siguiente manera: cada vez que se encola un elemento se lo guarda en la pila A, y cada vez que se desencola una elemento se lo extrae del tope de la pila B, si la pila B está vacía entonces se desapilan todos los elementos de la pila A y se los apila en B. Los métodos que debe soportar son Encolar, Desencolar, EstaVacia y Frente (para ver el primer elemento de la Cola)

2. Agregar a la implementación de lista enlazada simple un método denominado Revertir, que revierta los elementos de la lista, tal que el primer nodo pase a ser el último y viceversa Justificar el orden de la solución

        func (l *LinkedList[T]) Revertir()

3. Implementar un TAD denominado Glosario que permita asociar una palabra con varias definiciones. El Glosario debe proveer los métodos Agregar que permite agregarles definiciones a una palabra dada y un método Buscar que dada una palabra muestre por pantalla todas las definiciones de la misma.

4. Escribir una función que devuelva el elemento máximo de una lista de enteros que utilice la técnica de división y conquista. Calcular el orden de la solución aplicando el Teorema del Maestro.

