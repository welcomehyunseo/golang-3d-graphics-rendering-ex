package model

import (
	"github.com/welcomehyunseo/golang-3d-graphics-rendering-ex/physics/vector"
)git stat

type Model struct {
	vertices  []*vector.Vector
	triangles []*Triangle
}

func NewModel(
	vertices []*vector.Vector,
	triangles []*Triangle,
) *Model {

	return &Model{
		vertices:  vertices,
		triangles: triangles,
	}
}

func (m *Model) GetVertices() []*vector.Vector {
	return m.vertices
}

func (m *Model) GetTriangles() []*Triangle {
	return m.triangles
}
