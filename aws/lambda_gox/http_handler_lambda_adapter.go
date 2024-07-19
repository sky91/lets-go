package lambda_gox

import (
	"bytes"
	"context"
	"encoding/base64"
	"github.com/aws/aws-lambda-go/events"
	"io"
	"net/http"
	"strings"
)

type bufferedHttpRespWriter struct {
	statusCode int
	header     http.Header
	body       bytes.Buffer
}

func (thisP *bufferedHttpRespWriter) Header() http.Header {
	return thisP.header
}

func (thisP *bufferedHttpRespWriter) Write(b []byte) (int, error) {
	return thisP.body.Write(b)
}

func (thisP *bufferedHttpRespWriter) WriteHeader(statusCode int) {
	thisP.statusCode = statusCode
}

type HttpHandlerToLambdaAdapter struct {
	HttpHandler http.Handler
}

func (thisV HttpHandlerToLambdaAdapter) HandleHttpApi(ctx context.Context, request *events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
	body := io.Reader(strings.NewReader(request.Body))
	if request.IsBase64Encoded {
		body = base64.NewDecoder(base64.StdEncoding, body)
	}
	url := "https://" + request.RequestContext.DomainName + request.RawPath
	if request.RawQueryString != "" {
		url += "?" + request.RawQueryString
	}
	ctx = context.WithValue(ctx, ReqCtxKey{}, request)
	httpRequest, err := http.NewRequestWithContext(ctx, request.RequestContext.HTTP.Method, url, body)
	if err != nil {
		return nil, err
	}
	httpRequest.RemoteAddr = request.RequestContext.HTTP.SourceIP
	for k, v := range request.Headers {
		httpRequest.Header.Add(k, v)
	}

	respWriter := bufferedHttpRespWriter{header: make(http.Header, 32)}
	respWriter.body.Grow(8192)

	thisV.HttpHandler.ServeHTTP(&respWriter, httpRequest)

	resp := &events.APIGatewayV2HTTPResponse{
		StatusCode:      respWriter.statusCode,
		Body:            base64.StdEncoding.EncodeToString(respWriter.body.Bytes()),
		IsBase64Encoded: true}

	if hdrCnt := len(respWriter.header); hdrCnt > 0 {
		resp.Headers = make(map[string]string, hdrCnt)
		for k, v := range respWriter.header {
			if k == "Set-Cookie" {
				resp.Cookies = v
			} else {
				resp.Headers[k] = strings.Join(v, ",")
			}
		}
	}
	return resp, nil
}

func (thisV HttpHandlerToLambdaAdapter) HandleFunctionUrl(ctx context.Context, request *events.LambdaFunctionURLRequest) (*events.LambdaFunctionURLResponse, error) {
	body := io.Reader(strings.NewReader(request.Body))
	if request.IsBase64Encoded {
		body = base64.NewDecoder(base64.StdEncoding, body)
	}
	url := "https://" + request.RequestContext.DomainName + request.RawPath
	if request.RawQueryString != "" {
		url += "?" + request.RawQueryString
	}
	ctx = context.WithValue(ctx, ReqCtxKey{}, request)
	httpRequest, err := http.NewRequestWithContext(ctx, request.RequestContext.HTTP.Method, url, body)
	if err != nil {
		return nil, err
	}
	httpRequest.RemoteAddr = request.RequestContext.HTTP.SourceIP
	for k, v := range request.Headers {
		httpRequest.Header.Add(k, v)
	}

	respWriter := bufferedHttpRespWriter{header: make(http.Header, 32)}
	respWriter.body.Grow(8192)

	thisV.HttpHandler.ServeHTTP(&respWriter, httpRequest)

	resp := &events.LambdaFunctionURLResponse{
		StatusCode:      respWriter.statusCode,
		Body:            base64.StdEncoding.EncodeToString(respWriter.body.Bytes()),
		IsBase64Encoded: true}

	if hdrCnt := len(respWriter.header); hdrCnt > 0 {
		resp.Headers = make(map[string]string, hdrCnt)
		for k, v := range respWriter.header {
			if k == "Set-Cookie" {
				resp.Cookies = v
			} else {
				resp.Headers[k] = strings.Join(v, ",")
			}
		}
	}
	return resp, nil
}

type ReqCtxKey struct{}
