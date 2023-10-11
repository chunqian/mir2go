// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package common

import "github.com/ying32/govcl/vcl/types"

func RGB(r, g, b uint32) types.TColor {
	return types.TColor(r | (g << 8) | (b << 16))
}
