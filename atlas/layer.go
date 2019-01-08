package atlas

import (
	"github.com/dennisrutjes/tegola/provider"
	"github.com/go-spatial/geom"
)

type Layer struct {
	// optional. if not set, the ProviderLayerName will be used
	Name              string
	ProviderLayerName string
	MinZoom           uint
	MaxZoom           uint
	// instantiated provider
	Provider provider.Tiler
	// default tags to include when encoding the layer. provider tags take precedence
	DefaultTags map[string]interface{}
	GeomType    geom.Geometry
	// DontSimplify indicates wheather feature simplification should be applied.
	// We use a negative in the name so the default is to simplify
	DontSimplify bool
}

// MVTName will return the value that will be encoded in the Name field when the layer is encoded as MVT
func (l *Layer) MVTName() string {
	if l.Name != "" {
		return l.Name
	}

	return l.ProviderLayerName
}
