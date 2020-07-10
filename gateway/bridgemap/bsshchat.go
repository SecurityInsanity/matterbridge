// +build !nosshchat

package bridgemap

import (
	bsshchat "github.com/SecurityInsanity/matterbridge/bridge/sshchat"
)

func init() {
	FullMap["sshchat"] = bsshchat.New
}
