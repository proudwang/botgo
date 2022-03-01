package apitest

import (
	"strconv"
	"testing"
	"time"

	"github.com/proudwang/botgo/dto"
)

var (
	testMuteGuildID = "3326534247441079828" // replace your guild id
	testMuteUserID  = "144111002883982087"  // replace your user id
)

func Test_mute(t *testing.T) {
	t.Run(
		"频道禁言", func(t *testing.T) {
			mute := &dto.UpdateGuildMute{
				MuteEndTimestamp: strconv.FormatInt(time.Now().Unix()+600, 10),
			}
			err := api.GuildMute(ctx, testMuteGuildID, mute)
			if err != nil {
				t.Error(err)
			}
			t.Logf("Testing_Succ")
		},
	)
	t.Run(
		"频道指定成员禁言", func(t *testing.T) {
			mute := &dto.UpdateGuildMute{
				MuteEndTimestamp: strconv.FormatInt(time.Now().Unix()+600, 10),
			}
			err := api.MemberMute(ctx, testGuildID, testMuteUserID, mute)
			if err != nil {
				t.Error(err)
			}
			t.Logf("Testing_Succ")
		},
	)
}
