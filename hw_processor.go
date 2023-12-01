// custom_processor.go
package hw_processor

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel/metric"
)

type HelloWorldProcessor struct{}

func (p *HelloWorldProcessor) ProcessStart(ctx context.Context, record metric.Record) (context.Context, metric.Int64Number, metric.Float64Number) {
	fmt.Println("Hello World Start")
	return ctx, metric.NewInt64Number(0), metric.NewFloat64Number(0)
}

func (p *HelloWorldProcessor) ProcessEnd(ctx context.Context, startCtx context.Context, record metric.Record) {
	fmt.Println("Hello World End")
}
