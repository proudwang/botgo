package apitest

import (
	"testing"

	"github.com/proudwang/botgo/dto"
)

func TestMessageReaction(t *testing.T) {
	t.Run(
		"Create Message Reaction", func(t *testing.T) {
			err := api.CreateMessageReaction(ctx, testChannelID, testMessageID, dto.Emoji{Type: 1, ID: "43"})
			if err != nil {
				t.Error(err)
			}
			t.Logf("err:%+v", err)
		},
	)
	t.Run(
		"Delete Own Reaction", func(t *testing.T) {
			err := api.DeleteOwnMessageReaction(ctx, testChannelID, testMessageID, dto.Emoji{Type: 1, ID: "43"})
			if err != nil {
				t.Error(err)
			}
			t.Logf("err:%+v", err)
		},
	)
}
