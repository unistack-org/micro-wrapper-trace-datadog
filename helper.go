package datadog

import (
	"context"

	"github.com/unistack-org/micro/v3/metadata"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// StartSpanFromContext returns a new span with the given operation name and options. If a span
// is found in the context, it will be used as the parent of the resulting span.
func StartSpanFromContext(ctx context.Context, operationName string, opts ...tracer.StartSpanOption) (context.Context, tracer.Span, error) {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(metadata.Metadata, 1)
	}

	if spanCtx, err := tracer.Extract(tracer.TextMapCarrier(md)); err == nil {
		opts = append(opts, tracer.ChildOf(spanCtx))
	}

	span, ctx := tracer.StartSpanFromContext(ctx, operationName, opts...)

	if err := tracer.Inject(span.Context(), tracer.TextMapCarrier(md)); err != nil {
		return nil, nil, err
	}

	return metadata.NewContext(ctx, md), span, nil
}
