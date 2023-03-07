module github.com/eyedeekay/mtg-i2p

go 1.17

require (
	github.com/9seconds/mtg/v2 v2.1.7
	github.com/alecthomas/kong v0.7.1
	github.com/alecthomas/units v0.0.0-20211218093645-b94a6e3cc137
	github.com/eyedeekay/sam3 v0.33.5
	github.com/panjf2000/ants/v2 v2.7.1
	github.com/pelletier/go-toml v1.9.5
	github.com/rs/zerolog v1.29.0
	github.com/stretchr/testify v1.8.1
	golang.org/x/sys v0.6.0
)

require (
	github.com/OneOfOne/xxhash v1.2.8 // indirect
	github.com/babolivier/go-doh-client v0.0.0-20201028162107-a76cff4cb8b6 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/cretz/bine v0.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/eyedeekay/goSam v0.32.54 // indirect
	github.com/eyedeekay/i2pkeys v0.33.0 // indirect
	github.com/eyedeekay/onramp v0.0.0-20230116055507-c93b68c0ab01 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/kentik/patricia v1.2.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.14.0 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.41.0 // indirect
	github.com/prometheus/procfs v0.9.0 // indirect
	github.com/smira/go-statsd v1.3.2 // indirect
	github.com/txthinking/runnergroup v0.0.0-20230211072751-d11f16258c86 // indirect
	github.com/txthinking/socks5 v0.0.0-20230304014913-225d9357dddc // indirect
	github.com/txthinking/x v0.0.0-20220929041811-1b4d914e9133 // indirect
	github.com/tylertreat/BoomFilters v0.0.0-20210315201527-1a82519a3e43 // indirect
	github.com/yl2chen/cidranger v1.0.2 // indirect
	golang.org/x/crypto v0.7.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/9seconds/mtg/v2@v2.1.4/mtglib => ./mtglib

replace github.com/9seconds/mtg/v2/mtglib => ./mtglib
