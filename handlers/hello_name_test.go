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

var _ = Describe("Hello Name Handler", func() {
	var (
		response *httptest.ResponseRecorder
		router   *gin.Engine
		req      *http.Request
	)

	BeforeEach(func() {
		router = gin.Default()
		router.GET("/hello/:name", handlers.HelloNameHandler())
		response = httptest.NewRecorder()
	})

	When("a name is provided in the URL", func() {
		testName := "Gopher"

		BeforeEach(func() {
			req, _ = http.NewRequest("GET", fmt.Sprintf("/hello/%s", testName), nil)
			router.ServeHTTP(response, req)
		})

		It("should return 200 OK", func() {
			Expect(response.Code).To(Equal(http.StatusOK))
		})

		It("should return a JSON response with the correct greeting", func() {
			expectedJSON := fmt.Sprintf(`{"message":"Hello, %s!"}`, testName)
			Expect(response.Body.String()).To(MatchJSON(expectedJSON))
		})
	})

})
