package api

import (
	"fmt"
	"net/http"
)

func Api(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	fmt.Printf("email: %s", r.PostForm.Get("email"))
	fmt.Printf("password: %s", r.PostForm.Get("password"))
	fmt.Printf("cf-turnstile-response: %s", r.PostForm.Get("cf-turnstile-response"))

	_, _ = w.Write([]byte("ok"))
}
