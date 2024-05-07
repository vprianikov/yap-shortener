package models

type (
	ExternalURL string
	InternalURL string
	ShortKey    string

	Config interface {
		ServerAddress() string
		BaseURL() string
	}

	Storage interface {
		Set(url ExternalURL) (ShortKey, error)
		Get(key ShortKey) (ExternalURL, error)
	}
)
