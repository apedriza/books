package books_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"example.com/books"
)

// NOTE: Remember that Describe, Context, and When are functionally equivalent aliases.

var _ = Describe("Organizing in container nodes", func() {
	var book *books.Book

	BeforeEach(func() {
		book = &books.Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  2783,
		}
		Expect(book.IsValid()).To(BeTrue())
	})

	Describe("Extracting the author's first and last name", func() {
		Context("When the author has both names", func() {
			It("can extract the author's last name", func() {
				Expect(book.AuthorLastName()).To(Equal("Hugo"))
			})

			It("can extract the author's first name", func() {
				Expect(book.AuthorFirstName()).To(Equal("Victor"))
			})
		})

		Context("When the author only has one name", func() {
			BeforeEach(func() {
				book.Author = "Hugo"
			})

			It("interprets the single author name as a last name", func() {
				Expect(book.AuthorLastName()).To(Equal("Hugo"))
			})

			It("returns empty for the first name", func() {
				Expect(book.AuthorFirstName()).To(BeZero())
			})
		})

		Context("When the author has a middle name", func() {
			BeforeEach(func() {
				book.Author = "Victor Marie Hugo"
			})

			It("can extract the author's last name", func() {
				Expect(book.AuthorLastName()).To(Equal("Hugo"))
			})

			It("can extract the author's first name", func() {
				Expect(book.AuthorFirstName()).To(Equal("Victor"))
			})
		})

		Context("When the author has no name", func() {
			It("should not be a valid book and returns empty for first and last name", func() {
				book.Author = ""
				Expect(book.IsValid()).To(BeFalse())
				Expect(book.AuthorLastName()).To(BeZero())
				Expect(book.AuthorFirstName()).To(BeZero())
			})
		})
	})

	Describe("JSON encoding and decoding", func() {
		It("survives the round trip", func() {
			encoded, err := book.AsJSON()
			Expect(err).NotTo(HaveOccurred())

			decoded, err := books.NewBookFromJSON(encoded)
			Expect(err).NotTo(HaveOccurred())

			Expect(decoded).To(Equal(book))
		})

		Describe("some JSON decoding edge cases", func() {
			var err error

			When("the JSON fails to parse", func() {
				BeforeEach(func() {
					book, err = books.NewBookFromJSON(`{
				"title":"Les Miserables",
				"author":"Victor Hugo",
				"pages":2783oops
			  }`)
				})

				It("returns a nil book", func() {
					Expect(book).To(BeNil())
				})

				It("errors", func() {
					Expect(err).To(MatchError(books.ErrInvalidJSON))
				})
			})

			When("the JSON is incomplete", func() {
				BeforeEach(func() {
					book, err = books.NewBookFromJSON(`{
				"title":"Les Miserables",
				"author":"Victor Hugo"
			  }`)
				})

				It("returns a nil book", func() {
					Expect(book).To(BeNil())
				})

				It("errors", func() {
					Expect(err).To(MatchError(books.ErrIncompleteJSON))
				})
			})
		})
	})
})
