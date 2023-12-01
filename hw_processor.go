// custom_processor.go
package hw_processor

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel/api/metric"
)

type HelloWorldProcessor struct{}

func (p *HelloWorldProcessor) ProcessStart(ctx context.Context, result metric.Int64ObserverResult) {
	fmt.Println("Hello World Start")
	result.Observe(0, metric.NewLabelSet())
}

func (p *HelloWorldProcessor) ProcessEnd(ctx context.Context, result metric.Int64ObserverResult) {
	fmt.Println("Hello World End")
	result.Observe(0, metric.NewLabelSet())
}
