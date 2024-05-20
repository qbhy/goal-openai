package goal_openai

import (
	"github.com/goal-web/contracts"
	"sync"
)

type ServiceProvider struct {
}

func (service ServiceProvider) Register(application contracts.Application) {
	application.Singleton("openai", func(config contracts.Config) Factory {
		conf := config.Get("openai").(*Config)
		return &factory{
			config:  conf,
			clients: sync.Map{},
		}
	})
}

func (service ServiceProvider) Start() error {
	return nil
}

func (service ServiceProvider) Stop() {
}
