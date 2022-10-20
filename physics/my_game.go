package physics

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/welcomehyunseo/golang-3d-graphics-rendering-ex/physics/color"
	"github.com/welcomehyunseo/golang-3d-graphics-rendering-ex/physics/light"
	"github.com/welcomehyunseo/golang-3d-graphics-rendering-ex/physics/model"
	"github.com/welcomehyunseo/golang-3d-graphics-rendering-ex/physics/vector"
)

const (
	DefaultRecursionDepth = 3 // for reflections
)

var (
	BackgroundColor = color.NewColor(255, 255, 255)
)

type Viewport struct {
	width  int
	height int
}

type Camera struct {
	origin             *vector.Vector
	viewport           *Viewport
	distanceToViewport float64
	viewDistance       float64
}

type MyGame struct {
	framebuffer []*color.Color
	lights      []light.Light
	models      []*model.Model
	camera      *Camera
}

func NewMyGame(cameraCenter *vector.Vector, width, height int, distanceToViewport, viewDistance float64) *MyGame {
	framebuffer := make([]*color.Color, width*height)
	for i, _ := range framebuffer {
		framebuffer[i] = BackgroundColor
	}

	return &MyGame{
		framebuffer: framebuffer,
		models:      []*model.Model{},
		camera: &Camera{
			cameraCenter,
			&Viewport{width, height},
			distanceToViewport,
			viewDistance,
		},
	}
}

func (g *MyGame) AddLight(light light.Light) {
	g.lights = append(g.lights, light)
}

func (g *MyGame) AddModel(model *model.Model) {
	g.models = append(g.models, model)
}

func (g *MyGame) UpdateFramebuffer() {
	models := g.models
	for _, m := range models {
		vertices := m.GetVertices()
		triangles := m.GetTriangles()
		for _, t := range triangles {
			indices := t.GetIndices()
			i0 := indices[0]
			i1 := indices[1]
			i2 := indices[2]
			v0 := vertices[i0]
			v1 := vertices[i1]
			v2 := vertices[i2]

		}
	}
}

func (g *MyGame) Update() error {
	return nil
}

func (g *MyGame) Draw(screen *ebiten.Image) {

	vh := g.camera.viewport.height
	vw := g.camera.viewport.width

	for l := 0; l < vh; l++ {
		for k := 0; k < vw; k++ {
			i := l*vw + k
			c := g.framebuffer[i]
			screen.Set(k, l, c.ToRGBA())
		}
	}
}

func (g *MyGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
