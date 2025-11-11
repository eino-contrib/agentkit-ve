package ark

import "github.com/volcengine/volcengine-go-sdk/volcengine/request"

type GetRawApiKeyInput struct {
	Id any `json:"Id"`
}

func (a *ArkService) GetRawApiKey(input *GetRawApiKeyInput) (map[string]any, error) {
	if input == nil {
		input = &GetRawApiKeyInput{}
	}
	output := map[string]any{}

	req := a.ARK.NewRequest(&request.Operation{
		Name:       "GetRawApiKey",
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}, input, &output)

	err := req.Send()
	if err != nil {
		return nil, err
	}
	return output, nil
}
