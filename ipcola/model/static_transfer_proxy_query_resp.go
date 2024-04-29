// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    staticTransferProxyQueryResp, err := UnmarshalStaticTransferProxyQueryResp(bytes)
//    bytes, err = staticTransferProxyQueryResp.Marshal()

package model

import "encoding/json"

func UnmarshalStaticTransferProxyQueryResp(data []byte) (StaticTransferProxyQueryResp, error) {
	var r StaticTransferProxyQueryResp
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StaticTransferProxyQueryResp) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type StaticTransferProxyQueryResp struct {
	// current time stamp from server
	CurrentUnixtime int64 `json:"current_unixtime"`
	// query result
	Data []StaticTransferProxyQueryRespDatum `json:"data"`
	// success or other error msg
	MetaMessage string `json:"meta_message"`
	// 1 for success, negative for error
	MetaStatus int64 `json:"meta_status"`
	// total record quantity
	TotalCount int64 `json:"total_count"`
}

// static direct proxy info
type StaticTransferProxyQueryRespDatum struct {
	// auth
	AccountName *string `json:"account_name,omitempty"`
	// auth password
	AccountPassword *string `json:"account_password,omitempty"`
	City            *string `json:"city,omitempty"`
	ContinentCode   *string `json:"continent_code,omitempty"`
	CountryCode     *string `json:"country_code,omitempty"`
	CreatedUnixtime *int64  `json:"created_unixtime,omitempty"`
	// a unix timestamp
	ExpireUnixtime *int64 `json:"expire_unixtime,omitempty"`
	Forbidden      *bool  `json:"forbidden,omitempty"`
	ID             *int64 `json:"id,omitempty"`
	// ip array
	IPWhitelist           []string `json:"ip_whitelist,omitempty"`
	IPWhitelistCountLimit *int64   `json:"ip_whitelist_count_limit,omitempty"`
	IsDatacenter          *bool    `json:"is_datacenter,omitempty"`
	IsOnline              *bool    `json:"is_online,omitempty"`
	IsSatellite           *bool    `json:"is_satellite,omitempty"`
	// status, not in db, get value when query
	IsStable *bool   `json:"is_stable,omitempty"`
	ISP      *string `json:"isp,omitempty"`
	// ip info
	NodeIpv4 *string `json:"node_ipv4,omitempty"`
	Port     *int64  `json:"port,omitempty"`
	// if >0 means refunded
	RefundUnixtime *int64  `json:"refund_unixtime,omitempty"`
	Region         *string `json:"region,omitempty"`
	// update flag. set to true if any update happened, set to false when update send to node
	ToUpdate *bool `json:"to_update,omitempty"`
	// total capacity == old capacity + new bought
	TrafficCap *int64 `json:"traffic_cap,omitempty"`
	// traffic
	TrafficUsed *int64  `json:"traffic_used,omitempty"`
	UserID      *int64  `json:"user_id,omitempty"`
	UserTag     *string `json:"user_tag,omitempty"`

	Password *string `json:"password,omitempty"`
}
