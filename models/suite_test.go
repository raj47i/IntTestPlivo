package models_test

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/raj47i/IntTestPlivo/config"
)

func TestMain(m *testing.M) {
	// Enable Ginko-Gomega
	RegisterFailHandler(Fail)
	// Setup Env
	config.FlushCache()
	// Run Test Suite
	rs := m.Run()
	// TearDown
	config.FlushCache()
	os.Exit(rs)
}
func TestModels(t *testing.T) {
	// Enable Ginko-Gomega
	RegisterFailHandler(Fail)
	// Run tests
	RunSpecs(t, "IntTestPlivo:Models Test Suite")
}
