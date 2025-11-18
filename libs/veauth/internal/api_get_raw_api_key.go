package internal

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
)

type GetRawApiKeyInput struct {
	Id string `queryName:"Id"`
}
type GetRawApiKeyOutput struct {
	Metadata *response.ResponseMetadata `json:"ResponseMetadata"`
	ApiKey   *string                    `json:"ApiKey"`
}

func (a *ArkService) GetRawApiKey(input *GetRawApiKeyInput) (*GetRawApiKeyOutput, error) {
	if input == nil {
		input = &GetRawApiKeyInput{}
	}
	output := new(GetRawApiKeyOutput)

	req := a.ARK.NewRequest(&request.Operation{
		Name:       "GetRawApiKey",
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}, input, output)

	err := req.Send()
	if err != nil {
		return nil, err
	}
	return output, nil
}
