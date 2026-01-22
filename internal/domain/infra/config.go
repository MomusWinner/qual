package infra

type Config interface {
	GetHost() string
	GetDatabaseURL() string
}
