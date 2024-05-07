package handlers_test

import (
	"net/http"
)

func (s *Suite) TestExpand() {
	key, err := s.env.Storage.Set(`https://practicum.yandex.ru/`)
	s.Require().NoError(err, `failed to set`)

	tests := []Test{
		{
			Name: `positive test`,
			Send: Send{
				Method: http.MethodGet,
				Path: []Path{
					{
						Name:  `shortKey`,
						Value: string(key),
					},
				},
			},
			Want: Want{
				Code: http.StatusTemporaryRedirect,
				Headers: map[string][]string{
					`Location`: {
						`https://practicum.yandex.ru/`,
					},
				},
			},
		},
		{
			Name: `not found`,
			Send: Send{
				Method: http.MethodGet,
				Path: []Path{
					{
						Name:  `key`,
						Value: `gQXr0VLO`,
					},
				},
			},
			Want: Want{
				Code: http.StatusNotFound,
			},
		},
		{
			Name: `incorrect key format`,
			Send: Send{
				Method: http.MethodGet,
				Path: []Path{
					{
						Name:  `key`,
						Value: `gQXr0VL!`,
					},
				},
			},
			Want: Want{
				Code: http.StatusNotFound,
			},
		},
		{
			Name: `not allowed`,
			Send: Send{
				Method: http.MethodPost,
			},
			Want: Want{
				Code: http.StatusMethodNotAllowed,
			},
		},
	}

	Run(s, s.env.Expand, tests)
}
