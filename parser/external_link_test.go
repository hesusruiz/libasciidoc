package parser_test

import (
	"github.com/bytesparadise/libasciidoc/parser"
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("External Links", func() {

	It("external link", func() {
		actualContent := "a link to https://foo.bar"
		expectedDocument := &types.InlineContent{
			Elements: []types.InlineElement{
				&types.StringElement{Content: "a link to "},
				&types.ExternalLink{
					URL: "https://foo.bar",
				},
			},
		}
		verify(GinkgoT(), expectedDocument, actualContent, parser.Entrypoint("InlineContent"))
	})

	It("external link with empty text", func() {
		actualContent := "a link to https://foo.bar[]"
		expectedDocument := &types.InlineContent{
			Elements: []types.InlineElement{
				&types.StringElement{Content: "a link to "},
				&types.ExternalLink{
					URL:  "https://foo.bar",
					Text: "",
				},
			},
		}
		verify(GinkgoT(), expectedDocument, actualContent, parser.Entrypoint("InlineContent"))
	})

	It("external link with text", func() {
		actualContent := "a link to mailto:foo@bar[the foo@bar email]"
		expectedDocument := &types.InlineContent{
			Elements: []types.InlineElement{
				&types.StringElement{Content: "a link to "},
				&types.ExternalLink{
					URL:  "mailto:foo@bar",
					Text: "the foo@bar email",
				},
			},
		}
		verify(GinkgoT(), expectedDocument, actualContent, parser.Entrypoint("InlineContent"))
	})
})