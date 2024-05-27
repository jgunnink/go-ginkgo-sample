package handlers_test

import (
	"gobdd/handlers"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hello Handler", func() {
	var (
		response *httptest.ResponseRecorder
		router   *gin.Engine
		req      *http.Request
	)

	BeforeEach(func() {
		router = gin.Default()
		router.GET("/hello", handlers.HelloHandler())
		response = httptest.NewRecorder()
	})

	When("GET /hello is invoked", func() {
		BeforeEach(func() {
			req, _ = http.NewRequest("GET", "/hello", nil)
			router.ServeHTTP(response, req)
		})

		It("should return 200 OK", func() {
			Expect(response.Code).To(Equal(http.StatusOK))
		})

		It("should return 'Hello, World!'", func() {
			Expect(response.Body.String()).To(Equal("Hello, World!"))
		})
	})
})
