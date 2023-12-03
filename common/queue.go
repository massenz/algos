/*
 * Copyright (c) 2023 Marco Massenzio. All rights reserved.
 */

package common

import (
	"errors"
)

// TODO: Should use generics

var (
	QueueFullError = errors.New("queue is full")
)

// A Queue is a basic FIFO queue based on a list that has a fixed capacity.
type Queue struct {
	Capacity int
	Items    []interface{}
}

func NewQueue(capacity int) *Queue {
	return &Queue{Capacity: capacity, Items: make([]interface{}, 0, capacity)}
}

func (q *Queue) Enqueue(item interface{}) error {
	if q.IsFull() {
		return QueueFullError
	}
	q.Items = append(q.Items, item)
	return nil
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	item := q.Items[0]
	q.Items = q.Items[1:]
	return item
}

func (q *Queue) Peek() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.Items[0]
}

func (q *Queue) IsEmpty() bool {
	return len(q.Items) == 0
}

func (q *Queue) IsFull() bool {
	return len(q.Items) == q.Capacity
}
