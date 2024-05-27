package handlers_test

import (
	"fmt"
	"gobdd/handlers"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hello Calculator Handler", func() {
	var (
		response *httptest.ResponseRecorder
		router   *gin.Engine
		req      *http.Request
	)

	BeforeEach(func() {
		router = gin.Default()
		router.GET("/hello/calculate", handlers.HelloCalculatorHandler())
		response = httptest.NewRecorder()
	})

	When("a two numbers are provided as query parameters", func() {
		BeforeEach(func() {
			req, _ = http.NewRequest("GET", fmt.Sprintf("/hello/calculate?num1=%s&num2=%s", "1", "2"), nil)
			router.ServeHTTP(response, req)
		})

		It("should return 200 OK", func() {
			Expect(response.Code).To(Equal(http.StatusOK))
		})

		It("returns the sum of the two numbers in JSON", func() {
			expectedJSON := fmt.Sprintf(`{"result":%d}`, 3)
			Expect(response.Body.String()).To(MatchJSON(expectedJSON))
		})
	})

	When("a non-numeric number is provided as a query parameter", func() {
		BeforeEach(func() {
			req, _ = http.NewRequest("GET", fmt.Sprintf("/hello/calculate?num1=%s&num2=%s", "p", "2"), nil)
			router.ServeHTTP(response, req)
		})

		It("should return 400 Bad Request", func() {
			Expect(response.Code).To(Equal(http.StatusBadRequest))
		})

		It("returns the sum of the two numbers in JSON", func() {
			expectedJSON := `{"error": "Invalid num1 parameter"}`
			Expect(response.Body.String()).To(MatchJSON(expectedJSON))
		})
	})
})
