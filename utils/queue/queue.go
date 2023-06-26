package queue

import (
	"errors"
	"utils/stack"
)

// Queue implementa una cola genérica sobre un arreglo dinámico.
type Queue struct{
	data []any
}

// Enqueue agrega un elemento a la cola. O(1)
func (q *Queue) Enqueue(v any) {
	q.data = append(q.data, v)
}

// Dequeue saca un elemento de la cola. O(1)
func (q *Queue) Dequeue() (any, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	head := (q.data)[0]
	q.data = (q.data)[1:]
	return head, nil
}

// Front devuelve el elemento del frente de la cola. O(1)
func (q *Queue) Front() (any, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	return (q.data)[0], nil
}

// IsEmpty verifica si la cola esta vacia. O(1)
func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}


type QueueS struct{

	a, b stack.Stack
}

func (q *QueueS) Enqueue(v any) {

    q.a.Push(v)
}

func (q *QueueS) Dequeue() (any, error) {

	if q.IsEmpty(){
		return nil, errors.New("queueS is empty")
	}

    if q.b.IsEmpty(){

		for !q.a.IsEmpty(){

			aux, _ := q.a.Pop()

			q.b.Push(aux)
		}
	}
	aux, _ := q.b.Pop()
	return aux, nil
}

func (q *QueueS) IsEmpty() bool {
    
	return q.a.IsEmpty() && q.b.IsEmpty()
}

func (q *QueueS) Front() (any, error) {
    if q.IsEmpty(){
		return nil, errors.New("queueS is empty")
	}

    if q.b.IsEmpty(){

		for !q.a.IsEmpty(){

			aux, _ := q.a.Pop()

			q.b.Push(aux)
		}
	}
	aux, _ := q.b.Top()
	return aux, nil
}