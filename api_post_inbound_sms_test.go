package main_test

import (
	"encoding/base64"
	"net/http"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("REST API", func() {

	var (
		path    string
		headers map[string]string
	)

	BeforeEach(func() {
		path = "/inbound/sms/"
		headers = make(map[string]string)
		headers["Authorization"] = "Basic " + base64.StdEncoding.EncodeToString([]byte("plivo4:YHWE3HDLPQ"))
	})

	Describe("POST /inbound/sms/", func() {
		Context("With invalid or no Authorization header", func() {
			It("Should return 401", func() {
				hR, err := Http("POST", path, nil, nil)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(hR).ToNot(BeNil())
				Expect(hR.Code).To(Equal(http.StatusUnauthorized))
			})
			It("Should return 401", func() {
				headers := make(map[string]string)
				headers["Authorization"] = "invalid"
				hR, err := Http("POST", path, headers, nil)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(hR).ToNot(BeNil())
				Expect(hR.Code).To(Equal(http.StatusUnauthorized))
				Expect(hR.Message).To(BeEmpty())
			})
		})

		Context("POST with valid Authorization header", func() {
			It("Should not return 401", func() {
				hR, err := Http("POST", path, headers, nil)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(hR).ToNot(BeNil())
				Expect(hR.Code).ToNot(Equal(http.StatusUnauthorized))
				Expect(hR.Message).To(BeEmpty())
			})
		})

		Context("POST with out input parameters", func() {
			It("Should return 422", func() {
				hR, err := Http("POST", path, headers, nil)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(hR.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(hR.Message).To(BeEmpty())
				Expect(hR.Error).To(Equal("to is missing"))
			})
			It("Should return 422", func() {
				data := map[string]string{
					"to":   "987654321",
					"from": "",
					"text": "hello",
				}
				hR, err := Http("POST", path, headers, data)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(hR.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(hR.Message).To(BeEmpty())
				Expect(hR.Error).To(Equal("from is missing"))
			})
			It("Should return 422", func() {
				data := map[string]string{
					"to":   "987654321",
					"from": "123455678",
					"text": "",
				}
				hR, err := Http("POST", path, headers, data)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(hR.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(hR.Message).To(BeEmpty())
				Expect(hR.Error).To(Equal("text is missing"))
			})
		})

		Context("POST with invalid input parameters", func() {
			It("Should return 422", func() {
				data := map[string]string{
					"to":   "987",
					"from": "1234567",
					"text": "hello",
				}
				hR, err := Http("POST", path, headers, data)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(hR.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(hR.Message).To(BeEmpty())
				Expect(hR.Error).To(Equal("to is invalid"))
			})
			It("Should return 422", func() {
				data := map[string]string{
					"to":   "987654321",
					"from": "12345678901234567",
					"text": "hello",
				}
				hR, err := Http("POST", path, headers, data)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(hR.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(hR.Message).To(BeEmpty())
				Expect(hR.Error).To(Equal("from is invalid"))
			})
			It("Should return 422", func() {
				data := map[string]string{
					"to":   "987654321",
					"from": "1234567890123456",
					"text": strings.Repeat("x", 121),
				}
				hR, err := Http("POST", path, headers, data)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(hR.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(hR.Message).To(BeEmpty())
				Expect(hR.Error).To(Equal("text is invalid"))
			})
		})
		Context("POST with unknown to numners", func() {
			It("Should return 422", func() {
				data := map[string]string{
					"to":   "987654321", // does not exist in db
					"from": "1234567890123456",
					"text": "Hello",
				}
				hR, err := Http("POST", path, headers, data)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(hR.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(hR.Message).To(BeEmpty())
				Expect(hR.Error).To(Equal("to parameter not found"))
			})
			It("Should return 422", func() {
				data := map[string]string{
					"to":   "4924195509195", // belongs to user:plivo1
					"from": "1234567890123456",
					"text": "Hello",
				}
				hR, err := Http("POST", path, headers, data)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(hR.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(hR.Message).To(BeEmpty())
				Expect(hR.Error).To(Equal("to parameter not found"))
			})
		})
		Context("POST with valid parameters", func() {
			It("Should return 200", func() {
				data := map[string]string{
					"to":   "61871112920",
					"from": "1234567890123456",
					"text": "Hello",
				}
				hR, err := Http("POST", path, headers, data)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(hR.Code).To(Equal(http.StatusOK))
				Expect(hR.Error).To(BeEmpty())
				Expect(hR.Message).To(Equal("inbound sms ok"))
			})
		})

	})
})
