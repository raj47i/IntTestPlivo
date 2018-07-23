package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/raj47i/IntTestPlivo/models"
)

var _ = Describe("PhoneNumber", func() {
	var (
		pn1 PhoneNumber
		pn2 PhoneNumber
	)
	BeforeEach(func() {
	})

	Describe("LoadByNumberAndAccountID", func() {
		It("Should load details from database for valid params", func() {
			r, e := pn1.LoadByNumberAndAccountID("9999", 0)
			Expect(e).NotTo(HaveOccurred())
			Expect(r).To(BeFalse())
			Expect(pn1.ID).To(BeZero())

			r, e = pn1.LoadByNumberAndAccountID("4924195509198", 1)
			Expect(e).NotTo(HaveOccurred())
			Expect(r).To(BeTrue())
			Expect(pn1.ID).To(BeNumerically("==", 1))

			r, e = pn2.LoadByNumberAndAccountID("61871112902", 3)
			Expect(e).NotTo(HaveOccurred())
			Expect(r).To(BeTrue())
			Expect(pn2.ID).To(BeNumerically("==", 60))

		})
	})
})
