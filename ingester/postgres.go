package ingester

import "fmt"

type postgresIngester struct {
}

func (p *postgresIngester) Start() {
	fmt.Println("Starting Postgres Ingester")
}

func NewPostgresIngester() *postgresIngester {
	return &postgresIngester{}
}
