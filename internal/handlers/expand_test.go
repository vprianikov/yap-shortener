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
				Path:   `/` + string(key),
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
				Path:   `/gQXr0VLO`,
			},
			Want: Want{
				Code: http.StatusNotFound,
			},
		},
		{
			Name: `incorrect key format`,
			Send: Send{
				Method: http.MethodGet,
				Path:   `/gQXr0VL!`,
			},
			Want: Want{
				Code: http.StatusNotFound,
			},
		},
		{
			Name: `not allowed`,
			Send: Send{
				Method: http.MethodPost,
				Path:   `/gQXr0VLO`,
			},
			Want: Want{
				Code: http.StatusMethodNotAllowed,
			},
		},
	}

	Run(s, tests)
}
