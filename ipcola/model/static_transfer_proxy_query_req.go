// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    staticTransferProxyQueryReq, err := UnmarshalStaticTransferProxyQueryReq(bytes)
//    bytes, err = staticTransferProxyQueryReq.Marshal()

package model

import "encoding/json"

func UnmarshalStaticTransferProxyQueryReq(data []byte) (StaticTransferProxyQueryReq, error) {
	var r StaticTransferProxyQueryReq
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StaticTransferProxyQueryReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type StaticTransferProxyQueryReq struct {
	// query condition
	Filter StaticTransferProxyQueryReqFilter `json:"filter"`
	// from db or cache
	FromDB bool `json:"from_db"`
	// query quantity for pagination
	Limit *int64 `json:"limit,omitempty"`
	// query offset for pagination
	Offset *int64 `json:"offset,omitempty"`
}

// query condition
type StaticTransferProxyQueryReqFilter struct {
	// optional
	CountryCode *string `json:"country_code,omitempty"`
	// true for datacenter, false for residential
	IsDatacenter bool `json:"is_datacenter"`
	// optional, proxy ip
	NodeIpv4      *string  `json:"node_ipv4,omitempty"`
	NodeIpv4Array []string `json:"node_ipv4_array,omitempty"`

	Tag *string `json:"tag,omitempty"`
}
