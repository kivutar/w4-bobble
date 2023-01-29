package main

import (
	"cart/w4"
	"math"
	"math/rand"
)

var ball_spr = [8]byte{
	0b11100111,
	0b10000001,
	0b10000001,
	0b00000000,
	0b00000000,
	0b10000001,
	0b10000001,
	0b11100111,
}

type ball struct {
	rigid
}

func bounce() {
	w4.Tone(262, 5, 50, w4.TONE_PULSE1)
}

func (self *ball) update() {
	// counter++
	// if counter%4 == 0 {
	// 	old_old_b = old_b
	// 	old_b = b
	// }

	self.xspeed = math.Cos(self.angle) * self.speed
	self.yspeed = -math.Sin(self.angle) * self.speed

	if self.x+self.width > p1.x && self.x < p1.x+p1.width &&
		self.y+self.height > p1.y && self.y < p1.y+p1.height {
		self.xspeed = -self.xspeed
		bounce()
	}
	if self.x+self.width > p2.x && self.x < p2.x+p1.width &&
		self.y+self.height > p2.y && self.y < p2.y+p2.height {
		self.xspeed = -self.xspeed
		bounce()
	}

	if self.y <= 0 || self.y+self.height >= 160 {
		self.yspeed = -self.yspeed
		bounce()
	}

	if self.x <= 0 || self.x+self.width >= 160 {
		if self.x <= 0 {
			p2score++
		}
		if self.x+self.width >= 160 {
			p1score++
		}
		self.x = 80 - 4
		self.y = 80 - 4
		self.angle = rand.Float64() * 2 * math.Pi
		w4.Tone(200, 10, 50, w4.TONE_PULSE1)
		return
	}

	self.angle = math.Atan2(-self.yspeed, self.xspeed)

	self.x += self.xspeed
	self.y += self.yspeed
}

func (self *ball) draw() {
	// *w4.DRAW_COLORS = 2
	// w4.Blit(&ball_spr[0], int(old_old_b.x), int(old_old_b.y), 8, 8, w4.BLIT_1BPP)

	// *w4.DRAW_COLORS = 3
	// w4.Blit(&ball_spr[0], int(old_b.x), int(old_b.y), 8, 8, w4.BLIT_1BPP)

	*w4.DRAW_COLORS = 3
	w4.Blit(&ball_spr[0], int(self.x), int(self.y), 8, 8, w4.BLIT_1BPP)
}
