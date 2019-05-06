package config

import "go.zenithar.org/pkg/platform"

// Configuration contains kornflake settings
type Configuration struct {
	Debug struct {
		Enable bool `toml:"enable" default:"false" comment:"allow debug mode"`
	} `toml:"Debug" comment:"###############################\n Debug \n##############################"`

	Instrumentation platform.InstrumentationConfig `toml:"Instrumentation" comment:"###############################\n Instrumentation \n##############################"`

	Server struct {
		Network string `toml:"network" default:"tcp" comment:"Network class used for listen (tcp, tcp4, tcp6, unixsocket)"`
		Listen  string `toml:"listen" default:":5555" comment:"Listen address for gRPC server"`
		UseTLS  bool   `toml:"useTLS" default:"false" comment:"Enable TLS listener"`
		TLS     struct {
			CertificatePath              string `toml:"certificatePath" default:"" comment:"Certificate path"`
			PrivateKeyPath               string `toml:"privateKeyPath" default:"" comment:"Private Key path"`
			CACertificatePath            string `toml:"caCertificatePath" default:"" comment:"CA Certificate Path"`
			ClientAuthenticationRequired bool   `toml:"clientAuthenticationRequired" default:"false" comment:"Force client authentication"`
		} `toml:"TLS" comment:"TLS Socket settings"`
	}

	WorkerID uint64 `toml:"worker_id" default:"0" comment:"Defines id generator worker ID must be different per instance."`
}
