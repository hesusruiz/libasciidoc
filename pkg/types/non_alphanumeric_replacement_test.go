package types_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"

	. "github.com/onsi/ginkgo/extensions/table" // nolint:golint
	. "github.com/onsi/gomega"                  // nolint:golint
)

var _ = DescribeTable("replace non-alphanumeric chars",

	func(source []interface{}, valueWithDefaultSettings, valueWithCustomSettings string) {
		Expect(types.ReplaceNonAlphanumerics(source, "_", "_")).To(Equal(valueWithDefaultSettings))
		Expect(types.ReplaceNonAlphanumerics(source, "id_", "-")).To(Equal(valueWithCustomSettings))
	},

	Entry("hello",
		[]interface{}{
			&types.StringElement{
				Content: "hello",
			},
		},
		"_hello",
		"id_hello",
	),

	Entry("héllo with an accent",
		[]interface{}{
			&types.StringElement{
				Content: "  héllo 1.2   and 3 Spaces",
			},
		},
		"_héllo_1_2_and_3_spaces",
		"id_héllo-1-2-and-3-spaces",
	),

	Entry("a an accent and a swedish character",
		[]interface{}{
			&types.StringElement{
				Content: `A à ⌘`,
			},
		},
		"_a_à",
		"id_a-à",
	),

	Entry("AŁA",
		[]interface{}{
			&types.StringElement{
				Content: `AŁA 0.1 ?`,
			},
		},
		"_ała_0_1",
		"id_ała-0-1",
	),

	Entry("it's  2 spaces, here !",
		[]interface{}{
			&types.StringElement{
				Content: `it's  2 spaces, here !`,
			},
		},
		"_its_2_spaces_here",
		"id_its-2-spaces-here",
	),

	Entry("content with <strong> markup",
		[]interface{}{
			&types.StringElement{
				Content: "a section title, with",
			},
			&types.QuotedText{
				Kind: types.SingleQuoteBold,
				Elements: []interface{}{
					&types.StringElement{
						Content: "bold content",
					},
				},
			},
		},
		"_a_section_title_with_bold_content",
		"id_a-section-title-with-bold-content",
	),

	Entry("content with link",
		[]interface{}{
			&types.StringElement{
				Content: "link to ",
			},
			&types.InlineLink{
				Attributes: types.Attributes{},
				Location: &types.Location{
					Scheme: "https://",
					Path:   "foo.bar",
				},
			},
		},
		"_link_to_httpsfoo_bar",
		"id_link-to-httpsfoo-bar",
	),

	Entry("content with dots and special characters",
		[]interface{}{
			&types.StringElement{
				Content: "...and we're back!",
			},
		},
		"_and_were_back",
		"id_-and-were-back",
	),

	Entry("content with dots",
		[]interface{}{
			&types.StringElement{
				Content: "Section A.a",
			},
		},
		"_section_a_a",
		"id_section-a-a",
	),

	Entry("content with quoted string",
		// Block Quotes and "`Smart`" Ones
		[]interface{}{
			&types.StringElement{
				Content: "Block Quotes and ",
			},
			&types.QuotedString{
				Kind: types.DoubleQuote,
				Elements: []interface{}{
					&types.StringElement{
						Content: "Smart",
					},
				},
			},
			&types.StringElement{
				Content: "Ones",
			},
		},
		"_block_quotes_and_smart_ones",
		"id_block-quotes-and-smart-ones",
	),
)
