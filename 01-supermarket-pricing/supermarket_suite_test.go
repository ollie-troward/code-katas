package supermarket_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSupermarket(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Supermarket Suite")
}
