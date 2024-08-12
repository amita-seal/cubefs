module github.com/cubefs/cubefs

go 1.15

//replace (
//	github.com/jacobsa/fuse => ./depends/jacobsa/fuse
//	github.com/spf13/cobra => ./depends/spf13/cobra
//)

require (
	github.com/Shopify/sarama v1.33.0
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/aws/aws-sdk-go v1.33.1
	github.com/benbjohnson/clock v1.3.1
	github.com/bits-and-blooms/bitset v1.2.1
	github.com/brahma-adshonor/gohook v1.1.9
	github.com/deniswernert/go-fstab v0.0.0-20141204152952-eb4090f26517
	github.com/desertbit/grumble v1.1.3
	github.com/dustin/go-humanize v1.0.1
	github.com/edsrzf/mmap-go v1.1.0
	github.com/fatih/color v1.15.0
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.3
	github.com/google/btree v1.0.1
	github.com/google/uuid v1.3.0
	github.com/gorilla/mux v1.8.0
	github.com/hashicorp/consul/api v1.15.2
	github.com/hashicorp/golang-lru v0.5.4
	github.com/jacobsa/daemonize v0.0.0-20160101105449-e460293e890f
	github.com/julienschmidt/httprouter v1.3.0
	github.com/klauspost/reedsolomon v1.11.7
	github.com/opentracing/opentracing-go v1.2.0
	github.com/peterbourgon/diskv/v3 v3.0.1
	github.com/prometheus/client_golang v1.13.0
	github.com/rs/xid v1.5.0
	github.com/samsarahq/thunder v0.0.0-20211005041752-96f4331b7baa
	github.com/spf13/cobra v1.2.1
	github.com/stretchr/testify v1.8.4
	github.com/tecbot/gorocksdb v0.0.0-20191217155057-f0fad39f321c
	github.com/xtaci/smux v1.5.16
	go.etcd.io/etcd/raft/v3 v3.5.8
	go.uber.org/automaxprocs v1.5.1
	golang.org/x/net v0.27.0
	golang.org/x/sync v0.7.0
	golang.org/x/sys v0.22.0
	golang.org/x/time v0.0.0-20211116232009-f0f3c7e86c11
	gopkg.in/bsm/ratelimit.v1 v1.0.0-20170922094635-f56db5e73a5e
	gopkg.in/go-playground/validator.v9 v9.31.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/tklauser/go-sysconf v0.3.11 // indirect
	github.com/yusufpapurcu/wmi v1.2.3 // indirect
)

require (
	github.com/BurntSushi/toml v1.1.0 // indirect
	github.com/armon/go-metrics v0.4.1 // indirect
	github.com/eapache/go-resiliency v1.3.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20230111030713-bf00bc1b83b6 // indirect
	github.com/facebookgo/ensure v0.0.0-20200202191622-63f1cf65ac4c // indirect
	github.com/facebookgo/stack v0.0.0-20160209184415-751773369052 // indirect
	github.com/facebookgo/subset v0.0.0-20200203212716-c811ad88dec4 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/graphql-go/graphql v0.8.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.5.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/serf v0.10.1 // indirect
	github.com/leodido/go-urn v1.2.3 // indirect
	github.com/mattn/go-isatty v0.0.18 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.34.1 // indirect
	github.com/shirou/gopsutil v3.21.11+incompatible
	github.com/smartystreets/goconvey v1.8.0 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
)
