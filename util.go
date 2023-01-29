package main

import "cart/w4"

func xfriction(v float64, pad uint8) float64 {
	if v < 0 && (pad&w4.BUTTON_LEFT == 0) {
		v += 0.1
		if v > 0 {
			v = 0
		}
	} else if v > 0 && (pad&w4.BUTTON_RIGHT == 0) {
		v -= 0.1
		if v < 0 {
			v = 0
		}
	}
	return v
}

func yfriction(v float64, pad uint8) float64 {
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
