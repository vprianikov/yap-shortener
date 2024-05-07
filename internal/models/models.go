package models

type (
	Config interface {
		Host() string
		Port() string
	}
)
