package config

type GRPCConfig interface {
	Address() string
}

type PgConfig interface {
	DSN() string
}

type AuthGRPCClientConfig interface {
	Address() string
}
