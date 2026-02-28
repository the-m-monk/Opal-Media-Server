package librarymgmt

import (
	"fmt"
	"image/color"
	"io"
	"log/slog"
	"os"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

var nameCardFont *truetype.Font

func initNameCardRenderer() {
	//TODO: make fonts configurable
	fontPath := "assets/fonts/IBM_Plex_Sans/static/IBMPlexSans-Regular.ttf"

	fontBytes, err := os.ReadFile(fontPath)
	if err != nil {
		slog.Error("failed to load font, library namecards will not render correctly", "fontPath", fontPath)
		return
	}

	f, err := truetype.Parse(fontBytes)
	if err != nil {
		slog.Error("failed to parse font, library namecards will not render correctly", "fontPath", fontPath)
		return
	}

	nameCardFont = f
}

func RenderNameCard(libraryName string, outputPath string) error {
	if nameCardFont == nil {
		return fmt.Errorf("nameCardFont is nil, unable to render name cards")
	}

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

	fontSize := 150.0
	maxWidth := float64(width) - (margin * 2)

	var finalFace font.Face

	for fontSize > 5 {
		face := truetype.NewFace(nameCardFont, &truetype.Options{
			Size: fontSize,
		})

		dc.SetFontFace(face)
		w, _ := dc.MeasureString(libraryName)

		if w <= maxWidth {
			finalFace = face
			break
		}

		if closer, ok := face.(io.Closer); ok {
			closer.Close()
		}
		fontSize -= 5
	}

	dc.SetRGB(1, 1, 1)
	dc.DrawStringAnchored(libraryName, float64(width)/2, float64(height)/2, 0.5, 0.5)

	if finalFace != nil {
		if closer, ok := finalFace.(io.Closer); ok {
			closer.Close()
		}
	}

	dc.SetRGB(1, 1, 1)
	dc.DrawStringAnchored(libraryName, float64(width)/2, float64(height)/2, 0.5, 0.5)

	return dc.SavePNG(outputPath)
}
