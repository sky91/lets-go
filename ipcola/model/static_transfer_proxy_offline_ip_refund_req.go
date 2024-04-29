// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    staticTransferProxyOfflineIPRefundReq, err := UnmarshalStaticTransferProxyOfflineIPRefundReq(bytes)
//    bytes, err = staticTransferProxyOfflineIPRefundReq.Marshal()

package model

import "encoding/json"

func UnmarshalStaticTransferProxyOfflineIPRefundReq(data []byte) (StaticTransferProxyOfflineIPRefundReq, error) {
	var r StaticTransferProxyOfflineIPRefundReq
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StaticTransferProxyOfflineIPRefundReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type StaticTransferProxyOfflineIPRefundReq struct {
	// record id to refund
	ID int64 `json:"id"`
}
