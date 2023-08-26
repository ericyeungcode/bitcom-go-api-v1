package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	StatusCode int
	Buffer     []byte
}

func ResponseToString(rsp *Response) string {
	if rsp == nil {
		return "null"
	}
	return fmt.Sprintf("[status_code=%v, buffer=%v]", rsp.StatusCode, string(rsp.Buffer))
}

type RawResponse struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    *json.RawMessage `json:"data"`
}

func ParseCommonPayload(buf []byte, v interface{}) error {
	var rawResp RawResponse

	if err := json.Unmarshal(buf, &rawResp); err != nil {
		return fmt.Errorf("ParseCommonPayload fail to unmarshal raw buffer %v, err:%+v", string(buf), err.Error())
	}

	if rawResp.Code != 0 {
		return fmt.Errorf("ParseCommonPayload: errCode=%v, errMsg=%v", rawResp.Code, rawResp.Message)
	}

	if rawResp.Data == nil {
		// `Data` could be null, which means no data, not indicating error (user center case)
		// this leave input `v` unchanged
		return nil
	}

	if err := json.Unmarshal(*rawResp.Data, v); err != nil {
		return fmt.Errorf("ParseCommonPayload fail to unmarshal payload, resp = %+v", rawResp)
	}

	return nil
}

// DoHttp send http request and return: status_code and byte content
// if jsonBody is empty, input jsonBodyStr = ""
func DoHttp(client *http.Client, method string, url string, headers map[string]string, jsonBodyStr string) (*Response, error) {
	var body io.Reader = nil

	if len(jsonBodyStr) > 0 {
		body = bytes.NewBuffer([]byte(jsonBodyStr))
	}

	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		request.Header.Add(k, v)
	}

	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	output := &Response{
		StatusCode: resp.StatusCode,
		Buffer:     buf,
	}

	return output, nil
}

/*
DoHttpV send http request and parse whole response into value `V`
Useful to parse JsonResult
*/
func DoHttpData(client *http.Client, method string, url string, headers map[string]string, jsonBodyStr string, v interface{}) error {
	buRsp, err := DoHttp(client, method, url, headers, jsonBodyStr)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buRsp.Buffer, &v)
	if err != nil {
		return fmt.Errorf("DoHttpV fail to unmarshal data %v, err:%+v", string(buRsp.Buffer), err.Error())
	}
	return nil
}

/*
DoHttpEx send http request and parse payload (extract `RawResponse.Data`)
Useful to parse business object inside response, see `TestDoHttpEx`
*/
func DoHttpPayload(client *http.Client, method string, url string, headers map[string]string, jsonBodyStr string, output interface{}) error {
	buRsp, err := DoHttp(client, method, url, headers, jsonBodyStr)
	if err != nil {
		return err
	}
	return ParseCommonPayload(buRsp.Buffer, &output)
}

func ParamMapToReqBody(paramMap map[string]interface{}) (string, error) {
	if len(paramMap) > 0 {
		data, err := json.Marshal(paramMap)
		if err != nil {
			return "", err
		}
		return string(data), nil
	}
	return "", nil
}
