package goal_openai

import "github.com/sashabaranov/go-openai"

type Config struct {
	Default        string
	Configurations map[string]openai.ClientConfig
}
