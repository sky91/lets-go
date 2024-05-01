package model

type StaticTransferProxyUpdateTagReq struct {
	Ids []int64 `json:"ids"`
	Tag string  `json:"tag"`
}
