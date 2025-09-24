# Fractal-

rpcv2-historical/
├── cmd/
│   └── rpcv2-hist/
│       └── main.go
├── internal/
│   ├── api/
│   │   ├── jsonrpc.go
│   │   ├── handlers.go
│   │   └── middleware.go
│   ├── cache/
│   │   └── ttl_sharded.go
│   ├── clickhouse/
│   │   ├── conn.go
│   │   └── sql.go
│   ├── index/
│   │   └── fractal_idx.go
│   ├── parquet/
│   │   ├── writer.go
│   │   └── reader.go
│   ├── security/
│   │   ├── jwt.go
│   │   └── acl.go
│   └── domain/
│       └── types.go
├── spec/
│   └── openrpc.json
├── scripts/
│   ├── gen-keypair.sh
│   └── docker-compose.yml
├── go.mod
├── go.sum
└── LICENSE (AGPL-3.0)
