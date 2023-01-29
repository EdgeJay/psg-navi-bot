package commands

import "encoding/json"

type InlineQueryData struct {
	QueryType string `json:"query_type"`
}

type MakeDropboxFileRequestInlineQueryEmbeddedData struct {
	RequestName string `json:"request_name"`
}

type MakeDropboxFileRequestInlineQueryData struct {
	QueryType string                                        `json:"query_type"`
	Data      MakeDropboxFileRequestInlineQueryEmbeddedData `json:"data"`
}

func GetMakeDropboxFileRequestInlineQueryData() string {
	data := MakeDropboxFileRequestInlineQueryData{
		QueryType: "make_file_request",
		Data:      MakeDropboxFileRequestInlineQueryEmbeddedData{""},
	}
	str, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(str)
}

func GetListDropboxFileRequestsInlineQueryData() string {
	data := InlineQueryData{
		QueryType: "list_file_requests",
	}
	str, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(str)
}

func GetDropboxFileRequestInfoInlineQueryData() string {
	data := InlineQueryData{
		QueryType: "file_request_info",
	}
	str, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(str)
}
