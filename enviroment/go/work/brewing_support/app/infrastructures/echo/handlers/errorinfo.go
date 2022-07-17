package handlers

type Errorinfo struct {
	Message    string `json:"message"`
	Status     int    `json:"status"`
	StatusText string `json:"statusText"`
}

func newErrorInfo(message string, status int, statusTest string) *Errorinfo {

	return &Errorinfo{
		Message:    message,
		Status:     status,
		StatusText: statusTest,
	}
}
