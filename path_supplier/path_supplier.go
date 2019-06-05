package path_supplier

import (
	m "github.com/harfmuf/geolocation-generator/model"
)

type PathSupplier interface {
	FindPath(*m.Location, *m.Location, string) ([][]*m.Location, error)
}
