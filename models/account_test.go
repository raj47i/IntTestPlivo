package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/raj47i/IntTestPlivo/models"
)

var _ = Describe("Account", func() {
	var (
		a1 Account
		a2 Account
	)

	BeforeEach(func() {
		a1 = Account{
			ID:       1,
			AuthID:   "auth1st",
			Username: "userX1",
		}
		a2 = Account{}
	})

	Describe("Authentication", func() {
		Context("With correct secret", func() {
			It("should be successful", func() {
				Expect(a1.Authenticate("auth1st")).To(BeTrue())
			})
		})

		Context("With incorrect secret", func() {
			It("should fail", func() {
				Expect(a1.Authenticate("Auth1s")).To(BeFalse())
			})
		})
	})

	// The database is readonly, as per SRS - so safe to test againt.
	Describe("LoadByUserName", func() {
		It("should load an account by its username", func() {
			Expect(a2.LoadByUserName("plivo2")).To(BeTrue())
			Expect(a2.ID).Should(BeNumerically("==", 2))
			Expect(a2.AuthID).Should(Equal("54P2EOKQ47"))
		})
		It("should not load an account for invalid usernames", func() {
			Expect(a2.LoadByUserName("this-name-is-not-in-db")).To(BeFalse())
			Expect(a2.ID).To(BeZero())
			Expect(a2.AuthID).To(BeEmpty())
			Expect(a2.Username).To(BeEmpty())
		})
	})

})
