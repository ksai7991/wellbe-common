package apiclient

type CommonResponse struct {
	ResultCode int    `json:"result_code"`
	Message    string `json:"message"`
}