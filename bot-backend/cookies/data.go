package cookies

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
)

const ChecksumSecret = "MenuSessionChecksum"

type MenuSession struct {
	ID        string
	StartTime time.Time
	Checksum  string
}

type MenuSessionJson struct {
	ID        string `json:"id"`
	StartTime int64  `json:"start_time"`
	Checksum  string `json:"checksum"`
}

func NewMenuSession() (*MenuSession, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	id := uuid.String()

	session := MenuSession{
		ID:        id,
		StartTime: now,
	}

	csErr := session.RecalculateChecksum()
	if csErr != nil {
		return nil, csErr
	}

	return &session, nil
}

func (m *MenuSession) createChecksum() (string, error) {
	return utils.CreateHmacHexString(
		fmt.Sprintf("%s.%d", m.ID, m.StartTime.UnixNano()),
		[]byte(ChecksumSecret),
	)
}

func (m *MenuSession) RecalculateChecksum() error {
	checksum, err := m.createChecksum()

	if err != nil {
		return err
	}

	m.Checksum = checksum

	return nil
}

func (m *MenuSession) Map() gin.H {
	return gin.H{
		"id":         m.ID,
		"start_time": m.StartTime.UnixNano(),
		"checksum":   m.Checksum,
	}
}

func (m *MenuSession) IsExpired(duration int) bool {
	now := time.Now()
	limit := m.StartTime.Add(time.Second * time.Duration(duration))
	return now.After(limit)
}

func (m *MenuSession) IsChecksumValid() bool {
	checksum, err := m.createChecksum()

	if err != nil {
		return false
	}

	return checksum == m.Checksum
}

func (m *MenuSession) ParseJson(str string) error {
	var menuSessionJson MenuSessionJson
	err := json.Unmarshal([]byte(str), &menuSessionJson)
	if err != nil {
		return err
	}

	m.ID = menuSessionJson.ID
	m.StartTime = time.Unix(0, menuSessionJson.StartTime)
	m.Checksum = menuSessionJson.Checksum

	return nil
}
