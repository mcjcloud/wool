
# Wool

A simple path finding tool for Go dependency trees.


## Install

Install using Go

```bash
$ go install github.com/mcjcloud/wool
```
    
## Usage

An example with only one path
```bash
$ go mod graph | wool --src "github.com/mcjcloud/taurine" --dst "golang.org/x/text@v0.3.3"
constructing graph... done.
finding path from 'github.com/mcjcloud/taurine' -> 'golang.org/x/text@v0.3.3'
> 5 hops
github.com/mcjcloud/taurine -> github.com/llir/ll@v0.0.0-20210719001141-246f2b6b1fa9 -> golang.org/x/tools@v0.1.4 -> golang.org/x/net@v0.0.0-20210405180319-a5a99cb37ef4 -> golang.org/x/text@v0.3.3
```

An example with multiple paths
```bash
go mod graph | wool --src "github.com/sonr-io/sonr" --dst "github.com/99designs/keyring@v1.1.6"
constructing graph... done.
finding path from 'github.com/sonr-io/sonr' -> 'github.com/99designs/keyring@v1.1.6'
found 18 paths.
> 3 hops
github.com/sonr-io/sonr -> github.com/tendermint/spn@v0.2.1-0.20220609194312-7833ecf4454a -> github.com/99designs/keyring@v1.1.6
> 4 hops
github.com/sonr-io/sonr -> github.com/tendermint/spn@v0.2.1-0.20220609194312-7833ecf4454a -> github.com/cosmos/cosmos-sdk@v0.45.4 -> github.com/99designs/keyring@v1.1.6
> 5 hops
github.com/sonr-io/sonr -> github.com/tendermint/spn@v0.2.1-0.20220609194312-7833ecf4454a -> github.com/cosmos/ibc-go/v3@v3.0.0 -> github.com/cosmos/cosmos-sdk@v0.45.1 -> github.com/99designs/keyring@v1.1.6
> 4 hops
github.com/sonr-io/sonr -> github.com/tendermint/spn@v0.2.1-0.20220609194312-7833ecf4454a -> github.com/cosmos/ibc-go/v3@v3.0.0 -> github.com/99designs/keyring@v1.1.6
> 6 hops
github.com/sonr-io/sonr -> github.com/tendermint/spn@v0.2.1-0.20220609194312-7833ecf4454a -> github.com/ignite-hq/cli@v0.21.0 -> github.com/cosmos/ibc-go/v2@v2.0.3 -> github.com/cosmos/cosmos-sdk@v0.44.5 -> github.com/99designs/keyring@v1.1.6
> 6 hops
github.com/sonr-io/sonr -> github.com/tendermint/spn@v0.2.1-0.20220609194312-7833ecf4454a -> github.com/ignite-hq/cli@v0.21.0 -> github.com/tendermint/spn@v0.2.1-0.20220427143342-de7398284030 -> github.com/ignite-hq/cli@v0.20.0 -> github.com/99designs/keyring@v1.1.6
> 7 hops
github.com/sonr-io/sonr -> github.com/tendermint/spn@v0.2.1-0.20220609194312-7833ecf4454a -> github.com/ignite-hq/cli@v0.21.0 -> github.com/tendermint/spn@v0.2.1-0.20220427143342-de7398284030 -> github.com/ignite-hq/cli@v0.20.0 -> github.com/cosmos/cosmos-sdk@v0.45.3 -> github.com/99designs/keyring@v1.1.6
> 8 hops
github.com/sonr-io/sonr -> github.com/tendermint/spn@v0.2.1-0.20220609194312-7833ecf4454a -> github.com/ignite-hq/cli@v0.21.0 -> github.com/tendermint/spn@v0.2.1-0.20220427143342-de7398284030 -> github.com/ignite-hq/cli@v0.20.0 -> github.com/tendermint/spn@v0.1.1-0.20220407154406-5cfd1bf28150 -> github.com/tendermint/starport@v0.19.5 -> github.com/99designs/keyring@v1.1.6
> 11 hops
github.com/sonr-io/sonr -> github.com/tendermint/spn@v0.2.1-0.20220609194312-7833ecf4454a -> github.com/ignite-hq/cli@v0.21.0 -> github.com/tendermint/spn@v0.2.1-0.20220427143342-de7398284030 -> github.com/ignite-hq/cli@v0.20.0 -> github.com/tendermint/spn@v0.1.1-0.20220407154406-5cfd1bf28150 -> github.com/tendermint/starport@v0.19.5 -> github.com/tendermint/spn@v0.1.1-0.20211210094128-4ca78a240c57 -> github.com/tendermint/spm@v0.1.8 -> github.com/cosmos/cosmos-sdk@v0.44.3 -> github.com/99designs/keyring@v1.1.6
> 12 hops
github.com/sonr-io/sonr -> github.com/tendermint/spn@v0.2.1-0.20220609194312-7833ecf4454a -> github.com/ignite-hq/cli@v0.21.0 -> github.com/tendermint/spn@v0.2.1-0.20220427143342-de7398284030 -> github.com/ignite-hq/cli@v0.20.0 -> github.com/tendermint/spn@v0.1.1-0.20220407154406-5cfd1bf28150 -> github.com/tendermint/starport@v0.19.5 -> github.com/tendermint/spn@v0.1.1-0.20211210094128-4ca78a240c57 -> github.com/tendermint/spm@v0.1.8 -> github.com/cosmos/ibc-go@v1.2.2 -> github.com/cosmos/cosmos-sdk@v0.44.2 -> github.com/99designs/keyring@v1.1.6
> 10 hops
github.com/sonr-io/sonr -> github.com/tendermint/spn@v0.2.1-0.20220609194312-7833ecf4454a -> github.com/ignite-hq/cli@v0.21.0 -> github.com/tendermint/spn@v0.2.1-0.20220427143342-de7398284030 -> github.com/ignite-hq/cli@v0.20.0 -> github.com/tendermint/spn@v0.1.1-0.20220407154406-5cfd1bf28150 -> github.com/tendermint/starport@v0.19.5 -> github.com/tendermint/spn@v0.1.1-0.20211210094128-4ca78a240c57 -> github.com/cosmos/cosmos-sdk@v0.44.4 -> github.com/99designs/keyring@v1.1.6
> 8 hops
github.com/sonr-io/sonr -> github.com/tendermint/spn@v0.2.1-0.20220609194312-7833ecf4454a -> github.com/ignite-hq/cli@v0.21.0 -> github.com/tendermint/spn@v0.2.1-0.20220427143342-de7398284030 -> github.com/ignite-hq/cli@v0.20.0 -> github.com/tendermint/spn@v0.1.1-0.20220407154406-5cfd1bf28150 -> github.com/tendermint/fundraising@v0.2.0 -> github.com/99designs/keyring@v1.1.6
> 7 hops
github.com/sonr-io/sonr -> github.com/tendermint/spn@v0.2.1-0.20220609194312-7833ecf4454a -> github.com/ignite-hq/cli@v0.21.0 -> github.com/tendermint/spn@v0.2.1-0.20220427143342-de7398284030 -> github.com/ignite-hq/cli@v0.20.0 -> github.com/tendermint/spn@v0.1.1-0.20220407154406-5cfd1bf28150 -> github.com/99designs/keyring@v1.1.6
> 8 hops
github.com/sonr-io/sonr -> github.com/tendermint/spn@v0.2.1-0.20220609194312-7833ecf4454a -> github.com/ignite-hq/cli@v0.21.0 -> github.com/tendermint/spn@v0.2.1-0.20220427143342-de7398284030 -> github.com/ignite-hq/cli@v0.20.0 -> github.com/tendermint/spn@v0.1.1-0.20220407154406-5cfd1bf28150 -> github.com/cosmos/cosmos-sdk@v0.45.2 -> github.com/99designs/keyring@v1.1.6
> 5 hops
github.com/sonr-io/sonr -> github.com/tendermint/spn@v0.2.1-0.20220609194312-7833ecf4454a -> github.com/ignite-hq/cli@v0.21.0 -> github.com/tendermint/spn@v0.2.1-0.20220427143342-de7398284030 -> github.com/99designs/keyring@v1.1.6
> 4 hops
github.com/sonr-io/sonr -> github.com/tendermint/spn@v0.2.1-0.20220609194312-7833ecf4454a -> github.com/ignite-hq/cli@v0.21.0 -> github.com/99designs/keyring@v1.1.6
> 4 hops
github.com/sonr-io/sonr -> github.com/tendermint/spn@v0.2.1-0.20220609194312-7833ecf4454a -> github.com/tendermint/fundraising@v0.3.0 -> github.com/99designs/keyring@v1.1.6
> 3 hops
github.com/sonr-io/sonr -> github.com/ignite-hq/cli@v0.22.0 -> github.com/99designs/keyring@v1.1.6
```
## License

[MIT](https://choosealicense.com/licenses/mit/)

