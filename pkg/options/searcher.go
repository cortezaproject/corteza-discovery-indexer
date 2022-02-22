package options

import (
	"fmt"
	"github.com/cortezaproject/corteza-server/pkg/options"
	"os"
)

type (
	SearcherOpt struct {
		Enabled bool
		//HttpAddr     string
		JwtSecret    []byte
		ClientKey    string
		ClientSecret string
	}
)

const (
	discoverySearcher     = "DISCOVERY_SEARCHER_"
	searcherEnvKeyEnabled = discoverySearcher + "ENABLED"
	//searcherEnvKeyHttpAddr = discoverySearcher + "HTTP_ADDR"
	envKeyJwtSecret    = discoverySearcher + "CORTEZA_SERVER_JWT_SECRET"
	envKeyClientKey    = discoverySearcher + "CORTEZA_SERVER_CLIENT_KEY"
	envKeyClientSecret = discoverySearcher + "CORTEZA_SERVER_CLIENT_SECRET"
)

func Searcher() (*SearcherOpt, error) {
	o := &SearcherOpt{}

	return o, func() error {
		o.Enabled = options.EnvBool(searcherEnvKeyEnabled, true)

		//o.CortezaServerBaseUrl = options.EnvString(envKeyBaseUrl, "http://server:80")
		//if o.CortezaServerBaseUrl == "" {
		//	return fmt.Errorf("endpoint URL for corteza (%s) is empty or missing", envKeyAuthUrl)
		//}

		//o.HttpAddr = options.EnvString(searcherEnvKeyHttpAddr, "0.0.0.0:80")

		//o.CortezaServerAuthUrl = options.EnvString(envKeyAuthUrl, o.CortezaServerBaseUrl+"/auth")
		//if o.CortezaServerAuthUrl == "" {
		//	return fmt.Errorf("endpoint URL for corteza auth (%s) is empty or missing", envKeyAuthUrl)
		//}

		if tmp := os.Getenv(envKeyJwtSecret); tmp != "" {
			o.JwtSecret = []byte(tmp)
		}

		if o.ClientKey = os.Getenv(envKeyClientKey); o.ClientKey == "" {
			return fmt.Errorf("client key (%s) is empty or missing", envKeyClientKey)
		}

		if o.ClientSecret = os.Getenv(envKeyClientSecret); o.ClientSecret == "" {
			return fmt.Errorf("client secret (%s) is empty or missing", envKeyClientSecret)
		}

		//for _, a := range strings.Split(options.EnvString(envKeyEsAddr, "http://localhost:9200"), " ") {
		//	if a = strings.TrimSpace(a); a != "" {
		//		o.Es.Addresses = append(o.Es.Addresses, a)
		//	}
		//}

		//o.EnableRetryOnTimeout = options.EnvBool(searcherEnvKeyEsEnableRetryOnTimeout, true)
		//o.MaxRetries = options.EnvInt(searcherEnvKeyEsMaxRetries, 5)

		return nil
	}()
}
