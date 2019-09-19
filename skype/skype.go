package skype

import (
	"net/http"

	"github.com/drzhnin/go-skype/skype"
	"github.com/lbryio/lbry.go/extras/api"
)

func SendMessage(r *http.Request) api.Response {
	client := skype.NewClient("df0abf4f-7fa1-4942-9017-f27dceebf432", "C.AZFqhLqR6gR37Cvv]x4+c?rl[=FXyP")
	_, err := client.Authorization.Authorize()
	if err != nil {
		return api.Response{Error: err}
	}
	_, err = client.Messege.Send("LXib4WRXjbOy", "message/text", "Hello World")
	if err != nil {
		return api.Response{Error: err}
	}

	return api.Response{Data: "ok"}
}
