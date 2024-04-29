// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    staticTransferProxyReplaceReq, err := UnmarshalStaticTransferProxyReplaceReq(bytes)
//    bytes, err = staticTransferProxyReplaceReq.Marshal()

package model

import "encoding/json"

func UnmarshalStaticTransferProxyReplaceReq(data []byte) (StaticTransferProxyReplaceReq, error) {
	var r StaticTransferProxyReplaceReq
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StaticTransferProxyReplaceReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type StaticTransferProxyReplaceReq struct {
	// change location or not
	ChangeLocation bool `json:"change_location"`
	// new location code
	CountryCode *string `json:"country_code"`
	// record id to refund
	ID int64 `json:"id"`
}
