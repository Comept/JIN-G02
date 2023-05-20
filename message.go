package sth

type MessagePost struct {
	Sub     string `json: "sub"`
	Message string `json: "message"`
}

type Message struct {
	Id      int    `json: "-"`
	Pub     int    `json: "pub"`
	Sub     int    `json: "sub"`
	Message string `json: "message"`
}