package model

type QywxText struct {
	MsgType string   `json:"msgtype"`
	Text    *Content `json:"text"`
}

type QywxMarkdown struct {
	MsgType  string   `json:"msgtype"`
	Markdown *Content `json:"markdown"`
}

type Content struct {
	Content string `json:"content"`
}
