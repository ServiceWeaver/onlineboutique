// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package currencyservice

import (
	"context"
	"errors"
	"github.com/ServiceWeaver/onlineboutique/types/money"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:  "github.com/ServiceWeaver/onlineboutique/currencyservice/CurrencyService",
		Iface: reflect.TypeOf((*CurrencyService)(nil)).Elem(),
		Impl:  reflect.TypeOf(impl{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return currencyService_local_stub{impl: impl.(CurrencyService), tracer: tracer, convertMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/onlineboutique/currencyservice/CurrencyService", Method: "Convert", Remote: false}), getSupportedCurrenciesMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/onlineboutique/currencyservice/CurrencyService", Method: "GetSupportedCurrencies", Remote: false})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return currencyService_client_stub{stub: stub, convertMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/onlineboutique/currencyservice/CurrencyService", Method: "Convert", Remote: true}), getSupportedCurrenciesMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/onlineboutique/currencyservice/CurrencyService", Method: "GetSupportedCurrencies", Remote: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return currencyService_server_stub{impl: impl.(CurrencyService), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return currencyService_reflect_stub{caller: caller}
		},
		RefData: "",
	})
}

// weaver.InstanceOf checks.
var _ weaver.InstanceOf[CurrencyService] = (*impl)(nil)

// weaver.Router checks.
var _ weaver.Unrouted = (*impl)(nil)

// Local stub implementations.

type currencyService_local_stub struct {
	impl                          CurrencyService
	tracer                        trace.Tracer
	convertMetrics                *codegen.MethodMetrics
	getSupportedCurrenciesMetrics *codegen.MethodMetrics
}

// Check that currencyService_local_stub implements the CurrencyService interface.
var _ CurrencyService = (*currencyService_local_stub)(nil)

func (s currencyService_local_stub) Convert(ctx context.Context, a0 money.T, a1 string) (r0 money.T, err error) {
	// Update metrics.
	begin := s.convertMetrics.Begin()
	defer func() { s.convertMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "currencyservice.CurrencyService.Convert", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Convert(ctx, a0, a1)
}

func (s currencyService_local_stub) GetSupportedCurrencies(ctx context.Context) (r0 []string, err error) {
	// Update metrics.
	begin := s.getSupportedCurrenciesMetrics.Begin()
	defer func() { s.getSupportedCurrenciesMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "currencyservice.CurrencyService.GetSupportedCurrencies", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetSupportedCurrencies(ctx)
}

// Client stub implementations.

type currencyService_client_stub struct {
	stub                          codegen.Stub
	convertMetrics                *codegen.MethodMetrics
	getSupportedCurrenciesMetrics *codegen.MethodMetrics
}

// Check that currencyService_client_stub implements the CurrencyService interface.
var _ CurrencyService = (*currencyService_client_stub)(nil)

func (s currencyService_client_stub) Convert(ctx context.Context, a0 money.T, a1 string) (r0 money.T, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.convertMetrics.Begin()
	defer func() { s.convertMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "currencyservice.CurrencyService.Convert", trace.WithSpanKind(trace.SpanKindClient))
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

	// Encode arguments.
	enc := codegen.NewEncoder()
	(a0).WeaverMarshal(enc)
	enc.String(a1)
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

func (s currencyService_client_stub) GetSupportedCurrencies(ctx context.Context) (r0 []string, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.getSupportedCurrenciesMetrics.Begin()
	defer func() { s.getSupportedCurrenciesMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "currencyservice.CurrencyService.GetSupportedCurrencies", trace.WithSpanKind(trace.SpanKindClient))
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
	r0 = serviceweaver_dec_slice_string_4af10117(dec)
	err = dec.Error()
	return
}

// Note that "weaver generate" will always generate the error message below.
// Everything is okay. The error message is only relevant if you see it when
// you run "go build" or "go run".
var _ codegen.LatestVersion = codegen.Version[[0][20]struct{}](`

ERROR: You generated this file with 'weaver generate' (devel) (codegen
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

type currencyService_server_stub struct {
	impl    CurrencyService
	addLoad func(key uint64, load float64)
}

// Check that currencyService_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*currencyService_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s currencyService_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Convert":
		return s.convert
	case "GetSupportedCurrencies":
		return s.getSupportedCurrencies
	default:
		return nil
	}
}

func (s currencyService_server_stub) convert(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 money.T
	(&a0).WeaverUnmarshal(dec)
	var a1 string
	a1 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Convert(ctx, a0, a1)

	// Encode the results.
	enc := codegen.NewEncoder()
	(r0).WeaverMarshal(enc)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s currencyService_server_stub) getSupportedCurrencies(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.GetSupportedCurrencies(ctx)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_string_4af10117(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

// Reflect stub implementations.

type currencyService_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that currencyService_reflect_stub implements the CurrencyService interface.
var _ CurrencyService = (*currencyService_reflect_stub)(nil)

func (s currencyService_reflect_stub) Convert(ctx context.Context, a0 money.T, a1 string) (r0 money.T, err error) {
	err = s.caller("Convert", ctx, []any{a0, a1}, []any{&r0})
	return
}

func (s currencyService_reflect_stub) GetSupportedCurrencies(ctx context.Context) (r0 []string, err error) {
	err = s.caller("GetSupportedCurrencies", ctx, []any{}, []any{&r0})
	return
}

// Encoding/decoding implementations.

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
