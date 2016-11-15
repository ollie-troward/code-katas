package supermarket_test

import (
	. "github.com/ollie-troward/code-katas/01-supermarket-pricing"

	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Beans", func() {
	Describe("GetPrice", func() {
		Context("Retrieve the price of beans", func() {
			It("Retrieves the correct price for beans", func() {
				var priceOfBeans float32
				priceOfBeans = 0.65

				Expect(Beans{}.GetPrice()).
					To(Equal(priceOfBeans), fmt.Sprintf("Expected beans to cost %s", priceOfBeans))
			})
		})
	})
})
