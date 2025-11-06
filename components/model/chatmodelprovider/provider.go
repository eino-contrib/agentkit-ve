package chatmodelprovider

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino-ext/components/model/arkbot"
	"github.com/cloudwego/eino-ext/components/model/claude"
	"github.com/cloudwego/eino-ext/components/model/deepseek"
	"github.com/cloudwego/eino-ext/components/model/gemini"
	"github.com/cloudwego/eino-ext/components/model/ollama"
	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino-ext/components/model/qwen"
	"github.com/cloudwego/eino/components"
	"github.com/cloudwego/eino/components/model"
	"google.golang.org/genai"
)

type Config struct {
	APIKey string

	BaseURL string

	// Model is the model name.
	Model string

	// MaxTokens is the max number of tokens, if reached the max tokens, the model will stop generating, and mostly return an finish reason of "length".
	MaxTokens *int
	// Temperature is the temperature, which controls the randomness of the model.
	Temperature *float32
	// TopP is the top p, which controls the diversity of the model.
	TopP *float32
	// Stop is the stop words, which controls the stopping condition of the model.
	Stop []string
}

type modelType string

const (
	openaiModelType      modelType = "OpenAI"
	azureOpenaiModelType modelType = "AzureOpenAI"
	geminiModelType      modelType = "Gemini"
	claudeModelType      modelType = "Claude"
	ollamaModelType      modelType = "Ollama"
	deepSeekModelType    modelType = "DeepSeek"
	arkModelType         modelType = "Ark"
	arkBotModelType      modelType = "ArkBot"
	qwenModelType        modelType = "Qwen"
	qianFaModelType      modelType = "QianFan"
)

var providerPrefixToModelType = map[string]modelType{
	"openai":     openaiModelType,
	"azure":      azureOpenaiModelType,
	"openrouter": openaiModelType,

	"vertex_ai": geminiModelType,
	"gemini":    geminiModelType,

	"anthropic":  claudeModelType,
	"ollama":     ollamaModelType,
	"deepseek":   deepSeekModelType,
	"volcengine": arkModelType,
	"dashscope":  qwenModelType,
}

type ChatModel struct {
	cfg *Config
	model.ToolCallingChatModel
}

func NewChatModel(ctx context.Context, provider string, cfg *Config) (*ChatModel, error) {

	var (
		err    error
		cModel model.ToolCallingChatModel
	)

	mType, ok := providerPrefixToModelType[provider]
	if !ok {
		return nil, fmt.Errorf("not support provider %s", provider)
	}

	switch mType {
	case arkModelType:
		arkCfg := cfg.toArkConfig()
		cModel, err = ark.NewChatModel(ctx, arkCfg)
		if err != nil {
			return nil, err
		}
	case arkBotModelType:
		arkBotCfg := cfg.toArkBotConfig()
		cModel, err = arkbot.NewChatModel(ctx, arkBotCfg)
		if err != nil {
			return nil, err
		}
	case deepSeekModelType:
		deepseekCfg := cfg.toDeepSeekConfig()
		cModel, err = deepseek.NewChatModel(ctx, deepseekCfg)
		if err != nil {
			return nil, err
		}
	case claudeModelType:
		claudeCfg := cfg.toClaudeConfig()
		cModel, err = claude.NewChatModel(ctx, claudeCfg)
		if err != nil {
			return nil, err
		}
	case geminiModelType:
		geminiCfg, err := cfg.toGeminiConfig(ctx)
		if err != nil {
			return nil, err
		}
		cModel, err = gemini.NewChatModel(ctx, geminiCfg)
		if err != nil {
			return nil, err
		}
	case ollamaModelType:
		ollamaCfg := cfg.toOllamaConfig()
		cModel, err = ollama.NewChatModel(ctx, ollamaCfg)
		if err != nil {
			return nil, err
		}
	case azureOpenaiModelType:
		openaiCfg := cfg.toOpenAIConfig()
		openaiCfg.ByAzure = true
		cModel, err = openai.NewChatModel(ctx, openaiCfg)
		if err != nil {
			return nil, err
		}
	case openaiModelType:
		openaiCfg := cfg.toOpenAIConfig()
		cModel, err = openai.NewChatModel(ctx, openaiCfg)
		if err != nil {
			return nil, err
		}
	case qwenModelType:
		qwenCfg := cfg.toQwenConfig()
		cModel, err = qwen.NewChatModel(ctx, qwenCfg)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("invalid model type: %s", mType)
	}
	return &ChatModel{
		cfg:                  cfg,
		ToolCallingChatModel: cModel,
	}, nil
}

func (c *ChatModel) GetType() string {
	typer, ok := c.ToolCallingChatModel.(components.Typer)
	if !ok {
		return "unknown"
	}
	return typer.GetType()
}

