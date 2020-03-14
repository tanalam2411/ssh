package ssh

import "encoding/binary"

func parsePtyRequest(s []byte) (pty Pty, ok bool) {
	term, s, ok := parseString(s)
	if !ok {
		return
	}

	width32, s, ok := parseUint32(s)
	if !ok {
		return
	}

	height32, _, ok := parseUint32(s)
	if !ok {
		return
	}

	pty = Pty{
		Term: term,
		Window: Window{
			Width:  int(width32),
			Height: int(height32),
		},
	}

	return
}

func parseString(in []byte) (out string, rest []byte, ok bool) {
	if len(in) < 4 {
		return
	}

	length := binary.BigEndian.Uint32(in)
	if uint32(len(in)) < 4+length {
		return
	}

	out = string(in[4 : 4+length])
	rest = in[4+length:]
	return
}

func parseUint32(in []byte) (uint32, []byte, bool) {
	if len(in) < 4 {
		return 0, nil, false
	}
	return binary.BigEndian.Uint32(in), in[4:], true
}

func parseWinchRequest(s []byte) (win Window, ok bool) {
	width32, s, ok := parseUint32(s)
	if width32 < 1 {
		ok = false
	}
	if !ok {
		return
	}

	height32, _, ok := parseUint32(s)
	if height32 < 1 {
		ok = false
	}
	if !ok {
		return
	}

	win = Window{
		Width:  int(width32),
		Height: int(height32),
	}
	return
}