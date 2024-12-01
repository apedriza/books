package books_test

import (
	"example.com/books"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("some JSON decoding edge cases", func() {
	var book *books.Book
	var err error
	var json string
	JustBeforeEach(func() {
		book, err = books.NewBookFromJSON(json)
		Expect(book).To(BeNil())
	})

	When("the JSON fails to parse", func() {
		BeforeEach(func() {
			json = `{
		  "title":"Les Miserables",
		  "author":"Victor Hugo",
		  "pages":2783oops
		}`
		})

		It("errors", func() {
			Expect(err).To(MatchError(books.ErrInvalidJSON))
		})
	})

	When("the JSON is incomplete", func() {
		BeforeEach(func() {
			json = `{
		  "title":"Les Miserables",
		  "author":"Victor Hugo"
		}`
		})

		It("errors", func() {
			Expect(err).To(MatchError(books.ErrIncompleteJSON))
		})
	})
})
