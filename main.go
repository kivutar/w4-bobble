package main

import (
	"cart/w4"
)

var level = [][]byte{
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 1, 1},
	{1, 0, 0, 0, 0, 0, 0, 1, 1, 1},
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
}

type rigid struct {
	x      float64
	y      float64
	width  float64
	height float64
	speed  float64
	angle  float64
	xspeed float64
	yspeed float64
	xaccel float64
	yaccel float64
}

var tur1 = turnip{rigid: rigid{x: 32, y: 32, width: 12, height: 16}}
var tur2 = turnip{rigid: rigid{x: 16, y: 16, width: 12, height: 16}}

//go:export start
func start() {
	w4.PALETTE[0] = 0xfff6d3
	w4.PALETTE[1] = 0xf9a875
	w4.PALETTE[2] = 0xeb6b6f
	w4.PALETTE[3] = 0x7c3f58
}

var block_spr = [64]byte{0x55, 0x55, 0x55, 0x55, 0x7f, 0xff, 0xff, 0xfd, 0x6f, 0xff, 0xff, 0xf9, 0x6b, 0xff, 0xff, 0xe9, 0x6a, 0xff, 0xff, 0xa9, 0x6a, 0xbf, 0xfe, 0xa9, 0x6a, 0xaf, 0xfa, 0xa9, 0x6a, 0xab, 0xea, 0xa9, 0x6a, 0xa9, 0x6a, 0xa9, 0x6a, 0xa5, 0x5a, 0xa9, 0x6a, 0x95, 0x56, 0xa9, 0x6a, 0x55, 0x55, 0xa9, 0x69, 0x55, 0x55, 0x69, 0x65, 0x55, 0x55, 0x59, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55}

//go:export update
func update() {
	tur1.update(*w4.GAMEPAD1)
	tur2.update(*w4.GAMEPAD2)

	tur1.draw()
	tur2.draw()

	*w4.DRAW_COLORS = 0x241
	for y, line := range level {
		for x, tile := range line {
			if tile == 1 {
				w4.Blit(&block_spr[0], x*16, y*16, 16, 16, w4.BLIT_2BPP)
			}
		}
	}

	song.Step()
}
