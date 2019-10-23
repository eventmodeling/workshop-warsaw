package router

import (
	"net/http"

	"github.com/eventmodeling/workshop-warsaw/register/app/feedback"
)

type postFeedback struct {
	Handler feedback.FeedbackHandler
}

func (h postFeedback) Handle(w http.ResponseWriter, r *http.Request) {
	return
}
