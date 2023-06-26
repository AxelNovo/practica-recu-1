# Parcial 1 - Martes y Jueves

1. Dada la implementación de Set vista en clase (implementado sobre una lista enlazada simple) que tiene los métodos: Contains, Add, Remove, Size, String y Values y soporta las siguientes operaciones que en todos los casos devuelven un conjunto nuevo: NewSet, Union, Intersection, Difference, Subset y Equal. Se pide reescribir la operación Union para que reciba como parámetro un número indefinido de conjuntos y devuelva la Union de todos ellos.

        func Union[T comparable](conjuntos ...*Set[T]) *Set[T]

2. Escribir una función que recibe un diccionario cuyas claves son los nombres de los alumnos y como valor una lista con los días que asistieron a clase. Debe devolver un diccionario con clave fecha y valor la lista de alumnos que asistieron en dicha fecha. 

    Por ejemplo, si la entrada es {"Ana": [ "Mie 10", "Vie 12"], "Luz": [ "Vie 12",  "Mie 17"],  "Pedro": ["Mie  10", "Mie 17"]}, debe devolver  {“Mie 10”: [“Ana", "Pedro"]), “Vie 12”: [“Ana", "Luz”], “Mie 17”: [“Luz”, "Pedro"]}

3. Dada una tabla de hash de tamaño 7 y la función de hashing h(x) = x mod 7, inserte los números 1, 8, 27, 125, 216, 343, resolviendo los choques con el método de hashing cerrado. Dibuje las tablas resultantes en cada paso.

4. Aplicar el teorema del maestro para calcular el orden de la siguiente llamada a la función SubsecuenciaSumaMaxima(arreglo, 0, len(arreglo)-1)  que utiliza la técnica de división y conquista.

---

### Todos los ejercicios para codificar están en main.go.

A continuación, la resolución del ejercicio 2, en donde se escribe el paso a paso.

Primero, vamos a inicializar una tabla de hash de tamaño 7 con todas las posiciones vacías:

```
Posición: 0 1 2 3 4 5 6
Valor:    - - - - - - -
```

Donde "-" indica una posición vacía.

Ahora, vamos a insertar los números uno por uno usando la función de hashing h(x) = x mod 7.

1. Insertamos 1: h(1) = 1 mod 7 = 1. Así que insertamos 1 en la posición 1.

```
Posición: 0 1 2 3 4 5 6
Valor:    - 1 - - - - -
```

2. Insertamos 8: h(8) = 8 mod 7 = 1. Pero la posición 1 ya está ocupada, así que vamos a la siguiente posición disponible, que es la posición 2.

```
Posición: 0 1 2 3 4 5 6
Valor:    - 1 8 - - - -
```

3. Insertamos 27: h(27) = 27 mod 7 = 6. Insertamos 27 en la posición 6.

```
Posición: 0 1 2 3 4 5 6
Valor:    - 1 8 - - - 27
```

4. Insertamos 125: h(125) = 125 mod 7 = 6. Pero la posición 6 ya está ocupada, así que vamos a la siguiente posición disponible, que es la posición 0.

```
Posición: 0 1 2 3 4 5 6
Valor:    125 1 8 - - - 27
```

5. Insertamos 216: h(216) = 216 mod 7 = 6. Pero las posiciones 6 y 0 ya están ocupadas, así que vamos a la siguiente posición disponible, que es la posición 3.

```
Posición: 0 1 2 3 4 5 6
Valor:    125 1 8 216 - - 27
```

6. Insertamos 343: h(343) = 343 mod 7 = 0. Pero las posiciones 0, 1, 2 y 3 ya están ocupadas, así que vamos a la siguiente posición disponible, que es la posición 4.

```
Posición: 0 1 2 3 4 5 6
Valor:    125 1 8 216 343 - 27
```

Así es como se vería la tabla de hash después de insertar todos los números.