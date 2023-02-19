package auth

import (
	"encoding/base64"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAdminManagerFromJson(t *testing.T) {
	wd, wdErr := os.Getwd()
	assert.Nil(t, wdErr)

	data, err := os.ReadFile(wd + "/../../config/admins.json")
	assert.Nil(t, err)

	adminManager, amErr := NewAdminManagerFromJson(data)
	assert.Nil(t, amErr)

	assert.EqualValues(t, 1, len(adminManager.Admins))
	assert.EqualValues(t, "hjwusg", adminManager.Admins[0].UserName)
	assert.EqualValues(t, "dropbox", adminManager.Admins[0].Scope[0].Domain)
	assert.EqualValues(t, "all", adminManager.Admins[0].Scope[0].Task)
}

func TestNewAdminManagerFromB64Json(t *testing.T) {
	wd, wdErr := os.Getwd()
	assert.Nil(t, wdErr)

	data, err := os.ReadFile(wd + "/../../config/admins.json")
	assert.Nil(t, err)

	// encode json as base64
	b64 := base64.StdEncoding.EncodeToString(data)

	adminManager, amErr := NewAdminManagerFromB64Json(b64)
	assert.Nil(t, amErr)

	assert.EqualValues(t, 1, len(adminManager.Admins))
	assert.EqualValues(t, "hjwusg", adminManager.Admins[0].UserName)
	assert.EqualValues(t, "dropbox", adminManager.Admins[0].Scope[0].Domain)
	assert.EqualValues(t, "all", adminManager.Admins[0].Scope[0].Task)
}
