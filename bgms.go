package main

// Soundtrack: simplesong
var simplesong = Track{
	ticks:     5,
	sequences: 32,
	channels: []*Channel{
		{
			instrument: 0,
			volume:     100,
			tones: [][3]uint16{
				{0x00, 0x42, 0x05},
				{0x02, 0x49, 0x05},
				{0x04, 0x4e, 0x05},
				{0x06, 0x50, 0x05},
				{0x08, 0x49, 0x05},
				{0x0a, 0x4e, 0x05},
				{0x0c, 0x50, 0x05},
				{0x0e, 0x49, 0x05},
				{0x10, 0x53, 0x05},
				{0x12, 0x49, 0x05},
				{0x14, 0x53, 0x05},
				{0x16, 0x52, 0x05},
				{0x18, 0x49, 0x05},
				{0x1a, 0x52, 0x05},
				{0x1c, 0x50, 0x05},
				{0x1e, 0x4e, 0x05},
				{0x20, 0x42, 0x05},
				{0x22, 0x49, 0x05},
				{0x24, 0x4e, 0x05},
				{0x26, 0x50, 0x05},
				{0x28, 0x49, 0x05},
				{0x2a, 0x4e, 0x05},
				{0x2c, 0x50, 0x05},
				{0x2e, 0x49, 0x05},
				{0x30, 0x53, 0x05},
				{0x32, 0x49, 0x05},
				{0x34, 0x53, 0x05},
				{0x36, 0x52, 0x05},
				{0x38, 0x49, 0x05},
				{0x3a, 0x52, 0x05},
				{0x3c, 0x50, 0x05},
				{0x3e, 0x4e, 0x05},
				{0x40, 0x42, 0x05},
				{0x42, 0x49, 0x05},
				{0x44, 0x4e, 0x05},
				{0x46, 0x50, 0x05},
				{0x48, 0x49, 0x05},
				{0x4a, 0x4e, 0x05},
				{0x4c, 0x50, 0x05},
				{0x4e, 0x49, 0x05},
				{0x50, 0x53, 0x05},
				{0x52, 0x49, 0x05},
				{0x54, 0x53, 0x05},
				{0x56, 0x52, 0x05},
				{0x58, 0x49, 0x05},
				{0x5a, 0x52, 0x05},
				{0x5c, 0x50, 0x05},
				{0x5e, 0x4e, 0x05},
				{0x60, 0x42, 0x05},
				{0x62, 0x49, 0x05},
				{0x64, 0x4e, 0x05},
				{0x66, 0x50, 0x05},
				{0x68, 0x49, 0x05},
				{0x6a, 0x4e, 0x05},
				{0x6c, 0x50, 0x05},
				{0x6e, 0x49, 0x05},
				{0x70, 0x53, 0x05},
				{0x72, 0x49, 0x05},
				{0x74, 0x53, 0x05},
				{0x76, 0x52, 0x05},
				{0x78, 0x49, 0x05},
				{0x7a, 0x52, 0x05},
				{0x7c, 0x50, 0x05},
				{0x7e, 0x4e, 0x05},
				{0x80, 0x42, 0x05},
				{0x82, 0x49, 0x05},
				{0x84, 0x4e, 0x05},
				{0x86, 0x50, 0x05},
				{0x88, 0x49, 0x05},
				{0x8a, 0x4e, 0x05},
				{0x8c, 0x50, 0x05},
				{0x8e, 0x49, 0x05},
				{0x90, 0x53, 0x05},
				{0x92, 0x49, 0x05},
				{0x94, 0x53, 0x05},
				{0x96, 0x52, 0x05},
				{0x98, 0x49, 0x05},
				{0x9a, 0x52, 0x05},
				{0x9c, 0x50, 0x05},
				{0x9e, 0x4e, 0x05},
				{0xa0, 0x42, 0x05},
				{0xa2, 0x49, 0x05},
				{0xa4, 0x4e, 0x05},
				{0xa6, 0x50, 0x05},
				{0xa8, 0x49, 0x05},
				{0xaa, 0x4e, 0x05},
				{0xac, 0x50, 0x05},
				{0xae, 0x49, 0x05},
				{0xb0, 0x53, 0x05},
				{0xb2, 0x49, 0x05},
				{0xb4, 0x53, 0x05},
				{0xb6, 0x52, 0x05},
				{0xb8, 0x49, 0x05},
				{0xba, 0x52, 0x05},
				{0xbc, 0x50, 0x05},
				{0xbe, 0x4e, 0x05},
				{0xc0, 0x42, 0x05},
				{0xc2, 0x49, 0x05},
				{0xc4, 0x4e, 0x05},
				{0xc6, 0x50, 0x05},
				{0xc8, 0x49, 0x05},
				{0xca, 0x4e, 0x05},
				{0xcc, 0x50, 0x05},
				{0xce, 0x49, 0x05},
				{0xd0, 0x53, 0x05},
				{0xd2, 0x49, 0x05},
				{0xd4, 0x53, 0x05},
				{0xd6, 0x52, 0x05},
				{0xd8, 0x49, 0x05},
				{0xda, 0x52, 0x05},
				{0xdc, 0x50, 0x05},
				{0xde, 0x4e, 0x05},
				{0xe0, 0x42, 0x05},
				{0xe2, 0x49, 0x05},
				{0xe4, 0x4e, 0x05},
				{0xe6, 0x50, 0x05},
				{0xe8, 0x49, 0x05},
				{0xea, 0x4e, 0x05},
				{0xec, 0x50, 0x05},
				{0xee, 0x49, 0x05},
				{0xf0, 0x53, 0x05},
				{0xf2, 0x49, 0x05},
				{0xf4, 0x53, 0x05},
				{0xf6, 0x52, 0x05},
				{0xf8, 0x49, 0x05},
				{0xfa, 0x52, 0x05},
				{0xfc, 0x50, 0x05},
				{0xfe, 0x4e, 0x05},
				{0x100, 0x42, 0x05},
				{0x102, 0x49, 0x05},
				{0x104, 0x4e, 0x05},
				{0x106, 0x50, 0x05},
				{0x108, 0x49, 0x05},
				{0x10a, 0x4e, 0x05},
				{0x10c, 0x50, 0x05},
				{0x10e, 0x49, 0x05},
				{0x110, 0x53, 0x05},
				{0x112, 0x49, 0x05},
				{0x114, 0x53, 0x05},
				{0x116, 0x52, 0x05},
				{0x118, 0x49, 0x05},
				{0x11a, 0x52, 0x05},
				{0x11c, 0x50, 0x05},
				{0x11e, 0x4e, 0x05},
				{0x120, 0x42, 0x05},
				{0x122, 0x49, 0x05},
				{0x124, 0x4e, 0x05},
				{0x126, 0x50, 0x05},
				{0x128, 0x49, 0x05},
				{0x12a, 0x4e, 0x05},
				{0x12c, 0x50, 0x05},
				{0x12e, 0x49, 0x05},
				{0x130, 0x53, 0x05},
				{0x132, 0x49, 0x05},
				{0x134, 0x53, 0x05},
				{0x136, 0x52, 0x05},
				{0x138, 0x49, 0x05},
				{0x13a, 0x52, 0x05},
				{0x13c, 0x50, 0x05},
				{0x13e, 0x4e, 0x05},
				{0x140, 0x42, 0x05},
				{0x142, 0x49, 0x05},
				{0x144, 0x4e, 0x05},
				{0x146, 0x50, 0x05},
				{0x148, 0x49, 0x05},
				{0x14a, 0x4e, 0x05},
				{0x14c, 0x50, 0x05},
				{0x14e, 0x49, 0x05},
				{0x150, 0x53, 0x05},
				{0x152, 0x49, 0x05},
				{0x154, 0x53, 0x05},
				{0x156, 0x52, 0x05},
				{0x158, 0x49, 0x05},
				{0x15a, 0x52, 0x05},
				{0x15c, 0x50, 0x05},
				{0x15e, 0x4e, 0x05},
				{0x160, 0x42, 0x05},
				{0x162, 0x49, 0x05},
				{0x164, 0x4e, 0x05},
				{0x166, 0x50, 0x05},
				{0x168, 0x49, 0x05},
				{0x16a, 0x4e, 0x05},
				{0x16c, 0x50, 0x05},
				{0x16e, 0x49, 0x05},
				{0x170, 0x53, 0x05},
				{0x172, 0x49, 0x05},
				{0x174, 0x53, 0x05},
				{0x176, 0x52, 0x05},
				{0x178, 0x49, 0x05},
				{0x17a, 0x52, 0x05},
				{0x17c, 0x50, 0x05},
				{0x17e, 0x4e, 0x05},
				{0x180, 0x42, 0x05},
				{0x182, 0x49, 0x05},
				{0x184, 0x4e, 0x05},
				{0x186, 0x50, 0x05},
				{0x188, 0x49, 0x05},
				{0x18a, 0x4e, 0x05},
				{0x18c, 0x50, 0x05},
				{0x18e, 0x49, 0x05},
				{0x190, 0x53, 0x05},
				{0x192, 0x49, 0x05},
				{0x194, 0x53, 0x05},
				{0x196, 0x52, 0x05},
				{0x198, 0x49, 0x05},
				{0x19a, 0x52, 0x05},
				{0x19c, 0x50, 0x05},
				{0x19e, 0x4e, 0x05},
				{0x1a0, 0x42, 0x05},
				{0x1a2, 0x49, 0x05},
				{0x1a4, 0x4e, 0x05},
				{0x1a6, 0x50, 0x05},
				{0x1a8, 0x49, 0x05},
				{0x1aa, 0x4e, 0x05},
				{0x1ac, 0x50, 0x05},
				{0x1ae, 0x49, 0x05},
				{0x1b0, 0x53, 0x05},
				{0x1b2, 0x49, 0x05},
				{0x1b4, 0x53, 0x05},
				{0x1b6, 0x52, 0x05},
				{0x1b8, 0x49, 0x05},
				{0x1ba, 0x52, 0x05},
				{0x1bc, 0x50, 0x05},
				{0x1be, 0x4e, 0x05},
				{0x1c0, 0x42, 0x05},
				{0x1c2, 0x49, 0x05},
				{0x1c4, 0x4e, 0x05},
				{0x1c6, 0x50, 0x05},
				{0x1c8, 0x49, 0x05},
				{0x1ca, 0x4e, 0x05},
				{0x1cc, 0x53, 0x05},
				{0x1ce, 0x52, 0x05},
				{0x1d0, 0x4e, 0x05},
				{0x1d2, 0x49, 0x05},
				{0x1d4, 0x47, 0x05},
				{0x1d6, 0x46, 0x05},
				{0x1d8, 0x42, 0x05},
				{0x1da, 0x3d, 0x05},
				{0x1dc, 0x3b, 0x05},
				{0x1de, 0x3a, 0x05},
				{0x1f2, 0x44, 0x05},
				{0x1f4, 0x45, 0x05},
				{0x1f6, 0x4c, 0x05},
				{0x1fa, 0x50, 0x05},
				{0x1fc, 0x51, 0x05},
				{0x1fe, 0x4c, 0x05},
				{0x260, 0x42, 0x05},
				{0x262, 0x49, 0x05},
				{0x264, 0x4e, 0x05},
				{0x266, 0x50, 0x05},
				{0x268, 0x49, 0x05},
				{0x26a, 0x4e, 0x05},
				{0x26c, 0x50, 0x05},
				{0x26e, 0x49, 0x05},
				{0x270, 0x53, 0x05},
				{0x272, 0x49, 0x05},
				{0x274, 0x53, 0x05},
				{0x276, 0x52, 0x05},
				{0x278, 0x49, 0x05},
				{0x27a, 0x52, 0x05},
				{0x27c, 0x50, 0x05},
				{0x27e, 0x4e, 0x05},
			},
		},
		{
			instrument: 1,
			volume:     100,
			tones: [][3]uint16{
				{0x40, 0x31, 0x5f},
				{0x54, 0x36, 0x0f},
				{0x58, 0x38, 0x0f},
				{0x5c, 0x3b, 0x0f},
				{0x60, 0x3b, 0x0f},
				{0x68, 0x3b, 0x0a},
				{0x6c, 0x3a, 0x32},
				{0x78, 0x38, 0x0f},
				{0x7c, 0x36, 0x0f},
				{0x80, 0x3d, 0x2d},
				{0x8c, 0x36, 0x1e},
				{0x94, 0x42, 0x37},
				{0xa0, 0x42, 0x0f},
				{0xb0, 0x42, 0x0f},
				{0xb4, 0x44, 0x0f},
				{0xb8, 0x42, 0x0f},
				{0xc0, 0x49, 0x96},
				{0xe8, 0x3a, 0x05},
				{0xea, 0x3b, 0x05},
				{0xec, 0x3d, 0x05},
				{0xee, 0x3d, 0x14},
				{0xf4, 0x42, 0x0f},
				{0xf8, 0x41, 0x0f},
				{0xfc, 0x42, 0x0f},
				{0x100, 0x44, 0x0f},
				{0x104, 0x46, 0x05},
				{0x106, 0x42, 0x41},
				{0x11c, 0x41, 0x05},
				{0x11e, 0x42, 0x05},
				{0x120, 0x49, 0x19},
				{0x126, 0x42, 0x05},
				{0x128, 0x42, 0x41},
				{0x13c, 0x42, 0x05},
				{0x13e, 0x42, 0x05},
				{0x140, 0x42, 0x2d},
				{0x14a, 0x44, 0x05},
				{0x14c, 0x42, 0x05},
				{0x14e, 0x41, 0x05},
				{0x150, 0x41, 0x4b},
				{0x160, 0x3a, 0x23},
				{0x168, 0x3a, 0x05},
				{0x16a, 0x38, 0x05},
				{0x16c, 0x36, 0x05},
				{0x16e, 0x3d, 0x2d},
				{0x178, 0x42, 0x0f},
				{0x17c, 0x44, 0x0f},
				{0x180, 0x47, 0x19},
				{0x186, 0x46, 0x05},
				{0x188, 0x46, 0x23},
				{0x198, 0x49, 0x0f},
				{0x19c, 0x47, 0x0f},
				{0x1a0, 0x46, 0x05},
				{0x1a2, 0x47, 0x05},
				{0x1a4, 0x46, 0x05},
				{0x1a6, 0x42, 0x0f},
				{0x1aa, 0x3d, 0x23},
				{0x1b2, 0x44, 0x05},
				{0x1b4, 0x42, 0x05},
				{0x1b6, 0x41, 0x05},
				{0x1b8, 0x3f, 0x05},
				{0x1ba, 0x41, 0x14},
				{0x1c0, 0x42, 0x0f},
				{0x1c6, 0x42, 0x0f},
				{0x1cc, 0x42, 0x14},
				{0x1e0, 0x42, 0x05},
				{0x1e2, 0x42, 0x05},
				{0x1e8, 0x44, 0x05},
				{0x1ea, 0x44, 0x05},
				{0x1ee, 0x40, 0x0a},
				{0x1f6, 0x3d, 0x05},
				{0x1f8, 0x40, 0x05},
				{0x1fa, 0x3b, 0x05},
				{0x1fc, 0x3d, 0x05},
				{0x1fe, 0x3b, 0x05},
				{0x1ff, 0x39, 0x05},
				{0x200, 0x3b, 0x05},
				{0x202, 0x3b, 0x05},
				{0x204, 0x39, 0x05},
				{0x206, 0x3b, 0x05},
				{0x20a, 0x39, 0x05},
				{0x20e, 0x3d, 0x41},
				{0x220, 0x42, 0x05},
				{0x222, 0x42, 0x05},
				{0x228, 0x44, 0x05},
				{0x22a, 0x44, 0x05},
				{0x22e, 0x45, 0x0f},
				{0x232, 0x40, 0x0a},
				{0x238, 0x40, 0x05},
				{0x23c, 0x45, 0x05},
				{0x23e, 0x47, 0x05},
				{0x240, 0x47, 0x41},
				{0x24e, 0x45, 0x05},
				{0x250, 0x49, 0x23},
				{0x258, 0x49, 0x05},
				{0x25a, 0x49, 0x05},
				{0x25c, 0x49, 0x05},
				{0x25e, 0x49, 0x05},
				{0x260, 0x49, 0x4b},
			},
		},
		{
			instrument: 2,
			volume:     100,
			tones: [][3]uint16{
				{0x40, 0x36, 0x19},
				{0x46, 0x3d, 0x19},
				{0x4c, 0x42, 0x37},
				{0x60, 0x34, 0x19},
				{0x66, 0x38, 0x19},
				{0x6c, 0x40, 0x37},
				{0x80, 0x33, 0x19},
				{0x86, 0x3a, 0x19},
				{0x8c, 0x3f, 0x4b},
				{0xa0, 0x32, 0x19},
				{0xa6, 0x39, 0x19},
				{0xac, 0x3e, 0x2d},
				{0xbe, 0x2a, 0x05},
				{0xc0, 0x36, 0x05},
				{0xc2, 0x36, 0x05},
				{0xc4, 0x36, 0x05},
				{0xc6, 0x36, 0x05},
				{0xc8, 0x36, 0x05},
				{0xca, 0x36, 0x05},
				{0xcc, 0x36, 0x05},
				{0xce, 0x36, 0x05},
				{0xd0, 0x36, 0x05},
				{0xd2, 0x36, 0x05},
				{0xd4, 0x36, 0x05},
				{0xd6, 0x36, 0x05},
				{0xd8, 0x36, 0x05},
				{0xda, 0x36, 0x05},
				{0xdc, 0x36, 0x05},
				{0xde, 0x36, 0x05},
				{0xe0, 0x36, 0x05},
				{0xe2, 0x36, 0x05},
				{0xe4, 0x36, 0x05},
				{0xe6, 0x36, 0x05},
				{0xe8, 0x36, 0x05},
				{0xea, 0x36, 0x05},
				{0xec, 0x36, 0x05},
				{0xee, 0x36, 0x05},
				{0xf0, 0x35, 0x05},
				{0xf2, 0x35, 0x05},
				{0xf4, 0x35, 0x05},
				{0xf6, 0x35, 0x05},
				{0xf8, 0x35, 0x05},
				{0xfa, 0x35, 0x05},
				{0xfc, 0x35, 0x05},
				{0xfe, 0x35, 0x05},
				{0x100, 0x33, 0x05},
				{0x102, 0x33, 0x05},
				{0x104, 0x33, 0x05},
				{0x106, 0x33, 0x05},
				{0x108, 0x33, 0x05},
				{0x10a, 0x33, 0x05},
				{0x10c, 0x33, 0x05},
				{0x10e, 0x33, 0x05},
				{0x110, 0x31, 0x05},
				{0x112, 0x31, 0x05},
				{0x114, 0x31, 0x05},
				{0x116, 0x31, 0x05},
				{0x118, 0x31, 0x05},
				{0x11a, 0x31, 0x05},
				{0x11c, 0x31, 0x05},
				{0x11e, 0x31, 0x05},
				{0x120, 0x2f, 0x05},
				{0x122, 0x2f, 0x05},
				{0x124, 0x2f, 0x05},
				{0x126, 0x2f, 0x05},
				{0x128, 0x2f, 0x05},
				{0x12a, 0x2f, 0x05},
				{0x12c, 0x2f, 0x05},
				{0x12e, 0x2f, 0x05},
				{0x130, 0x30, 0x05},
				{0x132, 0x30, 0x05},
				{0x134, 0x30, 0x05},
				{0x136, 0x30, 0x05},
				{0x138, 0x30, 0x05},
				{0x13a, 0x30, 0x05},
				{0x13c, 0x30, 0x05},
				{0x13e, 0x30, 0x05},
				{0x140, 0x31, 0x05},
				{0x142, 0x31, 0x05},
				{0x144, 0x31, 0x05},
				{0x146, 0x31, 0x05},
				{0x148, 0x31, 0x05},
				{0x14a, 0x31, 0x05},
				{0x14c, 0x31, 0x05},
				{0x14e, 0x31, 0x05},
				{0x150, 0x31, 0x05},
				{0x152, 0x31, 0x05},
				{0x154, 0x31, 0x05},
				{0x156, 0x31, 0x05},
				{0x158, 0x31, 0x05},
				{0x15a, 0x31, 0x05},
				{0x15c, 0x31, 0x05},
				{0x15e, 0x31, 0x05},
				{0x160, 0x36, 0x05},
				{0x162, 0x36, 0x05},
				{0x164, 0x36, 0x05},
				{0x166, 0x36, 0x05},
				{0x168, 0x36, 0x05},
				{0x16a, 0x36, 0x05},
				{0x16c, 0x36, 0x05},
				{0x16e, 0x36, 0x05},
				{0x170, 0x35, 0x05},
				{0x172, 0x35, 0x05},
				{0x174, 0x35, 0x05},
				{0x176, 0x35, 0x05},
				{0x178, 0x35, 0x05},
				{0x17a, 0x35, 0x05},
				{0x17c, 0x35, 0x05},
				{0x17e, 0x35, 0x05},
				{0x180, 0x33, 0x05},
				{0x182, 0x33, 0x05},
				{0x184, 0x33, 0x05},
				{0x186, 0x33, 0x05},
				{0x188, 0x33, 0x05},
				{0x18a, 0x33, 0x05},
				{0x18c, 0x33, 0x05},
				{0x18e, 0x33, 0x05},
				{0x190, 0x31, 0x05},
				{0x192, 0x31, 0x05},
				{0x194, 0x31, 0x05},
				{0x196, 0x31, 0x05},
				{0x198, 0x31, 0x05},
				{0x19a, 0x31, 0x05},
				{0x19c, 0x31, 0x05},
				{0x19e, 0x31, 0x05},
				{0x1a0, 0x2f, 0x05},
				{0x1a2, 0x2f, 0x05},
				{0x1a4, 0x2f, 0x05},
				{0x1a6, 0x2f, 0x05},
				{0x1a8, 0x2f, 0x05},
				{0x1aa, 0x2f, 0x05},
				{0x1ac, 0x2f, 0x05},
				{0x1ae, 0x2f, 0x05},
				{0x1b0, 0x2e, 0x05},
				{0x1b2, 0x2e, 0x05},
				{0x1b4, 0x2e, 0x05},
				{0x1b6, 0x2e, 0x05},
				{0x1b8, 0x2f, 0x05},
				{0x1ba, 0x2f, 0x05},
				{0x1bc, 0x31, 0x05},
				{0x1be, 0x31, 0x05},
				{0x1c0, 0x32, 0x0f},
				{0x1c6, 0x34, 0x0f},
				{0x1cc, 0x36, 0x14},
				{0x1e0, 0x2f, 0x05},
				{0x1e2, 0x2f, 0x05},
				{0x1e8, 0x34, 0x05},
				{0x1ea, 0x34, 0x05},
				{0x1ee, 0x2d, 0x0a},
				{0x1f2, 0x2d, 0x05},
				{0x1f4, 0x2d, 0x05},
				{0x1f6, 0x2d, 0x05},
				{0x1f8, 0x2d, 0x05},
				{0x1fa, 0x2d, 0x05},
				{0x1fc, 0x2d, 0x05},
				{0x1fe, 0x2d, 0x05},
				{0x200, 0x2f, 0x05},
				{0x202, 0x2f, 0x05},
				{0x204, 0x2f, 0x05},
				{0x206, 0x2f, 0x05},
				{0x208, 0x2f, 0x05},
				{0x20a, 0x2f, 0x05},
				{0x20c, 0x2f, 0x05},
				{0x20e, 0x2f, 0x05},
				{0x210, 0x31, 0x05},
				{0x212, 0x31, 0x05},
				{0x214, 0x31, 0x05},
				{0x216, 0x31, 0x05},
				{0x218, 0x31, 0x05},
				{0x21a, 0x31, 0x05},
				{0x21c, 0x31, 0x05},
				{0x21e, 0x31, 0x05},
				{0x220, 0x2f, 0x05},
				{0x222, 0x2f, 0x05},
				{0x228, 0x34, 0x05},
				{0x22a, 0x34, 0x05},
				{0x22e, 0x2d, 0x0a},
				{0x232, 0x2d, 0x05},
				{0x234, 0x2d, 0x05},
				{0x236, 0x2d, 0x05},
				{0x238, 0x2d, 0x05},
				{0x23a, 0x2d, 0x05},
				{0x23c, 0x2d, 0x05},
				{0x23e, 0x2d, 0x05},
				{0x240, 0x2f, 0x05},
				{0x242, 0x2f, 0x05},
				{0x244, 0x2f, 0x05},
				{0x246, 0x2f, 0x05},
				{0x248, 0x2f, 0x05},
				{0x24a, 0x2f, 0x05},
				{0x24c, 0x2f, 0x05},
				{0x24e, 0x2f, 0x05},
				{0x250, 0x31, 0x05},
				{0x252, 0x31, 0x05},
				{0x254, 0x31, 0x05},
				{0x256, 0x31, 0x05},
				{0x258, 0x31, 0x05},
				{0x25a, 0x31, 0x05},
				{0x25c, 0x31, 0x05},
				{0x25e, 0x31, 0x05},
				{0x260, 0x32, 0x05},
				{0x262, 0x32, 0x05},
				{0x264, 0x32, 0x05},
				{0x266, 0x32, 0x05},
				{0x268, 0x32, 0x05},
				{0x26a, 0x32, 0x05},
				{0x26c, 0x32, 0x05},
				{0x26e, 0x32, 0x05},
				{0x270, 0x32, 0x0f},
				{0x278, 0x32, 0x0f},
				{0x27c, 0x34, 0x0f},
			},
		},
		{
			instrument: 3,
			volume:     70,
			tones: [][3]uint16{
				{0xc0, 0x00, 0x05},
				{0xc4, 0x05, 0x0f},
				{0xc8, 0x00, 0x05},
				{0xca, 0x00, 0x05},
				{0xcc, 0x05, 0x0f},
				{0xd0, 0x00, 0x05},
				{0xd4, 0x05, 0x0f},
				{0xd8, 0x00, 0x05},
				{0xda, 0x00, 0x05},
				{0xdc, 0x05, 0x0f},
				{0xe0, 0x00, 0x05},
				{0xe4, 0x05, 0x0f},
				{0xe8, 0x00, 0x05},
				{0xea, 0x00, 0x05},
				{0xec, 0x05, 0x0f},
				{0xf0, 0x00, 0x05},
				{0xf4, 0x05, 0x0f},
				{0xf8, 0x00, 0x05},
				{0xfa, 0x00, 0x05},
				{0xfc, 0x05, 0x0f},
				{0x100, 0x00, 0x05},
				{0x104, 0x05, 0x0f},
				{0x108, 0x00, 0x05},
				{0x10a, 0x00, 0x05},
				{0x10c, 0x05, 0x0f},
				{0x110, 0x00, 0x05},
				{0x114, 0x05, 0x0f},
				{0x118, 0x00, 0x05},
				{0x11a, 0x00, 0x05},
				{0x11c, 0x05, 0x0f},
				{0x120, 0x00, 0x05},
				{0x124, 0x05, 0x0f},
				{0x128, 0x00, 0x05},
				{0x12a, 0x00, 0x05},
				{0x12c, 0x05, 0x0f},
				{0x130, 0x00, 0x05},
				{0x134, 0x05, 0x0f},
				{0x138, 0x00, 0x05},
				{0x13a, 0x00, 0x05},
				{0x13c, 0x05, 0x0f},
				{0x140, 0x00, 0x05},
				{0x144, 0x05, 0x0f},
				{0x148, 0x00, 0x05},
				{0x14a, 0x00, 0x05},
				{0x14c, 0x05, 0x0f},
				{0x150, 0x00, 0x05},
				{0x154, 0x05, 0x0f},
				{0x158, 0x00, 0x05},
				{0x15a, 0x00, 0x05},
				{0x15c, 0x05, 0x0f},
				{0x160, 0x00, 0x05},
				{0x164, 0x05, 0x0f},
				{0x168, 0x00, 0x05},
				{0x16a, 0x00, 0x05},
				{0x16c, 0x05, 0x0f},
				{0x170, 0x00, 0x05},
				{0x174, 0x05, 0x0f},
				{0x178, 0x00, 0x05},
				{0x17a, 0x00, 0x05},
				{0x17c, 0x05, 0x0f},
				{0x180, 0x00, 0x05},
				{0x184, 0x05, 0x0f},
				{0x188, 0x00, 0x05},
				{0x18a, 0x00, 0x05},
				{0x18c, 0x05, 0x0f},
				{0x190, 0x00, 0x05},
				{0x194, 0x05, 0x0f},
				{0x198, 0x00, 0x05},
				{0x19a, 0x00, 0x05},
				{0x19c, 0x05, 0x0f},
				{0x1a0, 0x00, 0x05},
				{0x1a4, 0x05, 0x0f},
				{0x1a8, 0x00, 0x05},
				{0x1aa, 0x00, 0x05},
				{0x1ac, 0x05, 0x0f},
				{0x1b0, 0x00, 0x05},
				{0x1b4, 0x05, 0x0f},
				{0x1b8, 0x00, 0x05},
				{0x1ba, 0x00, 0x05},
				{0x1bc, 0x05, 0x0f},
				{0x1c0, 0x00, 0x0f},
				{0x1c6, 0x00, 0x0f},
				{0x1cc, 0x00, 0x0f},
				{0x1d8, 0x00, 0x05},
				{0x1da, 0x00, 0x05},
				{0x1dc, 0x00, 0x05},
				{0x1de, 0x00, 0x05},
				{0x1e0, 0x00, 0x05},
				{0x1e2, 0x00, 0x05},
				{0x1e4, 0x0a, 0x0f},
				{0x1e8, 0x00, 0x05},
				{0x1ea, 0x00, 0x05},
				{0x1ec, 0x0a, 0x0f},
				{0x1f0, 0x00, 0x0f},
				{0x1f4, 0x05, 0x0f},
				{0x1f8, 0x00, 0x05},
				{0x1fa, 0x00, 0x05},
				{0x1fc, 0x05, 0x0f},
				{0x200, 0x00, 0x05},
				{0x204, 0x05, 0x0f},
				{0x208, 0x00, 0x05},
				{0x20a, 0x00, 0x05},
				{0x20c, 0x05, 0x0f},
				{0x210, 0x00, 0x05},
				{0x214, 0x05, 0x0f},
				{0x218, 0x00, 0x05},
				{0x21a, 0x00, 0x05},
				{0x21c, 0x05, 0x0f},
				{0x220, 0x00, 0x05},
				{0x222, 0x00, 0x05},
				{0x224, 0x0a, 0x0f},
				{0x228, 0x00, 0x05},
				{0x22a, 0x00, 0x05},
				{0x22c, 0x0a, 0x0f},
				{0x230, 0x00, 0x0f},
				{0x234, 0x05, 0x0f},
				{0x238, 0x00, 0x05},
				{0x23a, 0x00, 0x05},
				{0x23c, 0x05, 0x0f},
				{0x240, 0x00, 0x05},
				{0x244, 0x05, 0x0f},
				{0x248, 0x00, 0x05},
				{0x24a, 0x00, 0x05},
				{0x24c, 0x05, 0x0f},
				{0x250, 0x00, 0x05},
				{0x254, 0x05, 0x0f},
				{0x258, 0x00, 0x05},
				{0x25a, 0x00, 0x05},
				{0x25c, 0x05, 0x0f},
				{0x260, 0x00, 0x05},
				{0x264, 0x05, 0x0f},
				{0x268, 0x00, 0x05},
				{0x26a, 0x00, 0x05},
				{0x26c, 0x05, 0x0f},
				{0x270, 0x00, 0x05},
				{0x274, 0x05, 0x0f},
				{0x278, 0x00, 0x05},
				{0x27a, 0x00, 0x05},
				{0x27c, 0x05, 0x0f},
			},
		},
	},
}
