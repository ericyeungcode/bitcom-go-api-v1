package bitws

const (
	SubTypeSubscribe = "subscribe"
)

type SubReq struct {
	Type        string   `json:"type"`
	Instruments []string `json:"instruments"`
	Channels    []string `json:"channels"`
	Interval    string   `json:"interval"`
}

type PrivateSubSeq struct {
	*SubReq
	Token string `json:"token"`
}
