package imaging

import (
	"image"
	"image/color"
	"math"
)

// Small number that we look at to compare floating points.
const EPSILON = 0.0000001

// Rotate90 rotates the image 90 degrees counterclockwise and returns the transformed image.
func Rotate90(img image.Image) *image.NRGBA {
	src := toNRGBA(img)
	srcW := src.Bounds().Max.X
	srcH := src.Bounds().Max.Y
	dstW := srcH
	dstH := srcW
	dst := image.NewNRGBA(image.Rect(0, 0, dstW, dstH))

	parallel(dstH, func(partStart, partEnd int) {

		for dstY := partStart; dstY < partEnd; dstY++ {
			for dstX := 0; dstX < dstW; dstX++ {
				srcX := dstH - dstY - 1
				srcY := dstX

				srcOff := srcY*src.Stride + srcX*4
				dstOff := dstY*dst.Stride + dstX*4

				copy(dst.Pix[dstOff:dstOff+4], src.Pix[srcOff:srcOff+4])
			}
		}

	})

	return dst
}

// Rotate180 rotates the image 180 degrees counterclockwise and returns the transformed image.
func Rotate180(img image.Image) *image.NRGBA {
	src := toNRGBA(img)
	srcW := src.Bounds().Max.X
	srcH := src.Bounds().Max.Y
	dstW := srcW
	dstH := srcH
	dst := image.NewNRGBA(image.Rect(0, 0, dstW, dstH))

	parallel(dstH, func(partStart, partEnd int) {

		for dstY := partStart; dstY < partEnd; dstY++ {
			for dstX := 0; dstX < dstW; dstX++ {
				srcX := dstW - dstX - 1
				srcY := dstH - dstY - 1

				srcOff := srcY*src.Stride + srcX*4
				dstOff := dstY*dst.Stride + dstX*4

				copy(dst.Pix[dstOff:dstOff+4], src.Pix[srcOff:srcOff+4])
			}
		}

	})

	return dst
}

// Rotate270 rotates the image 270 degrees counterclockwise and returns the transformed image.
func Rotate270(img image.Image) *image.NRGBA {
	src := toNRGBA(img)
	srcW := src.Bounds().Max.X
	srcH := src.Bounds().Max.Y
	dstW := srcH
	dstH := srcW
	dst := image.NewNRGBA(image.Rect(0, 0, dstW, dstH))

	parallel(dstH, func(partStart, partEnd int) {

		for dstY := partStart; dstY < partEnd; dstY++ {
			for dstX := 0; dstX < dstW; dstX++ {
				srcX := dstY
				srcY := dstW - dstX - 1

				srcOff := srcY*src.Stride + srcX*4
				dstOff := dstY*dst.Stride + dstX*4

				copy(dst.Pix[dstOff:dstOff+4], src.Pix[srcOff:srcOff+4])
			}
		}

	})

	return dst
}

// FlipH flips the image horizontally (from left to right) and returns the transformed image.
func FlipH(img image.Image) *image.NRGBA {
	src := toNRGBA(img)
	srcW := src.Bounds().Max.X
	srcH := src.Bounds().Max.Y
	dstW := srcW
	dstH := srcH
	dst := image.NewNRGBA(image.Rect(0, 0, dstW, dstH))

	parallel(dstH, func(partStart, partEnd int) {

		for dstY := partStart; dstY < partEnd; dstY++ {
			for dstX := 0; dstX < dstW; dstX++ {
				srcX := dstW - dstX - 1
				srcY := dstY

				srcOff := srcY*src.Stride + srcX*4
				dstOff := dstY*dst.Stride + dstX*4

				copy(dst.Pix[dstOff:dstOff+4], src.Pix[srcOff:srcOff+4])
			}
		}

	})

	return dst
}

