package middleware

import (
	"myapp/data"

	"github.com/cjfloss/celeritas"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}
