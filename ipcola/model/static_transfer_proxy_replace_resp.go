// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    staticTransferProxyReplaceResp, err := UnmarshalStaticTransferProxyReplaceResp(bytes)
//    bytes, err = staticTransferProxyReplaceResp.Marshal()

package model

import "encoding/json"

func UnmarshalStaticTransferProxyReplaceResp(data []byte) (StaticTransferProxyReplaceResp, error) {
	var r StaticTransferProxyReplaceResp
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StaticTransferProxyReplaceResp) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type StaticTransferProxyReplaceResp struct {
	// success or other error msg
	MetaMessage string `json:"meta_message"`
	// 1 for success, negative for error
	MetaStatus int64  `json:"meta_status"`
	NewIP      string `json:"new_ip"`
	OldIP      string `json:"old_ip"`
}
