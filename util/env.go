package util

import (
	env "github.com/Netflix/go-env"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

// LoadFromEnv func
func LoadFromEnv(v interface{}) {
	env.UnmarshalFromEnviron(v)
}
