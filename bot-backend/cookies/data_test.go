package cookies

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRecalculateChecksum(t *testing.T) {
	sess, err := NewMenuSession()
	assert.Nil(t, err)

	const expectedID = "9fbfaddb-1fb6-49c9-bad6-5f3a32d2b7cc"
	const expectedChecksum = "9f2b3ca66344e39a7e84cf41044950310bb2051ae406ba332db638ac38480de1"
	const expectedTimestamp int64 = 1676109328233770637

	sess.ID = expectedID
	sess.StartTime = time.Unix(0, expectedTimestamp)
	assert.EqualValues(t, expectedTimestamp, sess.StartTime.UnixNano())

	err = sess.RecalculateChecksum()
	assert.Nil(t, err)

	assert.EqualValues(t, expectedChecksum, sess.Checksum)
}

func TestMap(t *testing.T) {
	sess, err := NewMenuSession()
	assert.Nil(t, err)

	const expectedID = "9fbfaddb-1fb6-49c9-bad6-5f3a32d2b7cc"
	const expectedChecksum = "9f2b3ca66344e39a7e84cf41044950310bb2051ae406ba332db638ac38480de1"
	const expectedTimestamp int64 = 1676109328233770637

	sess.ID = expectedID
	sess.StartTime = time.Unix(0, expectedTimestamp)
	assert.EqualValues(t, expectedTimestamp, sess.StartTime.UnixNano())

	err = sess.RecalculateChecksum()
	assert.Nil(t, err)

	mapped := sess.Map()
	assert.EqualValues(t, expectedTimestamp, mapped["start_time"])
	assert.EqualValues(t, expectedChecksum, mapped["checksum"])
}

func TestIsExpired(t *testing.T) {
	sess, err := NewMenuSession()
	assert.Nil(t, err)

	sess.StartTime = sess.StartTime.Add(time.Second * time.Duration(-40))
	assert.True(t, sess.IsExpired(30))

	sess, err = NewMenuSession()
	assert.Nil(t, err)

	assert.False(t, sess.IsExpired(30))
}

func TestIsChecksumValid(t *testing.T) {
	sess, err := NewMenuSession()
	assert.Nil(t, err)
	const expectedID = "9fbfaddb-1fb6-49c9-bad6-5f3a32d2b7cc"
	const expectedTimestamp int64 = 1676109328233770637

	sess.ID = expectedID
	sess.StartTime = time.Unix(0, expectedTimestamp)
	err = sess.RecalculateChecksum()
	assert.Nil(t, err)

	assert.EqualValues(t, true, sess.IsChecksumValid())

	// Now change one attribute value, expected checksum to be invalid
	uuid, uuidErr := uuid.NewRandom()
	assert.Nil(t, uuidErr)

	sess.ID = uuid.String()
	assert.EqualValues(t, false, sess.IsChecksumValid())
}