func (c *Config) toArkConfig() *ark.ChatModelConfig {

	cfg := &ark.ChatModelConfig{
		APIKey:  c.APIKey,
		Model:   c.Model,
		BaseURL: c.BaseURL,
	}
	if c.MaxTokens != nil {
		cfg.MaxTokens = c.MaxTokens
	}
	if c.Temperature != nil {
		cfg.Temperature = c.Temperature
	}
	if c.TopP != nil {
		cfg.TopP = c.TopP
	}
	if c.Stop != nil {
		cfg.Stop = c.Stop
	}

	return cfg
}

func (c *Config) toArkBotConfig() *arkbot.Config {

	cfg := &arkbot.Config{
		APIKey:  c.APIKey,
		Model:   c.Model,
		BaseURL: c.BaseURL,
	}

	if c.MaxTokens != nil {
		cfg.MaxTokens = c.MaxTokens
	}

	if c.Temperature != nil {
		cfg.Temperature = c.Temperature
	}

	if c.TopP != nil {
		cfg.TopP = c.TopP
	}

	if c.Stop != nil {
		cfg.Stop = c.Stop
	}

	return cfg
}

func (c *Config) toDeepSeekConfig() *deepseek.ChatModelConfig {

	cfg := &deepseek.ChatModelConfig{
		APIKey:  c.APIKey,
		Model:   c.Model,
		BaseURL: c.BaseURL,
	}

	if c.MaxTokens != nil {
		cfg.MaxTokens = *c.MaxTokens
	}
	if c.Temperature != nil {
		cfg.Temperature = *c.Temperature
	}

	if c.TopP != nil {
		cfg.TopP = *c.TopP
	}

	if c.Stop != nil {
		cfg.Stop = c.Stop
	}

	return cfg
}

func (c *Config) toClaudeConfig() *claude.Config {
	cfg := &claude.Config{
		APIKey: c.APIKey,
		Model:  c.Model,
	}
	if c.BaseURL != "" {
		cfg.BaseURL = &c.BaseURL
	}
	if c.MaxTokens != nil {
		cfg.MaxTokens = *c.MaxTokens
	}
	if c.Temperature != nil {
		cfg.Temperature = c.Temperature
	}
	if c.TopP != nil {
		cfg.TopP = c.TopP
	}
	if c.Stop != nil {
		cfg.StopSequences = c.Stop
	}

	return cfg
}

func (c *Config) toGeminiConfig(ctx context.Context) (*gemini.Config, error) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: c.APIKey,
		HTTPOptions: genai.HTTPOptions{
			BaseURL: c.BaseURL,
		},
	})
	if err != nil {
		return nil, err
	}

	cfg := &gemini.Config{
		Client: client,
		Model:  c.Model,
	}
	if c.MaxTokens != nil {
		cfg.MaxTokens = c.MaxTokens
	}
	if c.Temperature != nil {
		cfg.Temperature = c.Temperature
	}
	if c.TopP != nil {
		cfg.TopP = c.TopP
	}
	return cfg, nil

}

func (c *Config) toOllamaConfig() *ollama.ChatModelConfig {

	cfg := &ollama.ChatModelConfig{
		BaseURL: c.BaseURL,
		Model:   c.Model,
	}
	var options = &ollama.Options{}
	if c.MaxTokens != nil {
		options.NumPredict = *c.MaxTokens
	}
	if c.Temperature != nil {
		options.Temperature = *c.Temperature
	}
	if c.TopP != nil {
		options.TopP = *c.TopP
	}
	if c.Stop != nil {
		options.Stop = c.Stop
	}
	cfg.Options = options
	return cfg
}

func (c *Config) toOpenAIConfig() *openai.ChatModelConfig {
	cfg := &openai.ChatModelConfig{
		APIKey:  c.APIKey,
		Model:   c.Model,
		BaseURL: c.BaseURL,
	}

	if c.MaxTokens != nil {
		cfg.MaxCompletionTokens = c.MaxTokens
	}
	if c.Temperature != nil {
		cfg.Temperature = c.Temperature
	}
	if c.TopP != nil {
		cfg.TopP = c.TopP
	}
	if c.Stop != nil {
		cfg.Stop = c.Stop
	}
	return cfg
}

func (c *Config) toQwenConfig() *qwen.ChatModelConfig {
	cfg := &qwen.ChatModelConfig{
		APIKey:  c.APIKey,
		Model:   c.Model,
		BaseURL: c.BaseURL,
	}
	if c.MaxTokens != nil {
		cfg.MaxTokens = c.MaxTokens
	}
	if c.Temperature != nil {
		cfg.Temperature = c.Temperature
	}
	if c.TopP != nil {
		cfg.TopP = c.TopP
	}
	if c.Stop != nil {
		cfg.Stop = c.Stop
	}
	return cfg
}
