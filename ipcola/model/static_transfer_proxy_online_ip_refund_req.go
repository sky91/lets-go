// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    staticTransferProxyOnlineIPRefundReq, err := UnmarshalStaticTransferProxyOnlineIPRefundReq(bytes)
//    bytes, err = staticTransferProxyOnlineIPRefundReq.Marshal()

package model

import "encoding/json"

func UnmarshalStaticTransferProxyOnlineIPRefundReq(data []byte) (StaticTransferProxyOnlineIPRefundReq, error) {
	var r StaticTransferProxyOnlineIPRefundReq
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StaticTransferProxyOnlineIPRefundReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type StaticTransferProxyOnlineIPRefundReq struct {
	// record id to refund
	ID int64 `json:"id"`
}
