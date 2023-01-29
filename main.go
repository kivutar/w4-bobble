package main

import (
	"cart/w4"
)

type rigid struct {
	x      float64
	y      float64
	width  float64
	height float64
	speed  float64
	angle  float64
	xspeed float64
	yspeed float64
}

var tur = turnip{rigid: rigid{x: 16, y: 16, width: 16, height: 16}}

//go:export start
func start() {
	w4.PALETTE[0] = 0xfff6d3
	w4.PALETTE[1] = 0xf9a875
	w4.PALETTE[2] = 0xeb6b6f
	w4.PALETTE[3] = 0x7c3f58
}

//go:export update
func update() {
	tur.update(*w4.GAMEPAD1)

	tur.draw()
}
