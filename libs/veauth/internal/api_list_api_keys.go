package internal

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
)

type ListApiKeysInput struct {
	ProjectName string `json:"ProjectName"`
}

type ListApiKeysOutput struct {
	Metadata *response.ResponseMetadata `json:"ResponseMetadata"`
	Items    []struct {
		ID int64 `json:"Id"`
	} `json:"Items"`
}

func (a *ArkService) ListApiKeys(input *ListApiKeysInput) (*ListApiKeysOutput, error) {
	if input == nil {
		input = &ListApiKeysInput{}
	}

	output := new(ListApiKeysOutput)
	req := a.ARK.NewRequest(&request.Operation{
		Name:       "ListApiKeys",
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}, input, output)

	err := req.Send()
	if err != nil {
		return nil, err
	}

	return output, nil
}
