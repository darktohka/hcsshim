//go:build windows

package wclayer

import (
	"context"

	"github.com/Microsoft/hcsshim/internal/hcserror"
	"github.com/Microsoft/hcsshim/internal/oc"
	"go.opentelemetry.io/otel/trace"
)

// DestroyLayer will remove the on-disk files representing the layer with the given
// path, including that layer's containing folder, if any.
func DestroyLayer(ctx context.Context, path string) (err error) {
	title := "hcsshim::DestroyLayer"
	ctx, span := oc.StartSpan(ctx, title) //nolint:ineffassign,staticcheck
	defer span.End()
	defer func() { oc.SetSpanStatus(span, err) }()
	span.AddAttributes(trace.StringAttribute("path", path))

	err = destroyLayer(&stdDriverInfo, path)
	if err != nil {
		return hcserror.New(err, title, "")
	}
	return nil
}
