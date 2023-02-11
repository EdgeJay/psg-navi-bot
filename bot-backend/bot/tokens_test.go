package bot

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func init() {
	if cfg, err := os.ReadFile("../.env"); err != nil {
		panic(err)
	} else {
		viper.SetConfigType("env")
		viper.ReadConfig(bytes.NewBuffer(cfg))
		fmt.Println("Env vars loaded")
	}
}

func TestUnMarshalWebAppInitData(t *testing.T) {
	data := `query_id=AAHFIPczAAAAAMUg9zNjREeS&user={"id":871833797,"first_name":"Huijie","last_name":"Wu","username":"hjwusg","language_code":"en"}&auth_date=1675788080&hash=455da5699cea26cd9eae265c951f542a8e7d63ab28cec6d5821428a50c7239bf`
	result, err := UnMarshalWebAppInitData(data)
	if err != nil {
		assert.Fail(t, "Expected error to be nil", err)
	}
	assert.EqualValues(t, "AAHFIPczAAAAAMUg9zNjREeS", result.QueryId)
	assert.EqualValues(t, 1675788080, result.AuthDate)
	assert.EqualValues(t, "455da5699cea26cd9eae265c951f542a8e7d63ab28cec6d5821428a50c7239bf", result.Hash)
	assert.EqualValues(t, "hjwusg", result.User.UserName)
	assert.EqualValues(t, 871833797, result.User.Id)
}

func TestUnMarshalWebAppInitData_WithUrlEscapedString(t *testing.T) {
	data := `query_id=AAHFIPczAAAAAMUg9zMK1xx8&user=%7B%22id%22%3A871833797%2C%22first_name%22%3A%22Huijie%22%2C%22last_name%22%3A%22Wu%22%2C%22username%22%3A%22hjwusg%22%2C%22language_code%22%3A%22en%22%7D&auth_date=1675956839&hash=c4fae47d0712a7768dc3cf957b08fec927d167143cebc1f1c00fb159add92c2a`
	result, err := UnMarshalWebAppInitData(data)
	if err != nil {
		assert.Fail(t, "Expected error to be nil", err)
	}
	assert.EqualValues(t, "AAHFIPczAAAAAMUg9zMK1xx8", result.QueryId)
	assert.EqualValues(t, 1675956839, result.AuthDate)
	assert.EqualValues(t, "c4fae47d0712a7768dc3cf957b08fec927d167143cebc1f1c00fb159add92c2a", result.Hash)
	assert.EqualValues(t, "hjwusg", result.User.UserName)
	assert.EqualValues(t, 871833797, result.User.Id)
}

func TestIsWebAppInitDataHashValid(t *testing.T) {
	data := `query_id=AAHFIPczAAAAAMUg9zMK1xx8&user=%7B%22id%22%3A871833797%2C%22first_name%22%3A%22Huijie%22%2C%22last_name%22%3A%22Wu%22%2C%22username%22%3A%22hjwusg%22%2C%22language_code%22%3A%22en%22%7D&auth_date=1675956839&hash=c4fae47d0712a7768dc3cf957b08fec927d167143cebc1f1c00fb159add92c2a`
	result, err := IsWebAppInitDataHashValid(data)
	if err != nil {
		assert.Fail(t, "Expected error to be nil", err)
	}
	assert.EqualValues(t, true, result)
}
