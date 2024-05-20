package goal_openai

import "github.com/sashabaranov/go-openai"

type Factory interface {
	Client(drivers ...string) *openai.Client
}
