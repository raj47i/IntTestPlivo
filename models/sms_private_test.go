package models

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SMS", func() {
	var (
		m2 SMS
	)

	BeforeEach(func() {
		m2 = SMS{
			To:   "987654321",
			From: "123456789",
			Text: "This is a sample Text Message used for testing",
		}
	})

	Describe("stopKey", func() {
		It("Should return valid string keys for redis", func() {
			Expect(m2.stopKey()).To(Equal(fmt.Sprintf("STOP:%s:%s", m2.From, m2.To)))
		})
	})

	Describe("daylimitKey", func() {
		It("Should return valid string keys for redis", func() {
			Expect(m2.daylimitKey()).To(Equal(fmt.Sprintf("LAST24:%s", m2.From)))
		})
	})
})
