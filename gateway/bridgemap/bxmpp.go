// +build !noxmpp

package bridgemap

import (
	bxmpp "github.com/SecurityInsanity/matterbridge/bridge/xmpp"
)

func init() {
	FullMap["xmpp"] = bxmpp.New
}
