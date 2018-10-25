package broker

import (
	"context"
	"encoding/json"
)

type Asset struct {
	AssetId  string `json:"assetId"`
	AssetKey string `json:"assetKey,omitempty"`
	ChainId  string `json:"chainId"`
	Icon     string `json:"icon"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	CoinId   uint   `json:"coinId,omitempty"`
}

type SnapshotUser struct {
	FoxId    uint   `json:"foxId,omitempty"`
	MixinId  string `json:"mixinId,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	Fullname string `json:"fullname,omitempty"`
}

type Snapshot struct {
	SnapshotId  string  `json:"snapshotId"`
	TraceId     string  `json:"traceId"`
	UserId      string  `json:"userId"`
	CreatedAt   int64   `json:"createdAt"`
	Source      string  `json:"source"`
	Amount      float64 `json:"amount"`
	InsideMixin bool    `json:"insideMixin"`
	Memo        string  `json:"memo"`

	Sender          string `json:"sender,omitempty"`
	Receiver        string `json:"receiver,omitempty"`
	TransactionHash string `json:"transactionHash,omitempty"`

	Asset    Asset        `json:"asset,omitempty"`
	Opponent SnapshotUser `json:"opponent,omitempty"`

	ExtraData map[string]interface{} `json:"extraData,omitempty"`
}

type Snapshots []Snapshot

type SnapshotResponse struct {
	Pagination OffsetPagination `json:"pagination"`
	Snapshots  Snapshots        `json:"snapshots"`
}

func (b Broker) PullSnapshots(ctx context.Context, userId, assetId, cursor string, asc bool) (*SnapshotResponse, error) {
	paras := []interface{}{}
	if len(userId) > 0 {
		paras = append(paras, "userId", userId)
	}

	if len(assetId) > 0 {
		paras = append(paras, "assetId", assetId)
	}

	if len(cursor) > 0 {
		paras = append(paras, "cursor", cursor)
	}

	if asc {
		paras = append(paras, "order", "ASC")
	} else {
		paras = append(paras, "order", "DESC")
	}

	resp, err := b.do(ctx, "POST", "broker/wallet/snapshots", paras...)
	if err != nil {
		return nil, err
	}

	data, err := resp.MarshalJSON()
	if err != nil {
		return nil, err
	}

	snapResp := &SnapshotResponse{}
	if err := json.Unmarshal(data, snapResp); err != nil {
		return nil, err
	}

	return snapResp, nil
}
