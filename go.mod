module github.com/eyedeekay/mtg-i2p

go 1.17

require (
	github.com/9seconds/mtg/v2 v2.1.4
	github.com/alecthomas/kong v0.2.22
	github.com/alecthomas/units v0.0.0-20211218093645-b94a6e3cc137
	github.com/eyedeekay/sam3 v0.32.32
	github.com/panjf2000/ants/v2 v2.4.7
	github.com/pelletier/go-toml v1.9.4
	github.com/rs/zerolog v1.26.1
	github.com/stretchr/testify v1.7.0
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e
)

require (
	github.com/OneOfOne/xxhash v1.2.8 // indirect
	github.com/babolivier/go-doh-client v0.0.0-20201028162107-a76cff4cb8b6 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/kentik/patricia v0.0.0-20210909164817-21603333b70e // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.11.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.32.1 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/smira/go-statsd v1.3.2 // indirect
	github.com/txthinking/runnergroup v0.0.0-20210608031112-152c7c4432bf // indirect
	github.com/txthinking/socks5 v0.0.0-20211121111206-e03c1217a50b // indirect
	github.com/txthinking/x v0.0.0-20210326105829-476fab902fbe // indirect
	github.com/tylertreat/BoomFilters v0.0.0-20210315201527-1a82519a3e43 // indirect
	golang.org/x/crypto v0.0.0-20211215165025-cf75a172585e // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/9seconds/mtg/v2/mtglib => ./mtglib/
