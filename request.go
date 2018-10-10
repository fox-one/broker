package broker

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/bitly/go-simplejson"
	uuid "github.com/satori/go.uuid"
)

func (b Broker) doRequest(req *http.Request) (*simplejson.Json, error) {
	resp, err := b.client.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := simplejson.NewFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return data.Get("data"), nil
	}

	if code := data.Get("code").MustInt(); code > 0 {
		return nil, Err{
			code: code,
			msg:  data.Get("msg").MustString(),
		}
	}

	return nil, errors.New(resp.Status)
}

func (b Broker) newRequest(method, endpoint string, paras ...interface{}) (*http.Request, error) {
	u, _ := url.Parse(b.baseApi())
	u.Path = path.Join(u.Path, endpoint)

	payload := map[string]interface{}{}
	for idx := 0; idx+1 < len(paras); idx += 2 {
		k, v := fmt.Sprintf("%v", paras[idx]), paras[idx+1]
		payload[k] = v
	}

	// auth
	ts := time.Now().Unix()
	nonce := uuid.Must(uuid.NewV4()).String()
	sign := b.Signature(ts, nonce)

	payload["appid"] = b.appid
	payload["timestamp"] = ts
	payload["nonce"] = nonce
	payload["sign"] = sign

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), bytes.NewReader(body))
	if err == nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, err
}

func (b Broker) do(ctx context.Context, method, endpoint string, paras ...interface{}) (*simplejson.Json, error) {
	req, err := b.newRequest(method, endpoint, paras...)

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	return b.doRequest(req)
}
