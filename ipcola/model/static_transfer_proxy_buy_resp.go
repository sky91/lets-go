// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    staticTransferProxyBuyResp, err := UnmarshalStaticTransferProxyBuyResp(bytes)
//    bytes, err = staticTransferProxyBuyResp.Marshal()

package model

import "encoding/json"

func UnmarshalStaticTransferProxyBuyResp(data []byte) (StaticTransferProxyBuyResp, error) {
	var r StaticTransferProxyBuyResp
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StaticTransferProxyBuyResp) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type StaticTransferProxyBuyResp struct {
	// return proxy ip if success
	AllocIPList []string `json:"alloc_ip_list"`
	// success or other error msg
	MetaMessage string `json:"meta_message"`
	// 1 for success, negative for error
	MetaStatus int64 `json:"meta_status"`
}
