package api

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func Api(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	logrus.Info("email: %s", r.PostForm.Get("email"))
	logrus.Info("password: %s", r.PostForm.Get("password"))
	logrus.Info("cf-turnstile-response: %s", r.PostForm.Get("cf-turnstile-response"))

	_, _ = w.Write([]byte("ok"))
}
