package ark

import "github.com/volcengine/volcengine-go-sdk/volcengine/request"

type ListApiKeysInput struct {
	ProjectName string `json:"ProjectName"`
}

func (a *ArkService) ListApiKeys(input *ListApiKeysInput) (map[string]any, error) {
	if input == nil {
		input = &ListApiKeysInput{}
	}
	output := map[string]any{}

	req := a.ARK.NewRequest(&request.Operation{
		Name:       "ListApiKeys",
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}, input, &output)

	err := req.Send()
	if err != nil {
		return nil, err
	}

	return output, nil
}
