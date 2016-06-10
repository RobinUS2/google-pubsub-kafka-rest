package main

// Json structs

type PostTopic struct {
	Records []PostRecord `json:"records"`
}

type PostRecord struct {
	Value string `json:"value"`
}
