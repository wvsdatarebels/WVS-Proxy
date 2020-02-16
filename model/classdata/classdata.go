package classdata

import "encoding/json"

func UnmarshalWVSClass(data []byte) (WVSClass, error) {
	var r WVSClass
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WVSClass) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type WVSClass struct {
	Success      bool    `json:"success"`
	ResponseTime int64   `json:"responseTime"`
	Data         []Data  `json:"data"`
}

type Data struct {
	Class string `json:"class"`
	URL   string `json:"url"`
}
