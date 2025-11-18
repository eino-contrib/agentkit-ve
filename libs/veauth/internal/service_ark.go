package internal

import (
	"github.com/volcengine/volcengine-go-sdk/service/ark"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/client"
)

func NewArkService(p client.ConfigProvider, cfgs ...*volcengine.Config) *ArkService {
	return &ArkService{
		ARK: ark.New(p, cfgs...),
	}
}

type ArkService struct {
	*ark.ARK
}
