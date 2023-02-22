package main

import (
	"cart/w4"
	"math"
)

type Channel struct {
	next       int
	counter    uint
	instrument uint8
	volume     uint8
	tones      [][3]uint16
}

// A sound track is basically one fragment of a music.
type Track struct {
	ticks    uint
	channels []*Channel
}

func noteFreq(n uint16) uint {
	return uint(math.Pow(2, (float64(n)-57)/12) * 440)
}

func (self *Track) Step() {
	for _, channel := range self.channels {
		instrument := uint(channel.instrument % 4)

		if uint(channel.tones[channel.next][0]) == channel.counter/(self.ticks/2) {
			// Musical tone to be played.
			tone := channel.tones[channel.next]
			note := tone[1]
			wait := uint(tone[2])

			if instrument == w4.TONE_NOISE {
				w4.Tone(
					noteFreq(note),
					(wait*self.ticks/8)<<8,
					uint(channel.volume),
					instrument,
				)
			} else {
				w4.Tone(
					noteFreq(note),
					(wait*self.ticks/4)<<8,
					uint(channel.volume),
					instrument,
				)
			}

			channel.next++
			channel.next %= len(channel.tones)
		}
		channel.counter++
		channel.counter %= (12 * 32 * self.ticks / 2)
	}
}
