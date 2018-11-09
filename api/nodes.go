package api

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/popshootjapan/getho-cli/config"
	"github.com/popshootjapan/getho-cli/utils"
)

type NodeResp struct {
	Subdomain string `json:"subdomain"`
}

func GetNodes(token string) ([]NodeResp, error) {
	headers := map[string]string{}
	headers["Authorization"] = token

	resp, err := utils.HttpGetWithHeaders(config.APIHost+config.NodesEndpoint, headers)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 {
		return nil, errors.Errorf("%s", resp.Status)
	}

	bytes, err := utils.Resp2Bytes(*resp)
	if err != nil {
		return nil, err
	}

	var data []NodeResp
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
