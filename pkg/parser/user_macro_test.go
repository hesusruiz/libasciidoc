package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("user macros", func() {

	Context("final documents", func() {

		Context("inline macros", func() {

			It("inline macro empty", func() {
				source := "AAA hello:[]"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "AAA ",
									},
									types.UserMacro{
										Kind:    types.InlineMacro,
										Name:    "hello",
										Value:   "",
										RawText: "hello:[]",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline macro with double quoted attributes", func() {
				source := `AAA hello:[prefix="hello ",suffix="!!"]`
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{types.StringElement{
									Content: "AAA ",
								},
									types.UserMacro{
										Kind:  types.InlineMacro,
										Name:  "hello",
										Value: "",
										Attributes: types.Attributes{
											"prefix": "hello ",
											"suffix": "!!",
										},
										RawText: `hello:[prefix="hello ",suffix="!!"]`,
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline macro with value", func() {
				source := `AAA hello:JohnDoe[]`
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{types.StringElement{
									Content: "AAA ",
								},
									types.UserMacro{
										Kind:    types.InlineMacro,
										Name:    "hello",
										Value:   "JohnDoe",
										RawText: "hello:JohnDoe[]",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline user macro with value and attributes", func() {
				source := "repository: git:some/url.git[key1=value1,key2=value2]"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{types.StringElement{
									Content: "repository: ",
								},
									types.UserMacro{
										Kind:  types.InlineMacro,
										Name:  "git",
										Value: "some/url.git",
										Attributes: types.Attributes{
											"key1": "value1",
											"key2": "value2",
										},
										RawText: "git:some/url.git[key1=value1,key2=value2]",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})

		Context("user macros", func() {

			It("user macro block without value", func() {

				source := "git::[]"
				expected := types.Document{
					Elements: []interface{}{
						types.UserMacro{
							Kind:    types.BlockMacro,
							Name:    "git",
							Value:   "",
							RawText: "git::[]",
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("user block macro with value and attributes", func() {
				source := "git::some/url.git[key1=value1,key2=value2]"
				expected := types.Document{
					Elements: []interface{}{
						types.UserMacro{
							Kind:  types.BlockMacro,
							Name:  "git",
							Value: "some/url.git",
							Attributes: types.Attributes{
								"key1": "value1",
								"key2": "value2",
							},
							RawText: "git::some/url.git[key1=value1,key2=value2]",
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("user macro block with attribute", func() {
				source := `git::[key1="value1"]`
				expected := types.Document{
					Elements: []interface{}{
						types.UserMacro{
							Kind:  types.BlockMacro,
							Name:  "git",
							Value: "",
							Attributes: types.Attributes{
								"key1": "value1",
							},
							RawText: `git::[key1="value1"]`,
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("user macro block with value", func() {
				source := `git::some/url.git[]`
				expected := types.Document{
					Elements: []interface{}{
						types.UserMacro{
							Kind:    types.BlockMacro,
							Name:    "git",
							Value:   "some/url.git",
							RawText: "git::some/url.git[]",
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

		})
	})
})
