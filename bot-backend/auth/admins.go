package auth

import (
	"encoding/base64"
	"encoding/json"
	"log"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/aws"
	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
)

var adminManager *AdminManager

func NewAdminManagerFromB64Json(data string) (*AdminManager, error) {
	// decode from B64 string first
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return NewAdminManagerFromJson(decoded)
}

func NewAdminManagerFromJson(data []byte) (*AdminManager, error) {
	var am AdminManager
	if err := json.Unmarshal(data, &am); err != nil {
		return nil, err
	}
	return &am, nil
}

func NewAdminManager() (*AdminManager, error) {
	if adminManager != nil {
		return adminManager, nil
	}

	cfgBase64 := aws.GetStringParameter(
		utils.GetAWSParamStoreKeyName("config_admins"),
		"",
	)

	am, err := NewAdminManagerFromB64Json(cfgBase64)
	if err != nil {
		return nil, err
	}

	log.Printf("init NewAdminManager with %d admin(s)\n", len(am.Admins))

	adminManager = am

	return adminManager, nil
}

func (am *AdminManager) CanPerformTask(username, domain, task string) bool {
	var admin Admin
	for _, v := range am.Admins {
		if v.UserName == username {
			admin = v
		}
	}

	if admin.UserName == "" {
		return false
	}

	return admin.CanPerformTask(domain, task)
}

func (a *Admin) FilterScope(domain string) []string {
	filtered := []string{}

	for _, scope := range a.Scope {
		if scope.Domain == domain {
			filtered = append(filtered, scope.Task)
		}
	}

	return filtered
}

func (a *Admin) CanPerformTask(domain, task string) bool {
	filtered := a.FilterScope(domain)
	for _, t := range filtered {
		if t == AllTasksPermission {
			return true
		}
		if t == task {
			return true
		}
	}
	return false
}
