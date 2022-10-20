package golang_3d_graphics_rendering_ex

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/welcomehyunseo/golang-3d-graphics-rendering-ex/physics"
	"github.com/welcomehyunseo/golang-3d-graphics-rendering-ex/physics/color"
	"github.com/welcomehyunseo/golang-3d-graphics-rendering-ex/physics/light"
	"github.com/welcomehyunseo/golang-3d-graphics-rendering-ex/physics/model"
	"github.com/welcomehyunseo/golang-3d-graphics-rendering-ex/physics/vector"
)

const (
	Width               = 1024
	Height              = 768
	DistanceToViewport  = 500
	DefaultViewDistance = 2000
)

func main() {
	ebiten.SetWindowSize(Width, Height)
	ebiten.SetWindowTitle("Basic Rendering")

	cameraCenter := vector.NewVector(0, 0, 0)
	myGame := physics.NewMyGame(cameraCenter, Width, Height, DistanceToViewport, DefaultViewDistance)

	m0 := model.NewModel(
		[]*vector.Vector{
			vector.NewVector(0, 0, 1000),
			vector.NewVector(0, 1000, 1000),
			vector.NewVector(1000, 0, 1000),
		},
		[]*model.Triangle{
			model.NewTriangle(
				color.NewColor(0xff, 0, 0),
				0, 1, 2,
				300,
				0,
			),
		},
	)
	myGame.AddModel(m0)

	l0 := light.NewAmbientLight(0.2)
	myGame.AddLight(l0)
	l1 := light.NewPointLight(0.3, vector.NewVector(4000, 0, 0))
	myGame.AddLight(l1)
	l2 := light.NewDirectionalLight(0.2, vector.NewVector(0, -1, 0))
	myGame.AddLight(l2)

	go func() {
		for {
			myGame.UpdateFramebuffer()
		}
	}()

	if err := ebiten.RunGame(myGame); err != nil {
		panic(err)
	}
}
