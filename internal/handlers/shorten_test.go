package handlers_test

import "net/http"

func (s *Suite) TestShorten() {
	tests := []Test{
		{
			Name: `positive test`,
			Send: Send{
				Method: http.MethodPost,
				Path:   `/`,
				Headers: map[string][]string{
					`Content-Type`: {
						`text/plain; charset=utf-8`,
					},
				},
				Body: []byte(`https://example.me`),
			},
			Want: Want{
				Code: http.StatusCreated,
				Headers: map[string][]string{
					`Content-Type`: {
						`text/plain; charset=utf-8`,
					},
				},
			},
		},
		{
			Name: `not allowed`,
			Send: Send{
				Method: http.MethodGet,
				Path:   `/`,
			},
			Want: Want{
				Code: http.StatusMethodNotAllowed,
			},
		},
		{
			Name: `unsupported media type`,
			Send: Send{
				Method: http.MethodPost,
				Path:   `/`,
				Headers: map[string][]string{
					`Content-Type`: {
						`application/json`,
					},
				},
			},
			Want: Want{
				Code: http.StatusBadRequest,
			},
		},
		{
			Name: `bad url`,
			Send: Send{
				Method: http.MethodPost,
				Path:   `/`,
				Headers: map[string][]string{
					`Content-Type`: {
						`text/plain; charset=utf-8`,
					},
				},
				Body: []byte(`https//example.me`),
			},
			Want: Want{
				Code: http.StatusBadRequest,
			},
		},
	}

	Run(s, tests)
}
