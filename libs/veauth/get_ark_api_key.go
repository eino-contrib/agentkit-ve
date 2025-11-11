package veauth

import (
	"errors"

	"github.com/bytedance/sonic"
	"github.com/eino-contrib/agentkit-ve/libs/veauth/ark"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

type ListApiKeysOutput struct {
	Result struct {
		Items []struct {
			ID int64 `json:"Id"`
		} `json:"Items"`
	} `json:"Result"`
}

func GetArkAPIKey(ak, sk string, token string) (string, error) {
	config := volcengine.NewConfig().
		WithRegion("cn-beijing").
		WithCredentials(credentials.NewStaticCredentials(ak, sk, token))
	sess, err := session.NewSession(config)
	if err != nil {
		return "", err
	}
	arkSrv := ark.NewArkService(sess)

	response, err := arkSrv.ListApiKeys(&ark.ListApiKeysInput{ProjectName: "default"})
	if err != nil {
		return "", err
	}
	output := &ListApiKeysOutput{}
	bs, _ := sonic.Marshal(response)
	err = sonic.Unmarshal(bs, output)
	if err != nil {
		return "", err
	}

	if len(output.Result.Items) == 0 {
		return "", errors.New("list api keys returned empty list")
	}
	firstApiKeyId := output.Result.Items[0].ID

	response, err = arkSrv.GetRawApiKey(&ark.GetRawApiKeyInput{
		Id: firstApiKeyId,
	})
	if err != nil {
		return "", err
	}
	result, ok := response["Result"].(map[string]any)
	if !ok {
		return "", errors.New("GetRawApiKey did not return Result field")
	}
	apiKey, ok := result["ApiKey"].(string)
	if !ok {
		return "", errors.New("GetRawApiKey did not return valid ApiKey")
	}
	return apiKey, nil

}
