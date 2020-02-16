package dsbnextdata

import "encoding/json"

func UnmarshalDSBNextData(data []byte) (DSBNextData, error) {
	var r DSBNextData
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DSBNextData) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DSBNextData struct {
	ResponseTime int64    `json:"responseTime"`
	Result       []Result `json:"result"`
}

type Result struct {
	Date string  `json:"date"`
	Data []Datum `json:"data"`
}

type Datum struct {
	Date              string `json:"date"`
	Time              string `json:"time"`
	Day               string `json:"day"`
	SchoolClassBefore string `json:"school_class_before"`
	LessonBefore      string `json:"lesson_before"`
	RoomBefore        string `json:"room_before"`
	Type              string `json:"type"`
	Representative    string `json:"representative"`
	LessonAfter       string `json:"lesson_after"`
	RoomAfter         string `json:"room_after"`
	Text              string `json:"text"`
	Cancelled         bool   `json:"cancelled"`
}
