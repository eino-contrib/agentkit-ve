package chatmodelprovider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewChatModel(t *testing.T) {
	ctx := t.Context()
	cm, err := NewChatModel(ctx, "openai", &Config{})
	assert.Nil(t, err)
	assert.Equal(t, cm.GetType(), "OpenAI")

	cm, err = NewChatModel(ctx, "azure", &Config{})
	assert.Nil(t, err)
	assert.Equal(t, cm.GetType(), "OpenAI")

	cm, err = NewChatModel(ctx, "openrouter", &Config{})
	assert.Nil(t, err)
	assert.Equal(t, cm.GetType(), "OpenAI")

	cm, err = NewChatModel(ctx, "vertex_ai", &Config{APIKey: "api-key"})
	assert.Nil(t, err)
	assert.Equal(t, cm.GetType(), "Gemini")

	cm, err = NewChatModel(ctx, "gemini", &Config{APIKey: "api-key"})
	assert.Nil(t, err)
	assert.Equal(t, cm.GetType(), "Gemini")

	cm, err = NewChatModel(ctx, "anthropic", &Config{APIKey: "api-key"})
	assert.Nil(t, err)
	assert.Equal(t, cm.GetType(), "Claude")

	cm, err = NewChatModel(ctx, "ollama", &Config{APIKey: "api-key"})
	assert.Nil(t, err)
	assert.Equal(t, cm.GetType(), "Ollama")

	cm, err = NewChatModel(ctx, "deepseek", &Config{APIKey: "api-key", Model: "deepseek-r1"})
	assert.Nil(t, err)
	assert.Equal(t, cm.GetType(), "DeepSeek")

	cm, err = NewChatModel(ctx, "volcengine", &Config{APIKey: "api-key"})
	assert.Nil(t, err)
	assert.Equal(t, cm.GetType(), "Ark")

	cm, err = NewChatModel(ctx, "dashscope", &Config{APIKey: "api-key"})
	assert.Nil(t, err)
	assert.Equal(t, cm.GetType(), "Qwen")

}
