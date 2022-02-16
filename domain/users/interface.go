package users

type Interface interface {
	ReadUsersFeed(message UserMsg) error
}
