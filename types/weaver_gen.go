// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package types

import (
	"fmt"
	"github.com/ServiceWeaver/onlineboutique/cartservice"
	"github.com/ServiceWeaver/onlineboutique/shippingservice"
	"github.com/ServiceWeaver/onlineboutique/types/money"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
)

// weaver.InstanceOf checks.

// weaver.Router checks.

// Local stub implementations.

// Client stub implementations.

// Note that "weaver generate" will always generate the error message below.
// Everything is okay. The error message is only relevant if you see it when
// you run "go build" or "go run".
var _ codegen.LatestVersion = codegen.Version[[0][24]struct{}](`

ERROR: You generated this file with 'weaver generate' v0.23.2 (codegen
version v0.24.0). The generated code is incompatible with the version of the
github.com/ServiceWeaver/weaver module that you're using. The weaver module
version can be found in your go.mod file or by running the following command.

    go list -m github.com/ServiceWeaver/weaver

We recommend updating the weaver module and the 'weaver generate' command by
running the following.

    go get github.com/ServiceWeaver/weaver@latest
    go install github.com/ServiceWeaver/weaver/cmd/weaver@latest

Then, re-run 'weaver generate' and re-build your code. If the problem persists,
please file an issue at https://github.com/ServiceWeaver/weaver/issues.

`)

// Server stub implementations.

// Reflect stub implementations.

// AutoMarshal implementations.

var _ codegen.AutoMarshal = (*Order)(nil)

type __is_Order[T ~struct {
	weaver.AutoMarshal
	OrderID            string
	ShippingTrackingID string
	ShippingCost       money.T
	ShippingAddress    shippingservice.Address
	Items              []OrderItem
}] struct{}

var _ __is_Order[Order]

func (x *Order) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("Order.WeaverMarshal: nil receiver"))
	}
	enc.String(x.OrderID)
	enc.String(x.ShippingTrackingID)
	(x.ShippingCost).WeaverMarshal(enc)
	(x.ShippingAddress).WeaverMarshal(enc)
	serviceweaver_enc_slice_OrderItem_7622e708(enc, x.Items)
}

func (x *Order) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("Order.WeaverUnmarshal: nil receiver"))
	}
	x.OrderID = dec.String()
	x.ShippingTrackingID = dec.String()
	(&x.ShippingCost).WeaverUnmarshal(dec)
	(&x.ShippingAddress).WeaverUnmarshal(dec)
	x.Items = serviceweaver_dec_slice_OrderItem_7622e708(dec)
}

func serviceweaver_enc_slice_OrderItem_7622e708(enc *codegen.Encoder, arg []OrderItem) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		(arg[i]).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_slice_OrderItem_7622e708(dec *codegen.Decoder) []OrderItem {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]OrderItem, n)
	for i := 0; i < n; i++ {
		(&res[i]).WeaverUnmarshal(dec)
	}
	return res
}

var _ codegen.AutoMarshal = (*OrderItem)(nil)

type __is_OrderItem[T ~struct {
	weaver.AutoMarshal
	Item cartservice.CartItem
	Cost money.T
}] struct{}

var _ __is_OrderItem[OrderItem]

func (x *OrderItem) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("OrderItem.WeaverMarshal: nil receiver"))
	}
	(x.Item).WeaverMarshal(enc)
	(x.Cost).WeaverMarshal(enc)
}

func (x *OrderItem) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("OrderItem.WeaverUnmarshal: nil receiver"))
	}
	(&x.Item).WeaverUnmarshal(dec)
	(&x.Cost).WeaverUnmarshal(dec)
}
