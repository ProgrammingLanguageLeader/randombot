package redis

type User struct {
	ID           int    `json:"id"`
	LanguageCode string `json:"languageCode"`
	State        string `json:"state"`
}
