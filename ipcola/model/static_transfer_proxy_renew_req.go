// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    staticTransferProxyRenewReq, err := UnmarshalStaticTransferProxyRenewReq(bytes)
//    bytes, err = staticTransferProxyRenewReq.Marshal()

package model

import "encoding/json"

func UnmarshalStaticTransferProxyRenewReq(data []byte) (StaticTransferProxyRenewReq, error) {
	var r StaticTransferProxyRenewReq
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StaticTransferProxyRenewReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type StaticTransferProxyRenewReq struct {
	// record id to renew
	ID int64 `json:"id"`
	// true for datacenter, false for residential
	IsDatacenter bool `json:"is_datacenter"`
}
