package main

import (
	"cart/w4"
	"math"
)

var ball_spr = [8]byte{
	0b11111111,
	0b11000011,
	0b10000001,
	0b10011001,
	0b10011001,
	0b10000001,
	0b11000011,
	0b11111111,
}

type ball struct {
	rigid
}

func bounce() {
	w4.Tone(262, 5, 50, w4.TONE_PULSE1)
}

func (self *ball) update() {
	counter++
	if counter%4 == 0 {
		old_old_b = old_b
		old_b = b
	}

	self.xspeed = math.Cos(self.angle) * self.speed
	self.yspeed = -math.Sin(self.angle) * self.speed

	if self.x > p1.x && self.x < p1.x+8 && self.y > p1.y && self.y < p1.y+32 {
		self.xspeed = -self.xspeed
		bounce()
	}
	if self.x > p2.x && self.x < p2.x+8 && self.y > p2.y && self.y < p2.y+32 {
		self.xspeed = -self.xspeed
		bounce()
	}

	if self.y <= 0 || self.y >= 160 {
		self.yspeed = -self.yspeed
		bounce()
	}

	if self.x <= 0 || self.x >= 160 {
		self.xspeed = -self.xspeed
		bounce()
	}

	self.angle = math.Atan2(-self.yspeed, self.xspeed)

	self.x += self.xspeed
	self.y += self.yspeed
}

func (self *ball) draw() {
	*w4.DRAW_COLORS = 2
	w4.Blit(&ball_spr[0], int(old_old_b.x)-4, int(old_old_b.y)-4, 8, 8, w4.BLIT_1BPP)

	*w4.DRAW_COLORS = 3
	w4.Blit(&ball_spr[0], int(old_b.x)-4, int(old_b.y)-4, 8, 8, w4.BLIT_1BPP)

	*w4.DRAW_COLORS = 4
	w4.Blit(&ball_spr[0], int(b.x)-4, int(b.y)-4, 8, 8, w4.BLIT_1BPP)
}
