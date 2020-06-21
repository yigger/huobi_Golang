package auth

import (
	"encoding/json"

	"github.com/yigger/huobi_golang/pkg/response/base"
)

type WebSocketV2AuthenticationResponse struct {
	base.WebSocketV2ResponseBase
}

func ParseWSV2AuthResp(message string) *WebSocketV2AuthenticationResponse {
	result := &WebSocketV2AuthenticationResponse{}
	err := json.Unmarshal([]byte(message), result)
	if err != nil {
		return nil
	}

	return result
}
