//go:build windows

package wclayer

import (
	"context"

	"github.com/Microsoft/hcsshim/internal/hcserror"
	"github.com/Microsoft/hcsshim/internal/oc"
	"go.opentelemetry.io/otel/trace"
)

// CreateLayer creates a new, empty, read-only layer on the filesystem based on
// the parent layer provided.
func CreateLayer(ctx context.Context, path, parent string) (err error) {
	title := "hcsshim::CreateLayer"
	ctx, span := oc.StartSpan(ctx, title) //nolint:ineffassign,staticcheck
	defer span.End()
	defer func() { oc.SetSpanStatus(span, err) }()
	span.AddAttributes(
		trace.StringAttribute("path", path),
		trace.StringAttribute("parent", parent))

	err = createLayer(&stdDriverInfo, path, parent)
	if err != nil {
		return hcserror.New(err, title, "")
	}
	return nil
}
