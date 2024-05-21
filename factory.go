package goal_openai

import (
	"github.com/sashabaranov/go-openai"
	"sync"
)

type factory struct {
	config *Config

	clients sync.Map
}

func (f *factory) Client(drivers ...string) *openai.Client {
	var key = f.config.Default
	if len(drivers) > 0 {
		key = drivers[0]
	}

	if client, exists := f.clients.Load(key); exists {
		return client.(*openai.Client)
	}

	client := openai.NewClientWithConfig(f.config.Configurations[key])

	f.clients.Store(key, client)
	return client
}
