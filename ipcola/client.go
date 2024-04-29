package ipcola

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/sky91/lets-go/gox/errorx"
	"github.com/sky91/lets-go/ipcola/model"
	"io"
	"net/http"
)

type Client struct {
	HttpClient *http.Client
	AuthToken  string
}

func (thisP *Client) QueryStaticTransferProxy(ctx context.Context, req *model.StaticTransferProxyQueryReq) (*model.StaticTransferProxyQueryResp, error) {
	return ipColaDoHttp[model.StaticTransferProxyQueryReq, model.StaticTransferProxyQueryResp](
		thisP, ctx, http.MethodPost, apiUrlQueryStaticTransferProxy, req)
}

func (thisP *Client) BuyStaticTransferProxy(ctx context.Context, req *model.StaticTransferProxyBuyReq) (*model.StaticTransferProxyBuyResp, error) {
	return ipColaDoHttp[model.StaticTransferProxyBuyReq, model.StaticTransferProxyBuyResp](
		thisP, ctx, http.MethodPost, apiUrlBuyStaticTransferProxy, req)
}

func (thisP *Client) ReplaceStaticTransferProxy(ctx context.Context, req *model.StaticTransferProxyReplaceReq) (*model.StaticTransferProxyReplaceResp, error) {
	return ipColaDoHttp[model.StaticTransferProxyReplaceReq, model.StaticTransferProxyReplaceResp](
		thisP, ctx, http.MethodPost, apiUrlReplaceStaticTransferProxy, req)
}

func (thisP *Client) OfflineIpRefundStaticTransferProxy(ctx context.Context, id int64) (*model.StaticTransferProxyOfflineIPRefundResp, error) {
	return ipColaDoHttp[model.StaticTransferProxyOfflineIPRefundReq, model.StaticTransferProxyOfflineIPRefundResp](
		thisP, ctx, http.MethodPost, apiUrlOfflineIpRefundStaticTransferProxy, &model.StaticTransferProxyOfflineIPRefundReq{ID: id})
}

func (thisP *Client) OnlineIpRefundStaticTransferProxy(ctx context.Context, req *model.StaticTransferProxyOnlineIPRefundReq) (*model.StaticTransferProxyOnlineIPRefundResp, error) {
	return ipColaDoHttp[model.StaticTransferProxyOnlineIPRefundReq, model.StaticTransferProxyOnlineIPRefundResp](
		thisP, ctx, http.MethodPost, apiUrlOnlineIpRefundStaticTransferProxy, req)
}

func ipColaDoHttp[Req, Resp any](client *Client, ctx context.Context, method, url string, req *Req) (*Resp, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "Marshal() error")
	}
	httpReq, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(reqBody))
	if err != nil {
		return nil, errors.Wrap(err, "NewRequestWithContext() error")
	}
	httpReq.Header.Set("Authorization", "Bearer "+client.AuthToken)
	httpReq.Header.Add("Content-Type", contentTypeApplicationJson)

	httpResp, err := client.HttpClient.Do(httpReq)
	if err != nil {
		return nil, errors.Wrap(err, "do req error")
	}
	defer errorx.DoIgnoreError(httpResp.Body.Close)
	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "ReadAll() error")
	}
	resp := new(Resp)
	if err = json.Unmarshal(respBody, resp); err != nil {
		return nil, errors.Wrap(err, "Unmarshal() error")
	}
	return resp, nil
}

const (
	contentTypeApplicationJson = "application/json"

	apiUrlQueryStaticTransferProxy           = "https://api.ipcola.com/api/static_transfer_proxy/query"
	apiUrlBuyStaticTransferProxy             = "https://api.ipcola.com/api/static_transfer_proxy/buy"
	apiUrlRenewStaticTransferProxy           = "https://api.ipcola.com/api/static_transfer_proxy/renew"
	apiUrlReplaceStaticTransferProxy         = "https://api.ipcola.com/api/static_transfer_proxy/replace_ip"
	apiUrlOfflineIpRefundStaticTransferProxy = "https://api.ipcola.com/api/static_transfer_proxy/refund"
	apiUrlOnlineIpRefundStaticTransferProxy  = "https://api.ipcola.com/api/static_transfer_proxy/online_ip_refund"

	apiUrlDeleteRecordStaticTransferProxy = "https://api.ipcola.com/api/static_transfer_proxy/delete_record"
	// {"id":32871}
	// {"meta_status":1,"meta_message":"success"}
)
