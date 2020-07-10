// +build !noapi

package bridgemap

import (
	"github.com/SecurityInsanity/matterbridge/bridge/api"
)

func init() {
	FullMap["api"] = api.New
}
