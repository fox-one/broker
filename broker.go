package broker

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

type Broker struct {
	appid     string
	appsecret string
	debug     bool
	client    *http.Client
}

func New(id, secret string) Broker {
	return Broker{
		appid:     id,
		appsecret: secret,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (b Broker) IsDebug() bool {
	return b.debug
}

func (b *Broker) Debug() {
	b.debug = true
}

func (b Broker) baseApi() string {
	if b.debug {
		return "https://dev.fox.one/api"
	}

	return "https://api.fox.one/api"
}

func (b Broker) Signature(ts int64, nonce string) string {
	h := hmac.New(sha256.New, []byte(b.appsecret))
	payload := fmt.Sprintf("appid=%s&ts=%d&nonce=%s", b.appid, ts, nonce)
	h.Write([]byte(payload))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
