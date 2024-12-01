package books_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBooks(t *testing.T) {
	// Glue code to connect Ginkgo to Gomega. This line tells
	// to our matcher library (Gomega) which function to call
	// (Ginkgo's Fail) in the event a failure is detected.
	RegisterFailHandler(Fail)

	// Tells Ginkgo to start the test suite, passing it the
	// *testing.T instance and a description of the suite
	RunSpecs(t, "Testing ginkgo examples")
}
