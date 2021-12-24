package xhtml5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("thematic breaks", func() {

	It("simple", func() {
		source := `before

'''

after`
		expected := `<div class="paragraph">
<p>before</p>
</div>
<hr/>
<div class="paragraph">
<p>after</p>
</div>
`
		Expect(RenderXHTML(source)).To(MatchHTML(expected))
	})
})