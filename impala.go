package impala

import (
	"database/sql"
	"io"
)

func init() {
	sql.Register("impala", &Driver{})
}

// Options for impala driver connection
type Options struct {
	Host     string
	Port     string
	Username string
	Password string

	UseLDAP          bool
	UseTLS           bool
	CACertPath       string
	BufferSize       int
	BatchSize        int
	MemoryLimit      string
	QueryTimeout     int
	FetchRowsTimeout int
	Database         string

	LogOut io.Writer
}

var (
	// DefaultOptions for impala driver
	DefaultOptions = Options{
		BatchSize:        1024,
		BufferSize:       4096,
		Port:             "21050",
		LogOut:           io.Discard,
		FetchRowsTimeout: 3 * 60 * 1000,
		Database:         "default"}
)
