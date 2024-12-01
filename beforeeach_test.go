package books_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"example.com/books"
)

// Ginkgo requires, by default, that specs be fully independent. This allows Ginkgo to shuffle the order of specs and run specs in parallel.
var _ = Describe("Using BeforeEach nodes", func() {
	// variables needs to be declared in container node in order to be visible from subject nodes.
	var book *books.Book

	// removing duplication by delcaring a book 'before each' subject node.
	// This node is executed for each subject node.
	BeforeEach(func() {
		book = &books.Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  2783,
		}
		Expect(book.IsValid()).To(BeTrue())
	})

	It("can extract the author's last name", func() {
		Expect(book.AuthorLastName()).To(Equal("Hugo"))
	})

	It("interprets a single author name as a last name", func() {
		book.Author = "Hugo"
		Expect(book.AuthorLastName()).To(Equal("Hugo"))
	})

	It("can extract the author's first name", func() {
		Expect(book.AuthorFirstName()).To(Equal("Victor"))
	})

	It("returns no first name when there is a single author name", func() {
		book.Author = "Hugo"
		Expect(book.AuthorFirstName()).To(BeZero()) //BeZero asserts the value is the zero-value for its type.  In this case: ""
	})
})
