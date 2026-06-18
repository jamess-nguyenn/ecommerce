package lib

import (
	"ecommerce/helpers"
	"github.com/go-resty/resty/v2"
	"net/http"
	"strings"
)

type hTTPRequest struct {
	QueryParams map[string]string
	FormData    map[string]string
	PathParams  map[string]string
	Headers     map[string]string
	Body        any
}

func NewHTTPRequest() *hTTPRequest {
	return &hTTPRequest{}
}

func (req *hTTPRequest) SetQueryParams(queryParams map[string]string) *hTTPRequest {
	req.QueryParams = queryParams
	return req
}

func (req *hTTPRequest) SetFormData(formData map[string]string) *hTTPRequest {
	req.FormData = formData
	return req
}

func (req *hTTPRequest) SetPathParams(pathParams map[string]string) *hTTPRequest {
	req.PathParams = pathParams
	return req
}

func (req *hTTPRequest) SetHeaders(headers map[string]string) *hTTPRequest {
	req.Headers = headers
	return req
}

func (req *hTTPRequest) SetBody(body any) *hTTPRequest {
	req.Body = body
	return req
}

type hTTPResponse struct {
	IsSuccess  bool
	Data       any
	Message    string
	StatusCode int
	Headers    http.Header
}

func NewHTTPResponse() *hTTPResponse {
	return &hTTPResponse{}
}

func (resp *hTTPResponse) SetIsSuccess(isSuccess bool) *hTTPResponse {
	resp.IsSuccess = isSuccess
	return resp
}

func (resp *hTTPResponse) SetData(data any) *hTTPResponse {
	resp.Data = data
	return resp
}

func (resp *hTTPResponse) SetMessage(message string) *hTTPResponse {
	resp.Message = message
	return resp
}

func (resp *hTTPResponse) SetStatusCode(statusCode int) *hTTPResponse {
	resp.StatusCode = statusCode
	return resp
}

func (resp *hTTPResponse) SetHeaders(headers http.Header) *hTTPResponse {
	resp.Headers = headers
	return resp
}

func getMethod() string {
	functionInfo, _ := helpers.GetTrace(helpers.FrameLevel2)

	method := helpers.GetFunctionAfterDot(functionInfo)

	return strings.ToUpper(method)
}

func (req *hTTPRequest) Get(url string) *hTTPResponse {
	return req.execute(getMethod(), url)
}

func (req *hTTPRequest) Post(url string) *hTTPResponse {
	return req.execute(getMethod(), url)
}

func (req *hTTPRequest) newClient() *resty.Request {
	client := resty.New().NewRequest()

	if req.QueryParams != nil {
		client.SetQueryParams(req.QueryParams)
	}
	if req.FormData != nil {
		client.SetFormData(req.FormData)
	}
	if req.PathParams != nil {
		client.SetPathParams(req.PathParams)
	}
	if req.Headers != nil {
		client.SetHeaders(req.Headers)
	}
	if req.Body != nil {
		client.SetBody(req.Body)
	}

	return client
}

func (req *hTTPRequest) execute(method, url string) *hTTPResponse {
	resp := NewHTTPResponse()

	respClient, err := req.newClient().Execute(method, url)
	if err != nil {
		resp.SetIsSuccess(false).SetMessage(err.Error()).SetData(make(map[string]any))
	} else {
		resp.SetIsSuccess(respClient.IsSuccess()).SetStatusCode(respClient.StatusCode()).SetHeaders(respClient.Header())

		var data map[string]any

		if err = helpers.UnmarshalValue(respClient.Body(), &data); err != nil {
			resp.SetData(data)
		} else {
			resp.SetData(data["data"])
		}
	}

	helpers.LogInfo("Calling HTTP request", getLogClient(method, url, req, resp), helpers.GetAdditionalData(helpers.FrameLevel4))

	return resp
}

func getLogClient(method, url string, req *hTTPRequest, resp *hTTPResponse) map[string]any {
	return map[string]any{
		"Method":   method,
		"URL":      url,
		"Request":  *req,
		"Response": *resp,
	}
}
