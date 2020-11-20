// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package image

import (
	"testing"

	"github.com/h2non/bimg"
	"github.com/stretchr/testify/assert"
)

func TestOptionsHash(t *testing.T) {
	o := &Options{
		Width:  400,
		Height: 400,
	}

	assert.Equal(t, "619a9e108e52e84031672a4ce9e1588bda14b54a3a2bd3b95267544e59753014", o.Hash())

	// Test static cache
	assert.Equal(t, "619a9e108e52e84031672a4ce9e1588bda14b54a3a2bd3b95267544e59753014", o.Hash())
}

func TestOptionsToBimg(t *testing.T) {
	o := &Options{
		Width:      400,
		Height:     250,
		Background: []uint8{255, 255, 255},
	}

	assert.Equal(t, bimg.Options{
		Width:         400,
		Height:        250,
		Enlarge:       true,
		NoProfile:     true,
		StripMetadata: true,
		Background: bimg.Color{
			R: 255,
			G: 255,
			B: 255,
		},
	}, o.ToBimg())
}

func TestOptionsToBimgWithCropZone(t *testing.T) {
	o := &Options{
		Crop: CropType{
			Width:  40,
			Height: 30,
			X:      10,
			Y:      20,
		},
	}

	assert.Equal(t, bimg.Options{
		Width:         0,
		Height:        0,
		AreaHeight:    30,
		AreaWidth:     40,
		Left:          10,
		Top:           20,
		Crop:          true,
		Enlarge:       true,
		NoProfile:     true,
		StripMetadata: true,
	}, o.ToBimg())
}

func TestOptionsToBimgWithFocalPoint(t *testing.T) {
	o := &Options{
		Width:  200,
		Height: 200,
		Fit:    FitCropFocalPoint,
	}

	assert.Equal(t, bimg.Options{
		Width:         200,
		Height:        200,
		Crop:          true,
		Gravity:       bimg.GravitySmart,
		Enlarge:       true,
		NoProfile:     true,
		StripMetadata: true,
	}, o.ToBimg())
}

func TestOptionsToBimgWithBlur(t *testing.T) {
	o := &Options{
		Width:  200,
		Height: 200,
		Blur:   5,
	}

	assert.Equal(t, bimg.Options{
		Width:         200,
		Height:        200,
		Crop:          false,
		Enlarge:       true,
		NoProfile:     true,
		StripMetadata: true,
		GaussianBlur: bimg.GaussianBlur{
			Sigma:   2.5,
			MinAmpl: 5,
		},
	}, o.ToBimg())
}
