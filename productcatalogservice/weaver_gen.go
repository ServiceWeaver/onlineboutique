// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package productcatalogservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/ServiceWeaver/onlineboutique/types/money"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:  "github.com/ServiceWeaver/onlineboutique/productcatalogservice/T",
		Iface: reflect.TypeOf((*T)(nil)).Elem(),
		Impl:  reflect.TypeOf(impl{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return t_local_stub{impl: impl.(T), tracer: tracer, getProductMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/onlineboutique/productcatalogservice/T", Method: "GetProduct", Remote: false}), listProductsMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/onlineboutique/productcatalogservice/T", Method: "ListProducts", Remote: false}), searchProductsMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/onlineboutique/productcatalogservice/T", Method: "SearchProducts", Remote: false})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return t_client_stub{stub: stub, getProductMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/onlineboutique/productcatalogservice/T", Method: "GetProduct", Remote: true}), listProductsMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/onlineboutique/productcatalogservice/T", Method: "ListProducts", Remote: true}), searchProductsMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/onlineboutique/productcatalogservice/T", Method: "SearchProducts", Remote: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return t_server_stub{impl: impl.(T), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return t_reflect_stub{caller: caller}
		},
		RefData: "",
	})
}

// weaver.InstanceOf checks.
var _ weaver.InstanceOf[T] = (*impl)(nil)

// weaver.Router checks.
var _ weaver.Unrouted = (*impl)(nil)

// Local stub implementations.

type t_local_stub struct {
	impl                  T
	tracer                trace.Tracer
	getProductMetrics     *codegen.MethodMetrics
	listProductsMetrics   *codegen.MethodMetrics
	searchProductsMetrics *codegen.MethodMetrics
}

// Check that t_local_stub implements the T interface.
var _ T = (*t_local_stub)(nil)

func (s t_local_stub) GetProduct(ctx context.Context, a0 string) (r0 Product, err error) {
	// Update metrics.
	begin := s.getProductMetrics.Begin()
	defer func() { s.getProductMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "productcatalogservice.T.GetProduct", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetProduct(ctx, a0)
}

func (s t_local_stub) ListProducts(ctx context.Context) (r0 []Product, err error) {
	// Update metrics.
	begin := s.listProductsMetrics.Begin()
	defer func() { s.listProductsMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "productcatalogservice.T.ListProducts", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.ListProducts(ctx)
}

func (s t_local_stub) SearchProducts(ctx context.Context, a0 string) (r0 []Product, err error) {
	// Update metrics.
	begin := s.searchProductsMetrics.Begin()
	defer func() { s.searchProductsMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "productcatalogservice.T.SearchProducts", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.SearchProducts(ctx, a0)
}

// Client stub implementations.

type t_client_stub struct {
	stub                  codegen.Stub
	getProductMetrics     *codegen.MethodMetrics
	listProductsMetrics   *codegen.MethodMetrics
	searchProductsMetrics *codegen.MethodMetrics
}

// Check that t_client_stub implements the T interface.
var _ T = (*t_client_stub)(nil)

func (s t_client_stub) GetProduct(ctx context.Context, a0 string) (r0 Product, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.getProductMetrics.Begin()
	defer func() { s.getProductMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "productcatalogservice.T.GetProduct", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	(&r0).WeaverUnmarshal(dec)
	err = dec.Error()
	return
}

func (s t_client_stub) ListProducts(ctx context.Context) (r0 []Product, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.listProductsMetrics.Begin()
	defer func() { s.listProductsMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "productcatalogservice.T.ListProducts", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	var shardKey uint64

	// Call the remote method.
	var results []byte
	results, err = s.stub.Run(ctx, 1, nil, shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_Product_2bf2ec03(dec)
	err = dec.Error()
	return
}

func (s t_client_stub) SearchProducts(ctx context.Context, a0 string) (r0 []Product, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.searchProductsMetrics.Begin()
	defer func() { s.searchProductsMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "productcatalogservice.T.SearchProducts", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 2, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_Product_2bf2ec03(dec)
	err = dec.Error()
	return
}

// Note that "weaver generate" will always generate the error message below.
// Everything is okay. The error message is only relevant if you see it when
// you run "go build" or "go run".
var _ codegen.LatestVersion = codegen.Version[[0][20]struct{}](`

ERROR: You generated this file with 'weaver generate' v0.22.0 (codegen
version v0.20.0). The generated code is incompatible with the version of the
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

type t_server_stub struct {
	impl    T
	addLoad func(key uint64, load float64)
}

// Check that t_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*t_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s t_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "GetProduct":
		return s.getProduct
	case "ListProducts":
		return s.listProducts
	case "SearchProducts":
		return s.searchProducts
	default:
		return nil
	}
}

func (s t_server_stub) getProduct(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.GetProduct(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	(r0).WeaverMarshal(enc)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) listProducts(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.ListProducts(ctx)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_Product_2bf2ec03(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) searchProducts(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.SearchProducts(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_Product_2bf2ec03(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

// Reflect stub implementations.

type t_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that t_reflect_stub implements the T interface.
var _ T = (*t_reflect_stub)(nil)

func (s t_reflect_stub) GetProduct(ctx context.Context, a0 string) (r0 Product, err error) {
	err = s.caller("GetProduct", ctx, []any{a0}, []any{&r0})
	return
}

func (s t_reflect_stub) ListProducts(ctx context.Context) (r0 []Product, err error) {
	err = s.caller("ListProducts", ctx, []any{}, []any{&r0})
	return
}

func (s t_reflect_stub) SearchProducts(ctx context.Context, a0 string) (r0 []Product, err error) {
	err = s.caller("SearchProducts", ctx, []any{a0}, []any{&r0})
	return
}

// AutoMarshal implementations.

var _ codegen.AutoMarshal = (*Product)(nil)

type __is_Product[T ~struct {
	weaver.AutoMarshal
	ID          string   "json:\"id\""
	Name        string   "json:\"name\""
	Description string   "json:\"description\""
	Picture     string   "json:\"picture\""
	PriceUSD    money.T  "json:\"priceUsd\""
	Categories  []string "json:\"categories\""
}] struct{}

var _ __is_Product[Product]

func (x *Product) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("Product.WeaverMarshal: nil receiver"))
	}
	enc.String(x.ID)
	enc.String(x.Name)
	enc.String(x.Description)
	enc.String(x.Picture)
	(x.PriceUSD).WeaverMarshal(enc)
	serviceweaver_enc_slice_string_4af10117(enc, x.Categories)
}

func (x *Product) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("Product.WeaverUnmarshal: nil receiver"))
	}
	x.ID = dec.String()
	x.Name = dec.String()
	x.Description = dec.String()
	x.Picture = dec.String()
	(&x.PriceUSD).WeaverUnmarshal(dec)
	x.Categories = serviceweaver_dec_slice_string_4af10117(dec)
}

func serviceweaver_enc_slice_string_4af10117(enc *codegen.Encoder, arg []string) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		enc.String(arg[i])
	}
}

func serviceweaver_dec_slice_string_4af10117(dec *codegen.Decoder) []string {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]string, n)
	for i := 0; i < n; i++ {
		res[i] = dec.String()
	}
	return res
}

// Encoding/decoding implementations.

func serviceweaver_enc_slice_Product_2bf2ec03(enc *codegen.Encoder, arg []Product) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		(arg[i]).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_slice_Product_2bf2ec03(dec *codegen.Decoder) []Product {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]Product, n)
	for i := 0; i < n; i++ {
		(&res[i]).WeaverUnmarshal(dec)
	}
	return res
}
