package sgml

import (
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func (r *sgmlRenderer) renderAttributeSubstitution(ctx *renderer.Context, e *types.AttributeSubstitution) (string, error) {
	if v, found := ctx.Attributes[e.Name]; found {
		switch v := v.(type) {
		case string:
			return v, nil
		default:
			return r.renderElement(ctx, v)
		}
	}
	return "{" + e.Name + "}", nil
}