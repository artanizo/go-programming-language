package main

import (
	"bytes"
)

//comma 3.11
func comma(s string) string {
	if len(s) == 0 {
		return s
	}

	var buf bytes.Buffer

	rs := s
	if s[0] == '-' {
		rs = s[1:]
		buf.WriteByte(s[0])
	}

	for i, v := range rs {
		if v == '.' {
			buf.Write([]byte(rs[i:]))
			break
		}
		if i > 1 && i%3 == 0 {
			buf.WriteString(",")
		}

		buf.WriteRune(v)
	}

	return buf.String()
}
