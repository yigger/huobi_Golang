package market

import (
	"github.com/yigger/huobi_golang/pkg/response/base"
)

type SubscribeLast24hCandlestickResponse struct {
	base.WebSocketResponseBase
	Data *Candlestick
	Tick *Candlestick
}
