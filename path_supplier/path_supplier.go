package path_supplier

import (
	m "github.com/harfmuf/geolocation-generator/model"
)

type PathSupplier interface {
	findPath(m.Location, m.Location) [][]m.Location
}
