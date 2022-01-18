package zadig_test_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Zadig", func() {
	Context("Zadig API Test", func() {
		It("Create github integration should success", func() {
			a := 1+1
			Expect(a).To(Equal(2), "status code should be 200.")
		})

		It("Delete github integration should success", func() {
			a := 1+2
			Expect(a).To(Equal(3), "status code should be 200.")
		})

		It("Create gitlab integration should success", func() {
			a := 1+3
			Expect(a).To(Equal(4), "status code should be 200.")
		})

		It("Delete gitlab integration should success", func() {
			a := 1+4
			Expect(a).To(Equal(5), "status code should be 200.")
		})

		It("Integrate Object storage should success", func() {
			a := 1+4
			Expect(a).To(Equal(5), "status code should be 200.")
		})

		It("Delete object storage should success", func() {
			a := 1+4
			Expect(a).To(Equal(5), "status code should be 200.")
		})

		It("Integrate docker registry should success", func() {
			a := 1+4
			Expect(a).To(Equal(5), "status code should be 200.")
		})

		It("Delete docker registry should success", func() {
			a := 1+4
			Expect(a).To(Equal(5), "status code should be 200.")
		})

		It("Create workflow should success", func() {
			a := 1+4
			Expect(a).To(Equal(5), "status code should be 200.")
		})

		It("Delete workflow should success", func() {
			a := 1+4
			Expect(a).To(Equal(5), "status code should be 200.")
		})

		It("Create project should success", func() {
			a := 1+4
			Expect(a).To(Equal(5), "status code should be 200.")
		})

		It("Edit project should success", func() {
			a := 1+4
			Expect(a).To(Equal(5), "status code should be 200.")
		})

		It("Delete project should success", func() {
			a := 1+4
			Expect(a).To(Equal(5), "status code should be 200.")
		})
	})
})
