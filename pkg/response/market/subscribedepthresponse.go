package market

import (
	"github.com/yigger/huobi_golang/pkg/response/base"
)

type SubscribeDepthResponse struct {
	base.WebSocketResponseBase
	Data *Depth
	Tick *Depth
}
