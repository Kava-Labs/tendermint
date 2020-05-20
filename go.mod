module github.com/kava-labs/tendermint

go 1.13

require (
	github.com/ChainSafe/go-schnorrkel v0.0.0-20200102211924-4bcbc698314f
	github.com/Workiva/go-datastructures v1.0.52
	github.com/btcsuite/btcd v0.0.0-20190115013929-ed77733ec07d
	github.com/btcsuite/btcutil v0.0.0-20180706230648-ab6388e0c60a
	github.com/fortytw2/leaktest v1.3.0
	github.com/go-kit/kit v0.10.0
	github.com/go-logfmt/logfmt v0.5.0
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.3.4
	github.com/gorilla/websocket v1.4.1
	github.com/gtank/merlin v0.1.1-0.20191105220539-8318aed1a79f
	github.com/kava-labs/tm-db v0.4.2-0.20200506040135-3f7b09feebcd
	github.com/libp2p/go-buffer-pool v0.0.2
	github.com/magiconair/properties v1.8.1
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.5.0
	github.com/rcrowley/go-metrics v0.0.0-20181016184325-3113b8401b8a
	github.com/rs/cors v1.7.0
	github.com/snikch/goodman v0.0.0-20171125024755-10e37e294daa
	github.com/spf13/cobra v0.0.6
	github.com/spf13/viper v1.6.2
	github.com/stretchr/testify v1.5.1
	github.com/tendermint/go-amino v0.14.1
	go.etcd.io/bbolt v1.3.4 // indirect
	golang.org/x/crypto v0.0.0-20191206172530-e9b2fee46413
	golang.org/x/net v0.0.0-20190813141303-74dc4d7220e7
	google.golang.org/grpc v1.27.1
)

replace github.com/tendermint/tm-db => github.com/kava-labs/tm-db v0.4.2-0.20200506040135-3f7b09feebcd
