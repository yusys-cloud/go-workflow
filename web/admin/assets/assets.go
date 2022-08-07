// Author: yangzq80@gmail.com
// Date: 2022/8/6
//
package assets

import "embed"

var (
	//go:embed web/dist/*
	Static embed.FS
)
