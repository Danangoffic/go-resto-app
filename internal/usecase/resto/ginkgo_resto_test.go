package resto_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"resto-app/internal/model"
	"resto-app/internal/usecase/resto"
)

var _ = Describe("GinkgoResto", func() {
	var usecase resto.Usecase
	Describe("request order info", func() {
		Context("it gave the correct inputs", func() {
			inputs := model.GetOrderDataRequest{
				OrderID: "valid_order_id",
				UserID:  "valid_user_id",
			}

			When("the request order id is not the user's", func() {
				It("returns unqauthorized error", func() {
					res, err := usecase.GetOrderData(context.Background(), inputs)
					Expect(err).Should(HaveOccurred())
					Expect(err.Error()).To(BeEquivalentTo("unauthorized"))
					Expect(res).To(BeEquivalentTo(model.Order{}))
				})
			})
		})
	})
})
