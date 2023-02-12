package cookies

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
)

type MenuSession struct {
	ID        string
	StartTime time.Time
	Checksum  string
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

func (m *MenuSession) RecalculateChecksum() error {
	checksum, err := utils.CreateHmacHexString(
		fmt.Sprintf("%s.%d", m.ID, m.StartTime.UnixNano()),
		[]byte("MenuSessionChecksum"),
	)

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
