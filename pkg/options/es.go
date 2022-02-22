package options

import (
	"github.com/cortezaproject/corteza-server/pkg/options"
	"strings"
)

type (
	EsOpt struct {
		Addresses            []string `env:"ES_ADDRESS"`
		EnableRetryOnTimeout bool     `env:"ES_ENABLE_RETRY_ON_TIMEOUT"`
		MaxRetries           int      `env:"ES_MAX_RETRIES"`
		IndexInterval        int      `env:"ES_INDEX_INTERVAL"`
	}
)

func ES() (o *EsOpt, err error) {
	o = &EsOpt{}
	return o, func() error {
		o.EnableRetryOnTimeout = options.EnvBool("ES_ENABLE_RETRY_ON_TIMEOUT", true)
		o.MaxRetries = options.EnvInt("ES_MAX_RETRIES", 5)
		o.IndexInterval = options.EnvInt("ES_INDEX_INTERVAL", 600)

		for _, a := range strings.Split(options.EnvString("ES_ADDRESS", "http://es:9200"), " ") {
			if a = strings.TrimSpace(a); a != "" {
				o.Addresses = append(o.Addresses, a)
			}
		}
		return nil
	}()
}
