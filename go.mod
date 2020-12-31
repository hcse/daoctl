module github.com/hypha-dao/daoctl

go 1.15

replace github.com/hypha-dao/dao-go => github.com/hypha-dao/dao-contracts/dao-go v0.0.0-20201229192602-b03263aebb57

replace github.com/hypha-dao/document-graph/docgraph => github.com/hypha-dao/document-graph/docgraph v0.0.0-20201229193929-e09f4b1c9e47

require (
	github.com/alexeyco/simpletable v0.0.0-20200730140406-5bb24159ccfb
	github.com/bronze1man/go-yaml2json v0.0.0-20150129175009-f6f64b738964
	github.com/eoscanada/eos-go v0.9.1-0.20200805141443-a9d5402a7bc5
	github.com/eoscanada/eosc v1.4.0
	github.com/hypha-dao/document-graph/docgraph v0.0.0-20201205184845-8c83c9c9a989
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
	golang.org/x/crypto v0.0.0-20201112155050-0c6587e931a9
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	google.golang.org/api v0.13.0
)
