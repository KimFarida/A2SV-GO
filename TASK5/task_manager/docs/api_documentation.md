`go get github.com/quic-go/qpack@v0.5.1`
go: downgraded github.com/quic-go/qpack v0.6.0 => v0.5.1`

=> The dependency `quic-go v0.560` was written against an older version of a library, but go get pulls qpack v0.6.0. so adowngrade i necessary to avoid errors