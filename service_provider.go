package goal_openai

import (
	"github.com/goal-web/contracts"
	"github.com/sashabaranov/go-openai"
	"sync"
)

type ServiceProvider struct {
}

func Service() contracts.ServiceProvider {
	return ServiceProvider{}
}

func (service ServiceProvider) Register(application contracts.Application) {
	application.Singleton("openai", func(config contracts.Config) Factory {
		conf := config.Get("openai").(*Config)
		return &factory{
			config:  conf,
			clients: sync.Map{},
		}
	})

	application.Singleton("openai.client", func(f Factory) *openai.Client {
		return f.Client()
	})
}

func (service ServiceProvider) Start() error {
	return nil
}

func (service ServiceProvider) Stop() {
}
