package supermarket_test

import (
	. "github.com/ollie-troward/code-katas/01-supermarket-pricing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Supermarket", func() {
	Describe("GetTotal", func() {
		Context("Passing in a valid product", func() {
			It("returns the expected price", func() {
				beans := Beans{}

				supermarket := Supermarket{
					Products: []Product{beans},
				}

				Expect(supermarket.GetTotal()).To(Equal(beans.GetPrice()))
			})
		})
		Context("Passing in multiple products", func() {
			It("returns the total of prices", func() {
				firstTinOfBeans := Beans{}
				secondTinOfBeans := Beans{}

				supermarket := Supermarket{
					Products: []Product{firstTinOfBeans, secondTinOfBeans},
				}

				Expect(supermarket.GetTotal()).To(Equal(firstTinOfBeans.GetPrice() + secondTinOfBeans.GetPrice()))
			})
		})
	})
})
