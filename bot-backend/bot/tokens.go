package bot

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/url"
	"sort"
	"strconv"
	"strings"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
)

type WebAppInitDataUser struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
	LangCode  string `json:"language_code"`
}

type WebAppInitData struct {
	QueryId  string
	User     *WebAppInitDataUser
	AuthDate int
	Hash     string
}

func UnMarshalWebAppInitData(data string) (*WebAppInitData, error) {
	result := WebAppInitData{}
	unescaped, err := url.QueryUnescape(data)
	if err != nil {
		return nil, err
	}

	arr := strings.Split(unescaped, "&")

	var userData WebAppInitDataUser
	for _, item := range arr {
		arr := strings.Split(item, "=")
		if strings.Index(item, "user=") == 0 {
			if err := json.Unmarshal([]byte(arr[1]), &userData); err != nil {
				return nil, err
			}
			result.User = &userData
		} else if strings.Index(item, "query_id=") == 0 {
			result.QueryId = arr[1]
		} else if strings.Index(item, "auth_date=") == 0 {
			val, err := strconv.Atoi(arr[1])
			if err != nil {
				return nil, err
			}
			result.AuthDate = val
		} else if strings.Index(item, "hash=") == 0 {
			result.Hash = arr[1]
		}
	}

	return &result, nil
}

func IsWebAppInitDataHashValid(data string) (bool, error) {
	unescaped, err := url.QueryUnescape(data)
	if err != nil {
		return false, err
	}

	hash := ""
	arr := strings.Split(unescaped, "&")
	arr2 := make([]string, 0)
	for _, item := range arr {
		if !strings.Contains(item, "hash=") {
			arr2 = append(arr2, item)
		} else {
			arr3 := strings.Split(item, "=")
			hash = arr3[1]
		}
	}

	sort.Strings(arr2)
	reordered := strings.Join(arr2, "\n")

	if secret, err := GetSecretKeyForWebApp(); err != nil {
		return false, err
	} else {
		h := hmac.New(sha256.New, secret)
		if _, err := h.Write([]byte(reordered)); err != nil {
			return false, err
		} else {
			sha := hex.EncodeToString(h.Sum(nil))
			return sha == hash, nil
		}
	}
}

func GetSecretKeyForWebApp() ([]byte, error) {
	token := utils.GetTelegramBotToken()
	secret := "WebAppData"
	h := hmac.New(sha256.New, []byte(secret))
	if _, err := h.Write([]byte(token)); err != nil {
		return nil, err
	} else {
		return h.Sum(nil), nil
	}
}
