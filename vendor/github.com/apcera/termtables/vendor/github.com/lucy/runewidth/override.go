// Defines exceptions for gen.go

// +build ignore

package main

// nabbed from https://github.com/jquast/wcwidth/blob/07cea7f70f298c7124c5a7838848c392516a676d/wcwidth/wcwidth.py#L159-L173
var overrides = []override{
	// Control codes
	{0x0000, 0x001F, -1},
	{0x007F, 0x009F, -1},
	// Misc zero width from Cf
	{0x0000, 0x0000, 0},
	{0x200B, 0x200F, 0},
	{0x2028, 0x2029, 0},
	{0x202A, 0x202E, 0},
	{0x2060, 0x2063, 0},
	{0x034F, 0x034F, 0},
}
