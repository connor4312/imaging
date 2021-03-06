package imaging

import (
	"image"
	"testing"
)

func TestRotate90(t *testing.T) {
	td := []struct {
		desc string
		src  image.Image
		want *image.NRGBA
	}{
		{
			"Rotate90 2x3",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x11, 0x22, 0x33, 0xcc, 0xdd, 0xee, 0xff,
					0xff, 0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00,
					0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0xff,
				},
			},
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 3, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0xcc, 0xdd, 0xee, 0xff, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff,
					0x00, 0x11, 0x22, 0x33, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x00,
				},
			},
		},
	}
	for _, d := range td {
		got := Rotate90(d.src)
		want := d.want
		if !compareNRGBA(got, want, 0) {
			t.Errorf("test [%s] failed: %#v", d.desc, got)
		}
	}
}

func TestRotate180(t *testing.T) {
	td := []struct {
		desc string
		src  image.Image
		want *image.NRGBA
	}{
		{
			"Rotate180 2x3",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x11, 0x22, 0x33, 0xcc, 0xdd, 0xee, 0xff,
					0xff, 0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00,
					0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0xff,
				},
			},
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 2, 3),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff, 0x00,
					0x00, 0xff, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00,
					0xcc, 0xdd, 0xee, 0xff, 0x00, 0x11, 0x22, 0x33,
				},
			},
		},
	}
	for _, d := range td {
		got := Rotate180(d.src)
		want := d.want
		if !compareNRGBA(got, want, 0) {
			t.Errorf("test [%s] failed: %#v", d.desc, got)
		}
	}
}

func TestRotate270(t *testing.T) {
	td := []struct {
		desc string
		src  image.Image
		want *image.NRGBA
	}{
		{
			"Rotate270 2x3",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x11, 0x22, 0x33, 0xcc, 0xdd, 0xee, 0xff,
					0xff, 0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00,
					0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0xff,
				},
			},
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 3, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0x00, 0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0x11, 0x22, 0x33,
					0x00, 0x00, 0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xcc, 0xdd, 0xee, 0xff,
				},
			},
		},
	}
	for _, d := range td {
		got := Rotate270(d.src)
		want := d.want
		if !compareNRGBA(got, want, 0) {
			t.Errorf("test [%s] failed: %#v", d.desc, got)
		}
	}
}

func TestFlipV(t *testing.T) {
	td := []struct {
		desc string
		src  image.Image
		want *image.NRGBA
	}{
		{
			"FlipV 2x3",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x11, 0x22, 0x33, 0xcc, 0xdd, 0xee, 0xff,
					0xff, 0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00,
					0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0xff,
				},
			},
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 2, 3),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0xff,
					0xff, 0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00,
					0x00, 0x11, 0x22, 0x33, 0xcc, 0xdd, 0xee, 0xff,
				},
			},
		},
	}
	for _, d := range td {
		got := FlipV(d.src)
		want := d.want
		if !compareNRGBA(got, want, 0) {
			t.Errorf("test [%s] failed: %#v", d.desc, got)
		}
	}
}

func TestFlipH(t *testing.T) {
	td := []struct {
		desc string
		src  image.Image
		want *image.NRGBA
	}{
		{
			"FlipH 2x3",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x11, 0x22, 0x33, 0xcc, 0xdd, 0xee, 0xff,
					0xff, 0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00,
					0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0xff,
				},
			},
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 2, 3),
				Stride: 2 * 4,
				Pix: []uint8{
					0xcc, 0xdd, 0xee, 0xff, 0x00, 0x11, 0x22, 0x33,
					0x00, 0xff, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff, 0x00,
				},
			},
		},
	}
	for _, d := range td {
		got := FlipH(d.src)
		want := d.want
		if !compareNRGBA(got, want, 0) {
			t.Errorf("test [%s] failed: %#v", d.desc, got)
		}
	}
}
