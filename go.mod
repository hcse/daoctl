module github.com/hypha-dao/daoctl

go 1.15

require (
	github.com/alexeyco/simpletable v0.0.0-20200730140406-5bb24159ccfb
	github.com/bronze1man/go-yaml2json v0.0.0-20150129175009-f6f64b738964
	github.com/dgraph-io/dgo/v200 v200.0.0-20201023081658-a9ad93fe6ebd
	github.com/eoscanada/eos-go v0.9.1-0.20200805141443-a9d5402a7bc5
	github.com/eoscanada/eosc v1.4.0
	github.com/hypha-dao/dao-contracts/dao-go v0.0.0-00010101000000-000000000000
	github.com/hypha-dao/document-graph/docgraph v0.0.0-20201229193929-e09f4b1c9e47
	github.com/leekchan/accounting v1.0.0
	github.com/manifoldco/promptui v0.8.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.8.0
	github.com/ryanuber/columnize v2.1.2+incompatible
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	github.com/tidwall/gjson v1.6.3
	github.com/tidwall/pretty v1.0.2
	github.com/tidwall/sjson v1.1.2
	go.uber.org/zap v1.16.0
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
	golang.org/x/net v0.0.0-20201224014010-6772e930b67b // indirect
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/sys v0.0.0-20210110051926-789bb1bd4061 // indirect
	golang.org/x/text v0.3.5 // indirect
	google.golang.org/api v0.13.0
	google.golang.org/genproto v0.0.0-20210108203827-ffc7fda8c3d7 // indirect
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.3.0
)

replace github.com/hypha-dao/dao-contracts/dao-go => ../dao-contracts/dao-go

replace github.com/hypha-dao/document-graph/docgraph => ../dao-contracts/document-graph/docgraph
