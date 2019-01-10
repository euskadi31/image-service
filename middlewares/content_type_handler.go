// Copyright 2017 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package middlewares

import (
	"net/http"

	"github.com/h2non/bimg"
	"github.com/hyperscale/hyperpic/httputil"
	"github.com/hyperscale/hyperpic/image"
	"github.com/rs/zerolog/log"
)

// NewContentTypeHandler negotiate content type
func NewContentTypeHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			options, err := OptionsFromContext(ctx)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)

				return
			}

			var mime string

			if options.Format == bimg.UNKNOWN {
				mime = httputil.NegotiateContentType(r, []string{
					"image/jpeg",
					"image/webp",
					"image/jpeg",
					"image/tiff",
					"image/png",
				}, "image/jpeg")

				format := image.ExtractImageTypeFromMime(mime)

				log.Debug().Msgf("Format extracted form mime: %s => %s", mime, format)

				/*if !IsFormatSupported(format) {
					http.Error(w, fmt.Sprintf("Format not supported"), http.StatusUnsupportedMediaType)

					return
				}*/

				options.Format = image.ExtensionToType(format)

				r = r.WithContext(NewOptionsContext(ctx, options))
			} else {
				mime = image.GetImageMimeType(options.Format)
			}

			w.Header().Set("Content-Type", mime)
			w.Header().Add("Vary", "Accept")

			next.ServeHTTP(w, r)
		})
	}
}
