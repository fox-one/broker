package broker

import (
	"context"
	"encoding/json"
)

type User struct {
	Id     string `json:"fox_id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type UserResponse struct {
	User   User   `json:"user"`
	Token  string `json:"token"`
	Expire int64  `json:"expires_in"`
}

func (b Broker) Register(ctx context.Context, traceId, name, avatar string) (*UserResponse, error) {
	paras := []interface{}{
		"trace_id", traceId,
		"name", name,
		"avatar", avatar,
	}

	resp, err := b.do(ctx, "POST", "broker/register", paras...)
	if err != nil {
		return nil, err
	}

	data, err := resp.MarshalJSON()
	if err != nil {
		return nil, err
	}

	userResp := &UserResponse{}
	if err := json.Unmarshal(data, userResp); err != nil {
		return nil, err
	}

	return userResp, nil
}
