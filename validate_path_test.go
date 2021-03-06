package apidVerifyApiKey

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Validate Path", func() {

	It("validation1", func() {
		s := validatePath("", "/foo")
		Expect(s).Should(BeTrue())
	})
	It("validation2", func() {
		s := validatePath("", "foo")
		Expect(s).Should(BeTrue())
	})
	It("validation3", func() {
		s := validatePath("{}", "foo")
		Expect(s).Should(BeTrue())
	})
	It("validation4", func() {
		s := validatePath("{/**}", "/foo")
		Expect(s).Should(BeTrue())
	})
	It("validation5", func() {
		s := validatePath("{/**}", "foo")
		Expect(s).Should(BeFalse())
	})
	It("validation6", func() {
		s := validatePath("{/**}", "/")
		Expect(s).Should(BeTrue())
	})
	It("validation7", func() {
		s := validatePath("{/foo/**}", "/")
		Expect(s).Should(BeFalse())
	})
	It("validation8", func() {
		s := validatePath("{/foo/**}", "/foo/")
		Expect(s).Should(BeTrue())
	})
	It("validation9", func() {
		s := validatePath("{/foo/**}", "/foo/bar")
		Expect(s).Should(BeTrue())
	})
	It("validation10", func() {
		s := validatePath("{/foo/**}", "foo")
		Expect(s).Should(BeFalse())
	})
	It("validation11", func() {
		s := validatePath("{/foo/bar/**}", "/foo/bar/xx/yy")
		Expect(s).Should(BeTrue())
	})
	It("validation12", func() {
		s := validatePath("/foo/bar/*}", "/foo/bar/xxx")
		Expect(s).Should(BeTrue())
	})
	It("validation13", func() {
		s := validatePath("{/foo/bar/*/}", "/foo/bar/xxx")
		Expect(s).Should(BeFalse())
	})
	It("validation14", func() {
		s := validatePath("{/foo/bar/**}", "/foo/bar/xx/yy")
		Expect(s).Should(BeTrue())
	})
	It("validation15", func() {
		s := validatePath("{/foo/*/**/}", "/foo/bar")
		Expect(s).Should(BeFalse())
	})
	It("validation16", func() {
		s := validatePath("{/foo/bar/*/xxx}", "/foo/bar/yyy/xxx")
		Expect(s).Should(BeTrue())
	})
	It("validation17", func() {
		s := validatePath("{/foo/bar/*/xxx/}", "/foo/bar/yyy/xxx")
		Expect(s).Should(BeFalse())
	})
	It("validation18", func() {
		s := validatePath("{/foo/bar/**/xxx/}", "/foo/bar/aaa/bbb/xxx/")
		Expect(s).Should(BeTrue())
	})
	It("validation19", func() {
		s := validatePath("{/foo/bar/***/xxx/}", "/foo/bar/aaa/bbb/xxx/")
		Expect(s).Should(BeTrue())
	})
	It("validation20", func() {
		s := validatePath("{/foo/, /bar/}", "/foo/")
		Expect(s).Should(BeTrue())
	})
	It("validation21", func() {
		s := validatePath("{/foo/bar/yy*/xxx}", "/foo/bar/yyy/xxx")
		Expect(s).Should(BeTrue())
	})
})
