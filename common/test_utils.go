package common

import (
	"testing"
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)


type RequestTest struct {
	Init           func(*http.Request)
	Url            string
	Method         string
	BodyData       string
	ExpectedCode   int
	ResponseRegexg string
	Msg            string
}

func ExecuteTests(t *testing.T, r *gin.Engine, requestTests []RequestTest) {
	asserts := assert.New(t)

	for _, testData := range requestTests {
		bodyData := testData.BodyData
		req, err := http.NewRequest(testData.Method, testData.Url, bytes.NewBufferString(bodyData))
		req.Header.Set("Content-Type", "application/json")
		asserts.NoError(err)

		testData.Init(req)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		asserts.Equal(testData.ExpectedCode, w.Code, "Response Status - "+testData.Msg)
		asserts.Regexp(testData.ResponseRegexg, w.Body.String(), "Response Content - "+testData.Msg)
	}
}
