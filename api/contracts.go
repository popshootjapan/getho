package api

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/popshootjapan/getho-cli/config"
	"github.com/popshootjapan/getho-cli/utils"
)

type ContractsResp struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	ABI     string `json:abi`
}

func GetContracts(token, subdomain string) ([]ContractsResp, error) {
	getContractsEndPoint := fmt.Sprintf(config.ContractsEndpoint, subdomain)
	headers := map[string]string{}
	headers["Authorization"] = token

	resp, err := utils.HttpGetWithHeaders(config.APIHost+getContractsEndPoint, headers)
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

	var data []ContractsResp
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func PostContracts(token, subdomain, name, address, abi string) (string, error) {
	jsonStr := fmt.Sprintf(`{"name": "%s", "address": "%s", "abi": "%s"}`, name, address, strings.Replace(abi, "\"", "\\\"", -1))
	deployEndPoint := fmt.Sprintf(config.ContractsEndpoint, subdomain)
	headers := map[string]string{}
	headers["Authorization"] = token

	resp, err := utils.HttpPostWithHeaders(config.APIHost+deployEndPoint, jsonStr, headers)
	if err != nil {
		return "", err
	}

	return resp.Status, nil
}
