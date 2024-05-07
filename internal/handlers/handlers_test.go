package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	"github.com/vprianikov/yap-shortener/internal/config"
	"github.com/vprianikov/yap-shortener/internal/handlers"
	"github.com/vprianikov/yap-shortener/internal/storage"
)

type (
	Suite struct {
		suite.Suite
		env handlers.Env
		r   *gin.Engine
	}

	Send struct {
		Method  string
		Path    string
		Headers http.Header
		Body    []byte
	}

	Want struct {
		Code    int
		Headers http.Header
		Body    []byte
	}

	Test struct {
		Name string
		Send Send
		Want Want
	}
)

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) SetupSuite() {
	co, errC := config.New()
	st, errS := storage.New()

	s.Require().NoError(errC, `failed to initialize the configuration`)
	s.Require().NoError(errS, `failed to initialize the storage`)

	s.env = handlers.Env{
		Config:  co,
		Storage: st,
	}

	r := gin.New()
	r.HandleMethodNotAllowed = true

	r.POST(`/`, s.env.Shorten)
	r.GET(`/:shortKey`, s.env.Expand)
	s.r = r
}

func setHeaders(r *http.Request, h http.Header) {
	for h, v := range h {
		for _, w := range v {
			r.Header.Add(h, w)
		}
	}
}

func checkHeaders(s *Suite, r *http.Response, hs http.Header) {
	for h, v := range hs {
		for _, w := range v {
			s.Contains(r.Header.Values(h), w)
		}
	}
}

func Run(s *Suite, tests []Test) {
	for _, test := range tests {
		s.Run(test.Name, func() {
			request := httptest.NewRequest(test.Send.Method, test.Send.Path, bytes.NewReader(test.Send.Body))
			setHeaders(request, test.Send.Headers)
			recorder := httptest.NewRecorder()

			s.r.ServeHTTP(recorder, request)

			result := recorder.Result()
			err := result.Body.Close()
			if err != nil {
				return
			}

			s.Equal(test.Want.Code, result.StatusCode)
			checkHeaders(s, result, test.Want.Headers)
		})
	}
}
