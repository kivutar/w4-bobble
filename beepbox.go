package main

import (
	"cart/w4"
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
var TrackNotes = [37]uint16{
	130, 140, 150, 160, 170, 180, 190, 200, 210,
	220, 230, 250, 260, 280, 290, 310, 330, 350,
	370, 390, 410, 440, 460, 490, 520, 550, 600,
	620, 660, 700, 750, 780, 840, 880, 940, 980,
	1000}

type Channel struct {
	next       int
	counter    uint
	instrument uint8
	flags      [2]uint8
	tones      [][3]byte
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

func (self *Track) Step() {
	for _, channel := range self.channels {
		// Instrument in use.
		instrument := channel.instrument % 4

		if uint(channel.tones[channel.next][0]) == channel.counter/self.ticks {
			// Musical tone to be played.
			tone := channel.tones[channel.next]
			note := tone[1]
			wait := tone[2]

			// Play tone...
			if int(note) < len(TrackNotes) {
				toneSub(ToneParams{
					freq1:   uint(TrackNotes[note]),
					freq2:   0,
					attack:  0,
					decay:   0,
					sustain: 0,
					release: uint(wait) * self.ticks,
					volume:  100,
					channel: uint(instrument),
				})
			}

			channel.next++
			channel.next %= len(channel.tones)
		}
		channel.counter++
		channel.counter %= 256 * self.ticks
	}
}
