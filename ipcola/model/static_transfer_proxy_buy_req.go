// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    staticTransferProxyBuyReq, err := UnmarshalStaticTransferProxyBuyReq(bytes)
//    bytes, err = staticTransferProxyBuyReq.Marshal()

package model

import "encoding/json"

func UnmarshalStaticTransferProxyBuyReq(data []byte) (StaticTransferProxyBuyReq, error) {
	var r StaticTransferProxyBuyReq
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StaticTransferProxyBuyReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type StaticTransferProxyBuyReq struct {
	// always 1
	Count int64 `json:"count"`
	// if no specific country, country_code=""
	CountryCode *string `json:"country_code,omitempty"`
	// true for datacenter, false for residential
	IsDatacenter bool `json:"is_datacenter"`
	// always "ipcola"
	UserTag string `json:"user_tag"`

	Tag *string `json:"tag,omitempty"`
}
