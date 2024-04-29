// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    staticTransferProxyRenewResp, err := UnmarshalStaticTransferProxyRenewResp(bytes)
//    bytes, err = staticTransferProxyRenewResp.Marshal()

package model

import "encoding/json"

func UnmarshalStaticTransferProxyRenewResp(data []byte) (StaticTransferProxyRenewResp, error) {
	var r StaticTransferProxyRenewResp
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StaticTransferProxyRenewResp) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type StaticTransferProxyRenewResp struct {
	// success or other error msg
	MetaMessage string `json:"meta_message"`
	// 1 success, negative for error
	MetaStatus int64 `json:"meta_status"`
}
