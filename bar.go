package main

import "cart/w4"

var bar_spr = [32]byte{
	0b10000001,
	0b01111110,
	0b01000010,
	0b01011010,
	0b01011010,
	0b01011010,
	0b01011010,
	0b01011010,
	0b01011010,
	0b01011010,
	0b01011010,
	0b01011010,
	0b01011010,
	0b01011010,
	0b01011010,
	0b01011010,
	0b01011010,
	0b01011010,
	0b01011010,
	0b01011010,
	0b01011010,
	0b01011010,
	0b01011010,
	0b01011010,
	0b01011010,
	0b01011010,
	0b01011010,
	0b01011010,
	0b01011010,
	0b01000010,
	0b01111110,
	0b10000001,
}

type bar struct {
	rigid
}

func friction(v float64, pad uint8) float64 {
	if v < 0 && (pad&w4.BUTTON_UP == 0) {
		v += 0.1
		if v > 0 {
			v = 0
		}
	} else if v > 0 && (pad&w4.BUTTON_DOWN == 0) {
		v -= 0.1
		if v < 0 {
			v = 0
		}
	}
	return v
}

func clamp(v, min, max float64) float64 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

func (self *bar) update(pad uint8) {
	if pad&w4.BUTTON_UP != 0 {
		self.yspeed -= 0.1
	}
	if pad&w4.BUTTON_DOWN != 0 {
		self.yspeed += 0.1
	}

	self.yspeed = clamp(self.yspeed, -2, 2)

	self.y += self.yspeed
	self.y = clamp(self.y, 0, 160-self.height)

	self.yspeed = friction(self.yspeed, pad)
}

func (self *bar) draw() {
	*w4.DRAW_COLORS = 2
	w4.Blit(&bar_spr[0], int(self.x), int(self.y), uint(self.width), uint(self.height), w4.BLIT_1BPP)
}
