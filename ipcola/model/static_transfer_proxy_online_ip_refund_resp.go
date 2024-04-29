// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    staticTransferProxyOnlineIPRefundResp, err := UnmarshalStaticTransferProxyOnlineIPRefundResp(bytes)
//    bytes, err = staticTransferProxyOnlineIPRefundResp.Marshal()

package model

import "encoding/json"

func UnmarshalStaticTransferProxyOnlineIPRefundResp(data []byte) (StaticTransferProxyOnlineIPRefundResp, error) {
	var r StaticTransferProxyOnlineIPRefundResp
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StaticTransferProxyOnlineIPRefundResp) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type StaticTransferProxyOnlineIPRefundResp struct {
	// expire time, unixtime
	ExpireUnixtime int64 `json:"expire_unixtime"`
	// success or other error msg
	MetaMessage string `json:"meta_message"`
	// 1 for success, negative for error
	MetaStatus int64 `json:"meta_status"`
	// refund balance
	RefundBalance int64 `json:"refund_balance"`
	// total capacity == old capacity + extra bought
	TrafficCap int64 `json:"traffic_cap"`
	// traffic bytes used
	TrafficUsed int64 `json:"traffic_used"`
}
