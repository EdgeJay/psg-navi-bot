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

func NewMenuSession(startTime time.Time) (*MenuSession, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	id := uuid.String()
	checksum, csErr := utils.CreateHmacHexString(
		fmt.Sprintf("%s.%d", id, now.UnixNano()),
		[]byte("MenuSessionChecksum"),
	)

	if csErr != nil {
		return nil, csErr
	}

	session := MenuSession{
		ID:        id,
		StartTime: now,
		Checksum:  checksum,
	}

	return &session, nil
}

func (m *MenuSession) Map() gin.H {
	return gin.H{
		"id":         m.ID,
		"start_time": m.StartTime.UnixNano(),
		"checksum":   m.Checksum,
	}
}
