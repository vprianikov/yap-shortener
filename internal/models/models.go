package models

type (
	ExternalURL string
	InternalURL string
	ShortKey    string

	// TODO(SSH): интерфейсы определяются там, где они используются
	Config interface {
		ServerAddress() string
		BaseURL() string
	}

	Storage interface {
		Set(url ExternalURL) (ShortKey, error)
		Get(key ShortKey) (ExternalURL, error)
	}
)
