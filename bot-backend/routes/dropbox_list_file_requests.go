package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
)

type FileRequest struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Desc      string    `json:"desc"`
	CreatedOn time.Time `json:"created_on"`
	Url       string    `json:"url"`
	FileCount int       `json:"file_count"`
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
					ID:        fileRequest.ID,
					Title:     fileRequest.Title,
					Desc:      fileRequest.Description,
					CreatedOn: fileRequest.Created,
					Url:       fileRequest.URL,
					FileCount: fileRequest.FileCount,
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
