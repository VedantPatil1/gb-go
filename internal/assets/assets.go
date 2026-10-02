// Package assets provides embedded application assets.
package assets

import "embed"

//go:embed static
var Files embed.FS
