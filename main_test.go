package main

import (
	"context"
	"testing"
	"time"

	"github.com/ServiceWeaver/onlineboutique/cartservice"
	"github.com/ServiceWeaver/onlineboutique/checkoutservice"
	"github.com/ServiceWeaver/onlineboutique/paymentservice"
	"github.com/ServiceWeaver/onlineboutique/productcatalogservice"
	"github.com/ServiceWeaver/onlineboutique/shippingservice"
	"github.com/ServiceWeaver/weaver/weavertest"
)

func TestPurchase(t *testing.T) {
	runner := weavertest.Local
	runner.Test(t, func(
		t *testing.T,
		catalog productcatalogservice.ProductCatalogService,
		cart cartservice.CartService,
		checkout checkoutservice.CheckoutService,
	) {
		// List all products.
		ctx := context.Background()
		products, err := catalog.ListProducts(ctx)
		if err != nil {
			t.Fatal(err)
		}

		// Add one of every product to our cart.
		const userID = "sundar_pichai"
		for _, product := range products {
			item := cartservice.CartItem{ProductID: product.ID, Quantity: 1}
			if err := cart.AddItem(ctx, userID, item); err != nil {
				t.Fatal(err)
			}
		}

		// Place the order.
		order := checkoutservice.PlaceOrderRequest{
			UserID:       userID,
			UserCurrency: "USD",
			Address: shippingservice.Address{
				StreetAddress: "1600 Amphitheatre Parkway",
				City:          "Mountain View",
				State:         "CA",
				Country:       "USA",
				ZipCode:       94043,
			},
			Email: "sundar@google.com",
			CreditCard: paymentservice.CreditCardInfo{
				Number:          "4432-8015-6152-0454",
				CVV:             672,
				ExpirationYear:  2025,
				ExpirationMonth: time.January,
			},
		}
		if _, err := checkout.PlaceOrder(ctx, order); err != nil {
			t.Fatal(err)
		}
	})
}
