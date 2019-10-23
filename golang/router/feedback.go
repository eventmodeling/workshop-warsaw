package router

import (
	"errors"
	"log"
	"net/http"

	"github.com/eventmodeling/workshop-warsaw/register/app/feedback"
)

type FeedbackSection struct {
    Name string
    Slug string
}

type getFeedback struct {
}

func (h getFeedback) Handle(w http.ResponseWriter, r *http.Request) {
	data := []FeedbackSection{
	    {"What's your opinion on shop cleanest?", "cleanest"},
	    {"What's your overal experience with the shop?", "exp"},
	    {"How easy was to find required goods?", "find"},
	}

	tpl := renderTemplate("feedback", data)
	tpl(w, r)
}

type postFeedback struct {
	Handler feedback.FeedbackHandler
}

func (h postFeedback) Handle(w http.ResponseWriter, r *http.Request) {
	cmd, err := parseFeedbackData(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	err = h.Handler.Execute(cmd)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_, _ = w.Write([]byte("success"))
}

func parseFeedbackData(r *http.Request) (feedback.Feedback, error) {
	err := r.ParseForm()
	if err != nil {
		return feedback.Feedback{}, err
	}

	purchaseID := r.PostForm.Get("purchase_id")
	if purchaseID == "" {
		return feedback.Feedback{}, errors.New("missing purchase ID")
	}

	return feedback.Feedback{
		PurchaseID:  purchaseID,
		Cleanliness: levelFromRequest(r, "cleanliness"),
		Experience:  levelFromRequest(r, "experience"),
		Queue:       levelFromRequest(r, "queue"),
		EasyToFind:  levelFromRequest(r, "easy_to_find"),
		Comments:    r.PostForm.Get("comments"),
	}, nil
}

func levelFromRequest(r *http.Request, key string) feedback.Level {
	level, err := feedback.NewLevel(r.PostForm.Get(key))
	if err != nil {
		log.Println(err)
		return feedback.Level{}
	}
	return level
}
