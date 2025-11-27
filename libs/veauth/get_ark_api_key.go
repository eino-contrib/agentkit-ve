package veauth

import (
	"errors"
	"strconv"

	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"

	"github.com/eino-contrib/agentkit-ve/libs/veauth/internal"
)

type option struct {
	region       string
	sessionToken string
}
type OptionFn func(*option)

func WithRegion(region string) OptionFn {
	return func(o *option) {
		o.region = region
	}
}

func WithSessionToken(sessionToken string) OptionFn {
	return func(o *option) {
		o.sessionToken = sessionToken
	}
}

func GetArkAPIKey(ak, sk string, opts ...OptionFn) (string, error) {
	opt := &option{
		region: "cn-beijing",
	}

	for _, o := range opts {
		o(opt)
	}

	config := volcengine.NewConfig().
		WithRegion(opt.region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, opt.sessionToken))
	sess, err := session.NewSession(config)
	if err != nil {
		return "", err
	}

	arkSrv := internal.NewArkService(sess)
	response, err := arkSrv.ListApiKeys(&internal.ListApiKeysInput{ProjectName: "default"})
	if err != nil {
		return "", err
	}

	if len(response.Items) == 0 {
		return "", errors.New("list api keys returned empty list")
	}

	rawApiKeyOutput, err := arkSrv.GetRawApiKey(&internal.GetRawApiKeyInput{
		Id: strconv.Itoa(int(response.Items[0].ID)),
	})
	if err != nil {
		return "", err
	}

	apiKey := rawApiKeyOutput.ApiKey
	if apiKey == nil {
		return "", errors.New("get raw api key returned nil")
	}

	return *apiKey, nil

}
