package main

import "cart/w4"

var turnip_spr0 = [64]byte{0x15, 0x55, 0x55, 0x54, 0x1f, 0xff, 0xfe, 0xa4, 0x1d, 0xff, 0xde, 0xa4, 0x1d, 0xff, 0xde, 0xa4, 0x1d, 0xff, 0xde, 0xa4, 0x1d, 0xff, 0xde, 0xa4, 0x1f, 0xff, 0xfe, 0xa4, 0x55, 0x55, 0x55, 0x54, 0x75, 0xaa, 0xaa, 0xa4, 0x6a, 0xaa, 0xaa, 0x69, 0x15, 0xaa, 0xaa, 0x99, 0x07, 0xff, 0xff, 0x5d, 0x07, 0xff, 0xfd, 0x04, 0x05, 0x55, 0x57, 0x40, 0x00, 0x00, 0x1f, 0xd0, 0x00, 0x00, 0x05, 0x50}
var turnip_spr1 = [64]byte{0x00, 0x00, 0x00, 0x00, 0x15, 0x55, 0x55, 0x54, 0x1f, 0xff, 0xfe, 0xa4, 0x1d, 0xff, 0xde, 0xa4, 0x1d, 0xff, 0xde, 0xa4, 0x1d, 0xff, 0xde, 0xa4, 0x1d, 0xff, 0xde, 0xa4, 0x1f, 0xff, 0xfe, 0xa4, 0x15, 0x55, 0x55, 0x54, 0x01, 0xaa, 0xaa, 0x90, 0x05, 0xaa, 0xa6, 0x90, 0x05, 0xaa, 0x9e, 0x90, 0x01, 0xff, 0xf5, 0x40, 0x01, 0xfd, 0xff, 0x40, 0x00, 0x7d, 0xfd, 0x00, 0x00, 0x15, 0x54, 0x00}
var turnip_spr2 = [64]byte{0x15, 0x55, 0x55, 0x54, 0x1f, 0xff, 0xfe, 0xa4, 0x1d, 0xff, 0xde, 0xa4, 0x1d, 0xff, 0xde, 0xa4, 0x1d, 0xff, 0xde, 0xa4, 0x1d, 0xff, 0xde, 0xa4, 0x1f, 0xff, 0xfe, 0xa4, 0x15, 0x55, 0x55, 0x54, 0x01, 0x9e, 0xaa, 0x40, 0x01, 0xa5, 0x56, 0x40, 0x55, 0xaa, 0xaa, 0x40, 0x7f, 0xff, 0xff, 0x40, 0x7f, 0xff, 0xff, 0x40, 0x55, 0x55, 0x7f, 0xd0, 0x00, 0x00, 0x1f, 0xf4, 0x00, 0x00, 0x05, 0x54}

var turnip_run_anim = Anim{
	frames: [][64]byte{
		turnip_spr0,
		turnip_spr1,
		turnip_spr2,
		turnip_spr1,
	},
}

var turnip_stand_anim = Anim{
	frames: [][64]byte{
		turnip_spr1,
	},
}

var turnip_jump_anim = Anim{
	frames: [][64]byte{
		turnip_spr2,
	},
}

type turnip struct {
	rigid
	anim *Anim
	flip uint
}

func (self *turnip) update(pad uint8) {
	if pad&w4.BUTTON_LEFT != 0 {
		self.xaccel = -0.1
	} else if pad&w4.BUTTON_RIGHT != 0 {
		self.xaccel = 0.1
	} else {
		self.xaccel = 0
	}

	self.xspeed += self.xaccel
	self.yspeed += self.yaccel

	if level[int((self.y+self.width)/16)][int(self.x/16)] != 0x1 && level[int((self.y+16)/16)][int((self.x+self.width-1)/16)] != 0x1 {
		self.yaccel = 0.2
	} else {
		self.y = float64(int(self.y/16)) * 16
		self.yspeed = 0
		self.yaccel = 0
		if pad&w4.BUTTON_1 != 0 {
			self.yspeed = -4
		}
	}

	self.xspeed = clamp(self.xspeed, -2, 2)
	self.yspeed = clamp(self.yspeed, -4, 4)
	self.x += self.xspeed
	self.y += self.yspeed
	self.xspeed = xfriction(self.xspeed, pad)

	if level[int((self.y)/16)][int((self.x+self.width)/16)] == 0x1 && self.xspeed > 0 {
		self.x = float64(int(self.x/16))*16 + 16 - self.width
		self.xspeed = 0
		self.xaccel = 0
	}

	if level[int((self.y)/16)][int((self.x)/16)] == 0x1 && self.xspeed < 0 {
		self.x = float64(int(self.x/16))*16 + 16
		self.xspeed = 0
		self.xaccel = 0
	}

	if self.xspeed > 0 {
		self.flip = w4.BLIT_FLIP_X
	} else if self.xspeed < 0 {
		self.flip = 0
	}

	if self.xspeed == 0 {
		self.anim = &turnip_stand_anim
	} else {
		self.anim = &turnip_run_anim
	}

	self.anim.counter++
}

func (self *turnip) draw() {
	*w4.DRAW_COLORS = 0x241
	w4.Blit(&self.anim.frames[(self.anim.counter/8)%len(self.anim.frames)][0], int(self.x)-2, int(self.y), 16, 16, w4.BLIT_2BPP|self.flip)
}
