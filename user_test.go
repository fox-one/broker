package broker

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	b := New("fox223eef78c1e4bb1c", "2b92cdcf529931d80410d2a5919d9199")
	b.Debug()

	ctx := context.TODO()
	resp, err := b.Register(ctx, "yiplee", "yiplee", "https://resources.kumiclub.com/1fi7Dk/52c5c839a5f9ed6ae0f7f5a153607571")
	if assert.Nil(t, err) {
		assert.Len(t, resp.User.Id, 36)
		assert.NotEmpty(t, resp.Token)
		assert.NotZero(t, resp.Expire)
	}
}
