package api

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	_ = r.ParseForm()
	if len(r.PostForm.Get("email")) < 1 {
		_, _ = w.Write(output(500, "email not empty!"))
		return
	}

	if len(r.PostForm.Get("password")) < 1 {
		_, _ = w.Write(output(500, "password not empty!"))
		return
	}

	if len(r.PostForm.Get("cf-turnstile-response")) < 1 {
		_, _ = w.Write(output(500, "cf-turnstile-response not empty!"))
		return
	}

	body := `{"secret":"` + os.Getenv("secret") + `", "response":"` + r.PostForm.Get("cf-turnstile-response") + `"}`
	re, _ := http.NewRequest(http.MethodPost, "https://challenges.cloudflare.com/turnstile/v0/siteverify", strings.NewReader(body))
	re.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(re)
	defer res.Body.Close()
	if err != nil {
		_, _ = w.Write(output(500, "http client fail: "+err.Error()))
		return
	}

	b, _ := ioutil.ReadAll(res.Body)
	logrus.Info("turnstile response " + string(b))

	response := struct {
		Success     bool     `json:"success"`
		ErrorCodes  []string `json:"error-codes"`
		ChallengeTs string   `json:"challenge_ts"`
		Hostname    string   `json:"hostname"`
	}{}

	_ = json.Unmarshal(b, &response)
	if response.Success != true {
		_, _ = w.Write(output(500, "验证失败"))
		return
	}

	_, _ = w.Write(output(200, "验证通过"))
}

type Output struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func output(code int, msg string) []byte {
	out := &Output{
		Code: code,
		Msg:  msg,
	}

	b, _ := json.Marshal(out)
	return b
}
