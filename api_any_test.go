package main_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("API: Any undefined", func() {
	Describe("PATH | VERB ", func() {
		Context("GET", func() {
			It("Should return 405", func() {
				hR, err := Http("GET", "/inbound/sms/", nil, nil)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(hR).ToNot(BeNil())
				Expect(hR.Code).To(Equal(http.StatusMethodNotAllowed))
			})
		})

		Context("GET", func() {
			It("Should return 405", func() {
				hR, err := Http("GET", "/", nil, nil)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(hR).ToNot(BeNil())
				Expect(hR.Code).To(Equal(http.StatusMethodNotAllowed))
			})
		})

		Context("POST", func() {
			It("Should return 405", func() {
				hR, err := Http("GET", "/", nil, nil)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(hR).ToNot(BeNil())
				Expect(hR.Code).To(Equal(http.StatusMethodNotAllowed))
			})
		})

		Context("DELETE", func() {
			It("Should return 405", func() {
				hR, err := Http("DELETE", "/outbound/sms/", nil, nil)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(hR).ToNot(BeNil())
				Expect(hR.Code).To(Equal(http.StatusMethodNotAllowed))
			})
		})

	})

})
