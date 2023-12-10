/*
 * Copyright (c) 2023 Marco Massenzio. All rights reserved.
 */

package common_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/massenz/algos/common"
)

var _ = Describe("Queue", func() {
	var (
		queue *Queue
	)

	BeforeEach(func() {
		queue = NewQueue(3)
	})

	Describe("Enqueue", func() {
		Context("when the queue is not full", func() {
			It("adds an item to the queue", func() {
				err := queue.Enqueue("item1")
				Expect(err).NotTo(HaveOccurred())
				Expect(queue.Items).To(ContainElement("item1"))
			})
		})

		Context("when the queue is full", func() {
			It("returns an error", func() {
				queue.Enqueue("item1")
				queue.Enqueue("item2")
				queue.Enqueue("item3")
				err := queue.Enqueue("item4")
				Expect(err).To(MatchError(QueueFullError))
			})
		})
	})

	Describe("Dequeue", func() {
		Context("when the queue is not empty", func() {
			It("removes and returns the first item from the queue", func() {
				queue.Enqueue("item1")
				queue.Enqueue("item2")
				item := queue.Dequeue()
				Expect(item).To(Equal("item1"))
				Expect(queue.Items).NotTo(ContainElement("item1"))
			})
		})

		Context("when the queue is empty", func() {
			It("returns nil", func() {
				item := queue.Dequeue()
				Expect(item).To(BeNil())
			})
		})
	})

	Describe("Peek", func() {
		Context("when the queue is not empty", func() {
			It("returns the first item from the queue without removing it", func() {
				queue.Enqueue("item1")
				item := queue.Peek()
				Expect(item).To(Equal("item1"))
				Expect(queue.Items).To(ContainElement("item1"))
			})
		})

		Context("when the queue is empty", func() {
			It("returns nil", func() {
				item := queue.Peek()
				Expect(item).To(BeNil())
			})
		})
	})

	Describe("IsEmpty", func() {
		Context("when the queue is empty", func() {
			It("returns true", func() {
				Expect(queue.IsEmpty()).To(BeTrue())
			})
		})

		Context("when the queue is not empty", func() {
			It("returns false", func() {
				queue.Enqueue("item1")
				Expect(queue.IsEmpty()).To(BeFalse())
			})
		})
	})

	Describe("IsFull", func() {
		Context("when the queue is full", func() {
			It("returns true", func() {
				queue.Enqueue("item1")
				queue.Enqueue("item2")
				queue.Enqueue("item3")
				Expect(queue.IsFull()).To(BeTrue())
			})
		})

		Context("when the queue is not full", func() {
			It("returns false", func() {
				queue.Enqueue("item1")
				Expect(queue.IsFull()).To(BeFalse())
			})
		})
	})
})
