package accounts

type Interface interface {
	SaveNewAccount(id int64, countryID string) error
}
