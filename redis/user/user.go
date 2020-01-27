package user

type User struct {
	ID              int      `json:"id"`
	LanguageCode    string   `json:"languageCode"`
	State           string   `json:"state"`
	Variants        []string `json:"variants"`
	MinRandomNumber int      `json:"minRandomNumber"`
	MaxRandomNumber int      `json:"maxRandomNumber"`
}
