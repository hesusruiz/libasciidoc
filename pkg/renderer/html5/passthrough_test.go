package html5_test

import . "github.com/onsi/ginkgo"

var _ = Describe("passthroughs", func() {

	Context("tripleplus passthrough", func() {

		It("an empty standalone tripleplus passthrough", func() {
			source := `++++++`
			expected := ``
			verify("test.adoc", expected, source)
		})

		It("an empty tripleplus passthrough in a paragraph", func() {
			source := `++++++ with more content afterwards...`
			expected := `<div class="paragraph">
<p> with more content afterwards&#8230;&#8203;</p>
</div>`
			verify("test.adoc", expected, source)
		})

		It("a standalone tripleplus passthrough", func() {
			source := `+++*bold content*+++`
			expected := `<div class="paragraph">
<p>*bold content*</p>
</div>`
			verify("test.adoc", expected, source)
		})

		It("tripleplus passthrough in paragraph", func() {
			source := `The text +++<u>underline & me</u>+++ is underlined.`
			expected := `<div class="paragraph">
<p>The text <u>underline & me</u> is underlined.</p>
</div>`
			verify("test.adoc", expected, source)
		})
	})

	Context("singleplus Passthrough", func() {

		It("an empty standalone singleplus passthrough", func() {
			source := `++`
			expected := `<div class="paragraph">
<p>&#43;&#43;</p>
</div>`
			verify("test.adoc", expected, source)
		})

		It("an empty singleplus passthrough in a paragraph", func() {
			source := `++ with more content afterwards...`
			expected := `<div class="paragraph">
<p>&#43;&#43; with more content afterwards&#8230;&#8203;</p>
</div>`
			verify("test.adoc", expected, source)
		})

		It("a singleplus passthrough", func() {
			source := `+*bold content*+`
			expected := `<div class="paragraph">
<p>*bold content*</p>
</div>`
			verify("test.adoc", expected, source)
		})

		It("singleplus passthrough in paragraph", func() {
			source := `The text +<u>underline me</u>+ is not underlined.`
			expected := `<div class="paragraph">
<p>The text &lt;u&gt;underline me&lt;/u&gt; is not underlined.</p>
</div>`
			verify("test.adoc", expected, source)
		})

		It("invalid singleplus passthrough in paragraph", func() {
			source := `The text + *hello*, world + is not passed through.`
			expected := `<div class="paragraph">
<p>The text &#43; <strong>hello</strong>, world &#43; is not passed through.</p>
</div>`
			verify("test.adoc", expected, source)
		})
	})

	Context("passthrough Macro", func() {

		It("passthrough macro with single word", func() {
			source := `pass:[hello]`
			expected := `<div class="paragraph">
<p>hello</p>
</div>`
			verify("test.adoc", expected, source)
		})

		It("passthrough macro with words", func() {
			source := `pass:[hello, world]`
			expected := `<div class="paragraph">
<p>hello, world</p>
</div>`
			verify("test.adoc", expected, source)
		})

		It("empty passthrough macro", func() {
			source := `pass:[]`
			expected := ``
			verify("test.adoc", expected, source)
		})

		It("passthrough macro with spaces", func() {
			source := `pass:[ *hello*, world ]`
			expected := `<div class="paragraph">
<p> *hello*, world </p>
</div>`
			verify("test.adoc", expected, source)
		})

		It("passthrough macro with line break", func() {
			source := "pass:[hello,\nworld]"
			expected := `<div class="paragraph">
<p>hello,
world</p>
</div>`
			verify("test.adoc", expected, source)
		})
	})

	Context("passthrough Macro with Quoted Text", func() {

		It("passthrough macro with single quoted word", func() {
			source := `pass:q[*hello*]`
			expected := `<div class="paragraph">
<p><strong>hello</strong></p>
</div>`
			verify("test.adoc", expected, source)
		})

		It("passthrough macro with quoted word in sentence and trailing spaces", func() {
			source := `pass:q[ a *hello*, world ]   `
			expected := `<div class="paragraph">
<p> a <strong>hello</strong>, world </p>
</div>`
			verify("test.adoc", expected, source)
		})

		It("passthrough macro within paragraph", func() {
			source := `an pass:q[ *hello*, world ] mention`
			expected := `<div class="paragraph">
<p>an  <strong>hello</strong>, world  mention</p>
</div>`
			verify("test.adoc", expected, source)
		})
	})
})
