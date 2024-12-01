package errors

type ErrorMeta struct {
	HTTPCode int    `json:"-"`
	Message  string `json:"message"`
}
