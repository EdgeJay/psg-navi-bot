package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
)

type FileRequest struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Url   string `json:"url"`
}

type ListFileRequestsResponse struct {
	FileRequests []FileRequest `json:"data"`
}

func DropboxListFileRequests(c *gin.Context) {
	// trigger add dropbox file request payload
	// Get Dropbox client
	dbx := utils.NewDropboxClient(
		utils.GetDropboxAppKey(),
		utils.GetDropboxAppSecret(),
		utils.GetDropboxRefreshToken(),
	)

	response := ListFileRequestsResponse{
		FileRequests: make([]FileRequest, 0),
	}

	if allFileRequests, err := dbx.GetFileRequests(); err == nil {
		for _, fileRequest := range *allFileRequests {
			if fileRequest.IsOpen {
				fileReq := FileRequest{
					ID:    fileRequest.ID,
					Title: fileRequest.Title,
					Url:   fileRequest.URL,
				}
				response.FileRequests = append(response.FileRequests, fileReq)
			}
		}
	}

	// response
	data, err := json.Marshal(response)
	if err != nil {
		c.Abort()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to parse file requests",
		})
		return
	}

	c.Data(http.StatusOK, "application/json", data)
}
