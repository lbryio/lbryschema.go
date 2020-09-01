module github.com/lbryio/lbryschema.go

go 1.14

replace github.com/btcsuite/btcd => github.com/lbryio/lbrycrd.go v0.0.0-20200203050410-e1076f12bf19

require (
	github.com/btcsuite/btcd v0.0.0-20190213025234-306aecffea32
	github.com/btcsuite/btcutil v0.0.0-20190425235716-9e5f4b9a998d
	github.com/golang/protobuf v1.4.2
	github.com/lbryio/lbry.go/v2 v2.6.1-0.20200710180140-fcade7475323
	github.com/lbryio/types v0.0.0-20200605192618-366870b2862d
	gotest.tools v2.2.0+incompatible
)
