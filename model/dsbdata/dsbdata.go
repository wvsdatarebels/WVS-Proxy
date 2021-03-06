package dsbdata

import "encoding/json"

func UnmarshalDSBData(data []byte) (DSBData, error) {
	var r DSBData
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DSBData) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DSBData struct {
	ResponseTime int64    `json:"responseTime"`
	Result       []Result `json:"result"`
}

type Result struct {
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
