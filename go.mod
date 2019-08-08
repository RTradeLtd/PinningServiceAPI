module github.com/RTradeLtd/PinningServiceAPI

go 1.12

require (
	github.com/RTradeLtd/go-ipfs-api v0.0.0-20190523020607-76503b15fe41 // indirect
	github.com/RTradeLtd/rtfs/v2 v2.2.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/ipfs/go-ds-badger v0.0.5 // indirect
	github.com/ipfs/go-ipfs-files v0.0.3 // indirect
	github.com/libp2p/go-libp2p-metrics v0.1.0 // indirect
	github.com/libp2p/go-libp2p-peer v0.2.0 // indirect
	github.com/multiformats/go-multiaddr-dns v0.0.2 // indirect
	github.com/opentracing/opentracing-go v1.1.0 // indirect
	golang.org/x/crypto v0.0.0-20190611184440-5c40567a22f8 // indirect
	golang.org/x/text v0.3.2 // indirect
	gopkg.in/dgrijalva/jwt-go.v3 v3.2.0
)

replace github.com/dgraph-io/badger v2.0.0-rc.2+incompatible => github.com/dgraph-io/badger v1.6.0-rc1

replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43
