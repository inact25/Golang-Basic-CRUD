package Rest

type Res struct {
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}
