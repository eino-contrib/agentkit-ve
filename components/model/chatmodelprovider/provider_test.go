package chatmodelprovider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewChatModel(t *testing.T) {
	ctx := t.Context()
	cm, err := NewChatModel(ctx, &Config{Provider: "openai"})
	assert.Nil(t, err)
	assert.Equal(t, cm.GetType(), "OpenAI")

	cm, err = NewChatModel(ctx, &Config{Provider: "azure"})
	assert.Nil(t, err)
	assert.Equal(t, cm.GetType(), "OpenAI")

	cm, err = NewChatModel(ctx, &Config{Provider: "openrouter"})
	assert.Nil(t, err)
	assert.Equal(t, cm.GetType(), "OpenAI")

	cm, err = NewChatModel(ctx, &Config{Provider: "vertex_ai", APIKey: "api-key"})
	assert.Nil(t, err)
	assert.Equal(t, cm.GetType(), "Gemini")

	cm, err = NewChatModel(ctx, &Config{Provider: "gemini", APIKey: "api-key"})
	assert.Nil(t, err)
	assert.Equal(t, cm.GetType(), "Gemini")

	cm, err = NewChatModel(ctx, &Config{Provider: "anthropic", APIKey: "api-key"})
	assert.Nil(t, err)
	assert.Equal(t, cm.GetType(), "Claude")

	cm, err = NewChatModel(ctx, &Config{Provider: "ollama", APIKey: "api-key"})
	assert.Nil(t, err)
	assert.Equal(t, cm.GetType(), "Ollama")

	cm, err = NewChatModel(ctx, &Config{Provider: "deepseek", APIKey: "api-key", Model: "deepseek-r1"})
	assert.Nil(t, err)
	assert.Equal(t, cm.GetType(), "DeepSeek")

	cm, err = NewChatModel(ctx, &Config{Provider: "volcengine", APIKey: "api-key"})
	assert.Nil(t, err)
	assert.Equal(t, cm.GetType(), "Ark")

	cm, err = NewChatModel(ctx, &Config{Provider: "dashscope", APIKey: "api-key"})
	assert.Nil(t, err)
	assert.Equal(t, cm.GetType(), "Qwen")

}
