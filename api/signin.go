package api

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"github.com/popshootjapan/getho-cli/config"
	"github.com/popshootjapan/getho-cli/utils"
)

type SigninResp struct {
	Token string `json:"token"`
}

func Singin(email, password string) (*SigninResp, error) {
	jsonStr := fmt.Sprintf(`{"email": "%s", "password": "%s"}`, email, password)
	resp, err := utils.HttpPost(config.APIHost+config.SinginEndpoint, jsonStr)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 {
		return nil, errors.Errorf("%s", resp.Status)
	}

	data := new(SigninResp)
	bytes, err := utils.Resp2Bytes(*resp)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
