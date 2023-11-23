package ingester

import (
	"fmt"
	_interface "github.com/metapox/logical-streaming-app/interface"
)

type Ingester struct {
}

func CreateIngester(ingester string) (_interface.Ingester, error) {
	switch ingester {
	case "postgres":
		return NewPostgresIngester(), nil
	}
	return nil, fmt.Errorf("ingester %s not found", ingester)
}