// FlipV flips the image vertically (from top to bottom) and returns the transformed image.
func FlipV(img image.Image) *image.NRGBA {
	src := toNRGBA(img)
	srcW := src.Bounds().Max.X
	srcH := src.Bounds().Max.Y
	dstW := srcW
	dstH := srcH
	dst := image.NewNRGBA(image.Rect(0, 0, dstW, dstH))

	parallel(dstH, func(partStart, partEnd int) {

		for dstY := partStart; dstY < partEnd; dstY++ {
			for dstX := 0; dstX < dstW; dstX++ {
				srcX := dstX
				srcY := dstH - dstY - 1

				srcOff := srcY*src.Stride + srcX*4
				dstOff := dstY*dst.Stride + dstX*4

				copy(dst.Pix[dstOff:dstOff+4], src.Pix[srcOff:srcOff+4])
			}
		}

	})

	return dst
}

// Scews the image around its center, adjusting for the new max width/height
func Skew(img image.Image, xOffset, yOffset int) *image.NRGBA {
	src := toNRGBA(img)
	srcW := src.Bounds().Max.X
	srcH := src.Bounds().Max.Y
	dstW := srcW + xOffset
	dstH := srcH + yOffset
	dst := image.NewNRGBA(image.Rect(0, 0, dstW, dstH))

	srcHalfW := float64(srcW) / 2
	srcHalfH := float64(srcH) / 2
	dstHalfW := float64(dstW) / 2
	dstHalfH := float64(dstH) / 2
	xOffsetFloat := float64(xOffset)
	yOffsetFloat := float64(yOffset)

	parallel(dstW*dstH, func(partStart, partEnd int) {
		for dstPixel := partStart; dstPixel < partEnd; dstPixel++ {
			// Desination coordinates on the skewed image
			dstY := float64(dstPixel / dstH) // Integer division
			dstX := float64(dstPixel - (dstPixel % dstH))

			// The change at the position caused by the screw
			xDelta := (dstX/dstHalfW - 1) * xOffsetFloat
			yDelta := (dstY/dstHalfH - 1) * yOffsetFloat

			// Based on that, get the destination positions
			srcX := int(srcHalfW + xDelta)
			srcY := int(srcHalfH + yDelta)

			offset := src.PixOffset(int(srcX), int(srcY))
			aliasedAdd(dst, dstX, dstY, color.NRGBA{
				R: src.Pix[offset+0],
				G: src.Pix[offset+1],
				B: src.Pix[offset+2],
				A: src.Pix[offset+3],
			})
		}
	})

	return dst
}

// Adds the color on the image in a bicubic fashion.
func aliasedAdd(img *image.NRGBA, x, y float64, col color.NRGBA) {
	aliasedAddPixel(img, x, y, int(x), int(y), col)
	aliasedAddPixel(img, x, y, int(x), int(y)+1, col)
	aliasedAddPixel(img, x, y, int(x)+1, int(y), col)
	aliasedAddPixel(img, x, y, int(x)+1, int(y)+1, col)
}

func aliasedAddPixel(img *image.NRGBA, x, y float64, px, py int, col color.NRGBA) {
	offset := img.PixOffset(px, py)
	opacity := float64(255*col.A) * (math.Pow(float64(px)-x, 2) + math.Pow(float64(py)-y, 2))

	if opacity < 0 {
		return
	}

	if img.Pix[offset+3] == 0 {
		// If this pixel has *no* color yet (alpha = 0), set it exactly
		img.Pix[offset+0] = col.R
		img.Pix[offset+1] = col.G
		img.Pix[offset+2] = col.B
	} else {
		// Otherwise mix it in
		img.Pix[offset+0] = (col.R + img.Pix[offset+0]) / 2
		img.Pix[offset+1] = (col.G + img.Pix[offset+1]) / 2
		img.Pix[offset+3] = (col.B + img.Pix[offset+3]) / 2
	}

	// Finally adjust the opacity
	img.Pix[offset+3] += uint8(opacity)
}
