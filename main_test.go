package main
import (
  "encoding/json"
  "net/http"
  "net/http/httptest"
  "strings"
  "net/url"
  "testing"
  "github.com/gin-gonic/gin"
  "github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string, params url.Values) *httptest.ResponseRecorder {
  req, _ :=  http.NewRequest(method, path, strings.NewReader(params.Encode()))
  req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

  w := httptest.NewRecorder()
  r.ServeHTTP(w, req)
  return w
}

func TestCreateProposal(t *testing.T) {
  body := gin.H{
    "message": "Proposal created successfully",
  }

  data := url.Values{}
  data.Set("name", "五石凪")
  data.Add("email", "goishi@naggi.club")
  data.Add("programming_experience", "有")
  data.Add("content", "契約がしたいです")

  router := SetupRouter()
  w := performRequest(router, "POST", "api/proposals", data)

  assert.Equal(t, http.StatusCreated, w.Code)

  var response map[string]interface{}
  err := json.Unmarshal([]byte(w.Body.String()), &response)
  value, exists := response["message"]

  assert.Nil(t, err)
  assert.True(t, exists)
  assert.Equal(t, body["message"], value)
}
