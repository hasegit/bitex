package infrastructure

type JsonRequest struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  Params `json:"params"`
	Id      int    `json:"id"`
}

type JsonResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  bool   `json:"result"`
}

type Params struct {
	Channel string `json:"channel"`
}

func NewJsonRequest(channel string, method string) JsonRequest {

	params := Params{
		Channel: channel,
	}

	return JsonRequest{
		Jsonrpc: "2.0",
		Method:  method,
		Params:  params,
		Id:      1,
	}
}
