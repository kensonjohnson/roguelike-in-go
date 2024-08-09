//go:build dev

package main

import "github.com/hajimehoshi/ebiten/v2"

func init() {
	ebiten.SetVsyncEnabled(false)
}
