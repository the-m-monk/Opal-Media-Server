package librarymgmt

import (
	"image/color"

	"github.com/fogleman/gg"
)

func RenderNameCard(libraryName string, outputPath string) error {
	const (
		width  = 400
		height = 225
		margin = 20.0
	)
	dc := gg.NewContext(width, height)

	grad := gg.NewLinearGradient(0, 0, width, height)
	grad.AddColorStop(0, color.RGBA{48, 10, 86, 255})
	grad.AddColorStop(1, color.RGBA{138, 43, 226, 255})
	dc.SetFillStyle(grad)
	dc.DrawRectangle(0, 0, float64(width), float64(height))
	dc.Fill()

	//TODO: make fonts configurable
	fontPath := "assets/fonts/IBM_Plex_Sans/static/IBMPlexSans-Regular.ttf"
	fontSize := 150.0
	maxWidth := float64(width) - (margin * 2)

	for fontSize > 5 {
		if err := dc.LoadFontFace(fontPath, fontSize); err != nil {
			return err
		}

		w, _ := dc.MeasureString(libraryName)
		if w <= maxWidth {
			break
		}
		fontSize -= 2
	}

	dc.SetRGB(1, 1, 1)
	dc.DrawStringAnchored(libraryName, width/2, height/2, 0.5, 0.5)

	return dc.SavePNG(outputPath)
}
