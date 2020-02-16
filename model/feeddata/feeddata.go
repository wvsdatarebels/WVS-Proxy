package feeddata

import "encoding/json"

func UnmarshalFeedData(data []byte) (FeedData, error) {
	var r FeedData
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FeedData) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type FeedData struct {
	Success bool   `json:"success,omitempty"`
	FeedURL string `json:"feedURL,omitempty"`
	Data    []Data `json:"data"`
}

type Data struct {
	Impediments []Impediment `json:"impediments"`
	Snippets    []string     `json:"snippets"`
	Author      string      `json:"author,omitempty"`
	Link        string      `json:"link,omitempty"`
	PubDate     string      `json:"pubDate,omitempty"`
}

type Impediment struct {
	Happening string `json:"happening,omitempty"`
	Lesson    string `json:"lesson"`
	Room      string `json:"room"`
	Time      string `json:"time,omitempty"`
}

