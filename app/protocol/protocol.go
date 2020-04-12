package protocol

type Resp struct {
	Code  int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type CallbackMessage struct {
	Msg string `json:"msg"`
}
