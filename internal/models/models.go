package models

type (
	ExternalURL string
	InternalURL string
	ShortKey    string

	Config interface {
		Host() string
		Port() string
	}

	Storage interface {
		Set(url ExternalURL) (ShortKey, error)
		Get(key ShortKey) (ExternalURL, error)
	}
)
