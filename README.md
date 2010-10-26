Goblin, the reimplementation of unix userspace in the Go programming language (Golang)
===================

The reason i started working on this is that I needed a few of the unix tools, preferrably in a fashion that's usable from go aswell, so I took uriel's idea and started it off.


Current problems: 
 * wc doesn't count as many characters as the GNU (or plan9) wc for now, i'm not sure why, maybe because I use ReadRune instead of ReadByte
 * cat is way slow, probably because it uses channels to communicate