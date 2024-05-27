package handlers_test

import (
	"bytes"
	"errors"
	"fmt"
	"gobdd/handlers"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type roundTripperFunc func(req *http.Request) (*http.Response, error)

// RoundTrip implements the RoundTripper interface.
func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req) // Call the underlying function with the request
}

var _ = Describe("HelloThirdPartyApiHandler", func() {
	var (
		mockClient *http.Client
		response   *httptest.ResponseRecorder
		ctx        *gin.Context

		testUrl string = "https://test-url.com/users/1"
	)

	BeforeEach(func() {
		mockClient = &http.Client{
			Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewBufferString(`{"id": 1, "name": "John Smith"}`)),
					Header:     make(http.Header),
				}, nil
			}),
		}

		gin.SetMode(gin.TestMode)
		response = httptest.NewRecorder()
		ctx, _ = gin.CreateTestContext(response)
		ctx.Request = httptest.NewRequest("GET", fmt.Sprintf("/hello/third-party?url=%s", testUrl), nil)
	})

	It("should return the user data from the third-party API on success", func() {
		handler := handlers.HelloThirdPartyApiHandler(mockClient)
		handler(ctx)

		Expect(response.Code).To(Equal(http.StatusOK))
		Expect(response.Body.String()).To(Equal(`{"id":1,"name":"John Smith"}`))
	})

	It("should return a bad request if the url is not provided", func() {
		handler := handlers.HelloThirdPartyApiHandler(mockClient)
		ctx.Request = httptest.NewRequest(http.MethodGet, "/hello/third-party", nil)
		handler(ctx)

		Expect(response.Code).To(Equal(http.StatusBadRequest))
		Expect(response.Body.String()).To(Equal(`{"error":"Invalid url parameter"}`))
	})

	It("should return an error if the API request fails", func() {
		mockClient.Transport = roundTripperFunc(func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("some unknown error")
		})

		handler := handlers.HelloThirdPartyApiHandler(mockClient)
		handler(ctx)

		Expect(response.Code).To(Equal(http.StatusInternalServerError))
	})
})
