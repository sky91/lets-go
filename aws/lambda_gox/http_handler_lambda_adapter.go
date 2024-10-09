package lambda_gox

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"io"
	"net/http"
	"net/url"
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

func (thisV HttpHandlerToLambdaAdapter) HandleHttpApi(ctx context.Context, request *events.APIGatewayV2HTTPRequest) (response *events.APIGatewayV2HTTPResponse, err error) {
	httpReq := commonHttpReq.WithContext(context.WithValue(ctx, ReqCtxKey{}, request))
	httpReq.Method = request.RequestContext.HTTP.Method
	httpReq.Proto = request.RequestContext.HTTP.Protocol

	var ok bool
	if httpReq.ProtoMajor, httpReq.ProtoMinor, ok = http.ParseHTTPVersion(httpReq.Proto); !ok {
		return nil, fmt.Errorf("malformed http proto: [%s]", httpReq.Proto)
	}

	httpReq.RequestURI = request.RawPath
	if request.RawQueryString != "" {
		httpReq.RequestURI += "?" + request.RawQueryString
	}
	if httpReq.URL, err = url.ParseRequestURI(httpReq.RequestURI); err != nil {
		return nil, err
	}

	httpReq.Header = make(http.Header, len(request.Headers))
	for k, v := range request.Headers {
		httpReq.Header.Add(k, v)
	}
	httpReq.Host = httpReq.Header.Get("Host")

	body := io.Reader(strings.NewReader(request.Body))
	if request.IsBase64Encoded {
		body = base64.NewDecoder(base64.StdEncoding, body)
		httpReq.ContentLength = int64(base64.StdEncoding.DecodedLen(len(request.Body)))
	} else {
		httpReq.ContentLength = int64(len(request.Body))
	}
	httpReq.Body = io.NopCloser(body)

	respWriter := bufferedHttpRespWriter{header: make(http.Header, 32)}
	respWriter.body.Grow(8192)

	thisV.HttpHandler.ServeHTTP(&respWriter, httpReq)

	response = &events.APIGatewayV2HTTPResponse{
		StatusCode:      respWriter.statusCode,
		Body:            base64.StdEncoding.EncodeToString(respWriter.body.Bytes()),
		IsBase64Encoded: true,
	}

	if hdrCnt := len(respWriter.header); hdrCnt > 0 {
		response.Headers = make(map[string]string, hdrCnt)
		for k, v := range respWriter.header {
			if k == "Set-Cookie" {
				response.Cookies = v
			} else {
				response.Headers[k] = strings.Join(v, ",")
			}
		}
	}
	return response, nil
}

func (thisV HttpHandlerToLambdaAdapter) HandleFunctionUrl(ctx context.Context, request *events.LambdaFunctionURLRequest) (response *events.LambdaFunctionURLResponse, err error) {
	httpReq := commonHttpReq.WithContext(context.WithValue(ctx, ReqCtxKey{}, request))
	httpReq.Method = request.RequestContext.HTTP.Method
	httpReq.Proto = request.RequestContext.HTTP.Protocol

	var ok bool
	if httpReq.ProtoMajor, httpReq.ProtoMinor, ok = http.ParseHTTPVersion(httpReq.Proto); !ok {
		return nil, fmt.Errorf("malformed http proto: [%s]", httpReq.Proto)
	}

	httpReq.RequestURI = request.RawPath
	if request.RawQueryString != "" {
		httpReq.RequestURI += "?" + request.RawQueryString
	}
	if httpReq.URL, err = url.ParseRequestURI(httpReq.RequestURI); err != nil {
		return nil, err
	}

	httpReq.Header = make(http.Header, len(request.Headers))
	for k, v := range request.Headers {
		httpReq.Header.Add(k, v)
	}
	httpReq.Host = httpReq.Header.Get("Host")

	body := io.Reader(strings.NewReader(request.Body))
	if request.IsBase64Encoded {
		body = base64.NewDecoder(base64.StdEncoding, body)
		httpReq.ContentLength = int64(base64.StdEncoding.DecodedLen(len(request.Body)))
	} else {
		httpReq.ContentLength = int64(len(request.Body))
	}
	httpReq.Body = io.NopCloser(body)

	respWriter := bufferedHttpRespWriter{header: make(http.Header, 32)}
	respWriter.body.Grow(8192)

	thisV.HttpHandler.ServeHTTP(&respWriter, httpReq)

	response = &events.LambdaFunctionURLResponse{
		StatusCode:      respWriter.statusCode,
		Body:            base64.StdEncoding.EncodeToString(respWriter.body.Bytes()),
		IsBase64Encoded: true,
	}

	if hdrCnt := len(respWriter.header); hdrCnt > 0 {
		response.Headers = make(map[string]string, hdrCnt)
		for k, v := range respWriter.header {
			if k == "Set-Cookie" {
				response.Cookies = v
			} else {
				response.Headers[k] = strings.Join(v, ",")
			}
		}
	}
	return response, nil
}

type ReqCtxKey struct{}

var commonHttpReq = http.Request{Close: true}
