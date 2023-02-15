package main

import (
	"cart/w4"
	"math"
)

type ToneParams struct {
	freq1   uint // First frequency.
	freq2   uint // Second frequency.
	attack  uint // Attack (aggressive start).
	decay   uint // Decay (smooth fading).
	sustain uint // Sustain time.
	release uint // Release time.
	volume  uint // Volume.
	channel uint // Audio channel. Each channel, except `0` and `1`, are different.
	mode    uint // Audio mode. Only works for channels `0` and `1`.
}

type Channel struct {
	next       int
	counter    uint
	instrument uint8
	tones      [][3]uint16
}

// A sound track is basically one fragment of a music.
type Track struct {
	ticks    uint
	channels []*Channel
}

func noteFreq(n uint16) uint16 {
	return uint16(math.Pow(2, (float64(n)-57)/12) * 440)
}

func (self *Track) Step() {
	for _, channel := range self.channels {
		// Instrument in use.
		instrument := channel.instrument % 4

		if uint(channel.tones[channel.next][0]) == channel.counter/(self.ticks) {
			// Musical tone to be played.
			tone := channel.tones[channel.next]
			note := tone[1]
			wait := tone[2]

			// Play tone...
			w4.Tone(
				uint(noteFreq(note)),
				(uint(wait)*self.ticks)<<8,
				100,
				uint(instrument),
			)

			channel.next++
			channel.next %= len(channel.tones)
		}
		channel.counter++
		channel.counter %= 256 * self.ticks
	}
}
