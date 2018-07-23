package models_test

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/raj47i/IntTestPlivo/models"
)

func generatePhone() string {
	return "9" + strings.Repeat(strconv.Itoa(rand.Intn(9)), rand.Intn(9)+6)
}

var _ = Describe("SMS", func() {
	var (
		m1   SMS
		m2   SMS
		to   string
		from string
		text string
	)

	BeforeEach(func() {
		to = "987654321"
		from = "123456789"
		text = "This is a sample Text Message used for testing"
		m1 = SMS{}
		m2 = SMS{
			To:   to,
			From: from,
			Text: text,
		}
	})

	Describe("DayLimit", func() {
		var mx [51]SMS
		It("Should count messages till 50 with out errors", func() {
			for i := 0; i < 50; i++ {
				mx[i].From = from
				mx[i].To = generatePhone()
				mx[i].Text = text
				exceeded, err := mx[i].DayLimit()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(exceeded).To(BeFalse())
			}
		})
		It("Should count messages and raise error for more than 50 messages in 24 hours", func() {
			mx[50].From = from
			mx[50].To = generatePhone()
			mx[50].Text = text
			exceeded, err := mx[50].DayLimit()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(exceeded).To(BeTrue())
		})
	})

	Describe("Block", func() {
		It("Should mark the from:to as blocked", func() {
			m1.Parse(to, from, "STOP")
			m1.Block()
			Expect(m2.IsBlocked()).To(BeTrue())
		})
	})

	Describe("IsBlocked", func() {
		It("Should be true only for blocked messages", func() {
			m1.Parse("9876543212", "1234567890", "STOP")
			Expect(m1.IsBlocked()).To(BeFalse())

			m1.Block()
			Expect(m1.IsBlocked()).To(BeTrue())
		})
	})

	Describe("IsBlockCommand", func() {
		It("Should be true only for valid STOP messages", func() {
			m1.Parse(to, from, "STOP")
			Expect(m1.IsBlockCommand()).To(BeTrue())

			m1.Parse(to, from, "STOP\r")
			Expect(m1.IsBlockCommand()).To(BeTrue())

			m1.Parse(to, from, "STOP\n")
			Expect(m1.IsBlockCommand()).To(BeTrue())

			m1.Parse(to, from, "STOP\r\n")
			Expect(m1.IsBlockCommand()).To(BeTrue())

			m1.Parse(to, from, "STOP\n\r")
			Expect(m1.IsBlockCommand()).To(BeTrue())

			m1.Parse(to, from, "stop")
			Expect(m1.IsBlockCommand()).To(BeFalse())

			m1.Parse(to, from, "\nSTOP")
			Expect(m1.IsBlockCommand()).To(BeFalse())

			m1.Parse(to, from, "Hello")
			Expect(m1.IsBlockCommand()).To(BeFalse())
		})
	})

	Describe("Parse", func() {
		Context("With empty To, From Or Text", func() {
			It("Should return Error", func() {
				Expect(m1.Parse("", "", "")).Should(MatchError(errors.New("to is missing")))
				Expect(m1.Parse("", "987654321", "Non empty text")).Should(MatchError(errors.New("to is missing")))
				Expect(m1.Parse("987654321", "", "Non empty text")).Should(MatchError(errors.New("from is missing")))
				Expect(m1.Parse("987654321", "123456789", "")).Should(MatchError(errors.New("text is missing")))
			})
		})

		Context("With invalid to", func() {
			It("Should return Error", func() {
				Expect(m1.Parse("98765", "123456789", "Non empty text")).Should(MatchError(errors.New("to is invalid")))
				Expect(m1.Parse("12345678901234567", "123456789", "Non empty text")).Should(MatchError(errors.New("to is invalid")))
				Expect(m1.Parse("ABCDEFGH", "123456789", "Non empty text")).Should(MatchError(errors.New("to is invalid")))
			})
		})

		Context("With invalid From", func() {
			It("Should return Error", func() {
				Expect(m1.Parse("123456789", "98765", "Non empty text")).Should(MatchError(errors.New("from is invalid")))
				Expect(m1.Parse("123456789", "12345678901234567", "Non empty text")).Should(MatchError(errors.New("from is invalid")))
				Expect(m1.Parse("123456789", "abcd$3sfg", "Non empty text")).Should(MatchError(errors.New("from is invalid")))
			})
		})

		Context("Only With invalid Text", func() {
			It("Should return Error", func() {
				text := strings.Repeat("abcdefghij", 12) // string with 120 chars length
				Expect(m1.Parse("123456789", "987654321", text)).Should(BeNil())
				Expect(m1.Parse("123456789", "987654321", text+"1")).Should(MatchError(errors.New("text is invalid")))
			})
		})

		Context("With valid To, From, Text", func() {
			It("Should return a valid SMS Object", func() {
				to := "123456789"
				from := "987654321"
				text := "Hello"
				Expect(m1.Parse(to, from, text)).Should(BeNil())
				Expect(m1.To).To(Equal(to))
				Expect(m1.From).To(Equal(from))
				Expect(m1.Text).To(Equal(text))
			})
		})
	})

})
