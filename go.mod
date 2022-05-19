module github.com/ipfs-force-community/venus-gateway

go 1.16

require (
	github.com/filecoin-project/go-address v0.0.6
	github.com/filecoin-project/go-jsonrpc v0.1.5
	github.com/filecoin-project/go-state-types v0.1.7
	github.com/filecoin-project/specs-actors/v5 v5.0.6-0.20220514165557-0b29a778685b
	github.com/filecoin-project/specs-storage v0.4.0
	github.com/filecoin-project/venus v1.3.0-rc2.0.20220530152607-c0a964715602
	github.com/filecoin-project/venus-auth v1.4.0
	github.com/gbrlsnchs/jwt/v3 v3.0.1
	github.com/google/uuid v1.3.0
	github.com/gorilla/mux v1.8.0
	github.com/ipfs-force-community/metrics v1.0.1-0.20211228055608-9462dc86e157
	github.com/ipfs/go-cid v0.1.0
	github.com/ipfs/go-log/v2 v2.5.0
	github.com/modern-go/reflect2 v1.0.2
	github.com/multiformats/go-multiaddr v0.5.0
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.1
	github.com/urfave/cli/v2 v2.3.0
	go.opencensus.io v0.23.0
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
)

replace github.com/filecoin-project/filecoin-ffi => ./extern/filecoin-ffi

replace github.com/ipfs-force-community/venus-gateway => ./

replace github.com/filecoin-project/go-jsonrpc => github.com/ipfs-force-community/go-jsonrpc v0.1.4-0.20210731021807-68e5207079bc
