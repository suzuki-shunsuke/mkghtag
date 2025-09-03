package github

import (
	"github.com/suzuki-shunsuke/ghtkn/pkg/api"
)

func NewTokenManager() *api.TokenManager {
	input := api.NewInput()
	tm := api.New(input)
	return tm
}
