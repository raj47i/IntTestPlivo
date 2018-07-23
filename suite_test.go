package main_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/raj47i/IntTestPlivo"
	"github.com/raj47i/IntTestPlivo/config"
)

func TestMain(m *testing.M) {
	//Enable Ginko-Gomega
	RegisterFailHandler(Fail)
	// Setup Env
	config.FlushCache()
	//Run Test Suite
	rs := m.Run()
	// TearDown
	config.FlushCache()
	os.Exit(rs)
}

func TestAll(t *testing.T) {
	// Enable Ginko-Gomega
	RegisterFailHandler(Fail)
	// Run tests
	RunSpecs(t, "IntTestPlivo Test Suite")
}

/* Helper functions and types for testing gin Routes/http */

type HttpResponse struct {
	Code    int
	Error   string `json:"error"`
	Message string `json:"message"`
}

func Http(verb string, path string, headers map[string]string, data map[string]string) (*HttpResponse, error) {
	var req *http.Request
	if data != nil {
		pData := url.Values{}
		for k, v := range data {
			pData.Set(k, v)
		}
		req, _ = http.NewRequest(verb, path, strings.NewReader(pData.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req, _ = http.NewRequest(verb, path, nil)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	r := gin.Default()
	InitRoutes(r)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	hR := HttpResponse{Code: w.Code}
	p, _ := ioutil.ReadAll(w.Body)
	if err := json.Unmarshal(p, &hR); err != nil {
		return nil, err
	}
	// hR.Code = w.Code
	return &hR, nil
}
