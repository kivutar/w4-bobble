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

// Musical note frequencies used by tracks.
// var TrackNotes = []uint16{
// 	44, 46, 49, 52, 55, 58, 62, 65,
// 	69, 73, 78, 82, 87, 92, 98, 104, 110, 117,
// 	123, 139, 147, 156, 165, 175, 185, 196, 208,
// 	220, 233, 247, 262, 277, 293, 311, 330, 349,
// 	370, 392, 415, 440, 466, 493, 523, 554, 587,
// 	622, 659, 698, 740, 784, 830, 880, 932, 987,
// 	1046}

type Channel struct {
	next       int
	counter    uint
	instrument uint8
	flags      [2]uint8
	tones      [][3]uint16
}

// A sound track is basically one fragment of a music.
type Track struct {
	ticks    uint
	channels []*Channel
}

func toneSub(p ToneParams) {
	w4.Tone(
		p.freq1|(p.freq2<<16),
		(p.attack<<24)|(p.decay<<16)|p.sustain|(p.release<<8),
		p.volume,
		p.channel,
	)
}

func noteFreq(n uint16) uint16 {
	// var freqs = []uint16{
	// 	28160, // A10
	// 	29834,
	// 	31609,
	// 	33488,
	// 	35479,
	// 	37589,
	// 	39824,
	// 	42192,
	// 	44701,
	// 	47359,
	// 	50175,
	// 	53159,
	// }
	//return freqs[n%12] >> (10 - n/12)
	return uint16(math.Pow(2, (float64(n)-57)/12) * 440)
}

func (self *Track) Step() {
	for _, channel := range self.channels {
		// Instrument in use.
		instrument := channel.instrument % 4

		if uint(channel.tones[channel.next][0]) == channel.counter/(self.ticks/2) {
			// Musical tone to be played.
			tone := channel.tones[channel.next]
			note := tone[1]
			wait := tone[2] / 2

			// Play tone...
			toneSub(ToneParams{
				//freq1: uint(TrackNotes[note-18]),
				freq1:   uint(noteFreq(note)),
				freq2:   0,
				attack:  0,
				decay:   0,
				sustain: 0,
				release: uint(wait) * self.ticks / 2,
				volume:  100,
				channel: uint(instrument),
			})

			channel.next++
			channel.next %= len(channel.tones)
		}
		channel.counter++
		channel.counter %= 256 * self.ticks
	}
}
