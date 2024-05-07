package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/vprianikov/yap-shortener/internal/config"
	"github.com/vprianikov/yap-shortener/internal/handlers"
	"github.com/vprianikov/yap-shortener/internal/storage"
)

type (
	Suite struct {
		suite.Suite
		env handlers.Env
	}

	Path struct {
		Name  string
		Value string
	}

	Send struct {
		Method  string
		Path    []Path
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
}

func setPath(r *http.Request, p []Path) {
	for _, p := range p {
		r.SetPathValue(p.Name, p.Value)
	}
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

func Run(s *Suite, handler func(http.ResponseWriter, *http.Request), tests []Test) {
	for _, test := range tests {
		s.Run(test.Name, func() {
			request := httptest.NewRequest(test.Send.Method, `/`, bytes.NewReader(test.Send.Body))
			setPath(request, test.Send.Path)
			setHeaders(request, test.Send.Headers)
			recorder := httptest.NewRecorder()

			handler(recorder, request)

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
