package model

import (
	"github.com/welcomehyunseo/golang-3d-graphics-rendering-ex/physics/color"
)

type Triangle struct {
	//center     *vector.Vector

	color *color.Color

	indices [3]int
	//normals

	specular   float64 // 0 < n <= 1000
	reflective float64 // 0 <= n < 1
}

func NewTriangle(
	color *color.Color,
	i0, i1, i2 int,
	specular float64,
	reflective float64,
) *Triangle {
	if specular < 0 || 1000 < specular {
		// TODO
	}

	return &Triangle{
		color:      color,
		indices:    [3]int{i0, i1, i2},
		specular:   specular,
		reflective: reflective,
	}
}

func (t *Triangle) GetColor() *color.Color {
	return t.color
}

func (t *Triangle) GetIndices() [3]int {
	return t.indices
}

func (t *Triangle) GetSpecular() float64 {
	return t.specular
}

func (t *Triangle) GetReflective() float64 {
	return t.reflective
}
