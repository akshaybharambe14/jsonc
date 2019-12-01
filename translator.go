package jsonc

const (
	quote   = 34 // ("")
	space   = 32 // ( )
	tab     = 9  // (	)
	newLine = 10 // (\n)
	star    = 42 // (*)
	slash   = 47 // (/)
)

const (
	stopped state = iota
	canStart
	started
	canStop
)

type (
	state   int
	comment struct {
		state       state
		isJSON      bool
		isMultiLine bool
	}
)

func extract(s []byte) []byte {
	var (
		res = make([]byte, len(s))
		i   = 0
		cmt = comment{}
	)

	for _, s := range s {

		switch cmt.state {

		case stopped:
			if s == quote {
				cmt.isJSON = !cmt.isJSON
			}

			if cmt.isJSON {
				goto addJSON
			}

			if s == space || s == tab || s == newLine {
				continue
			}

			if s == slash {
				cmt.state = canStart
				continue
			}

		case canStart:
			if s == slash || s == star {
				cmt.state = started
				if s == star {
					cmt.isMultiLine = true
				}
				continue
			}

		case started:
			if s == star {
				cmt.state = canStop
				continue
			}

			if s == newLine {
				if cmt.isMultiLine {
					continue
				}
				cmt.state = stopped
			}

			continue

		case canStop:
			if s == slash {
				cmt.state = stopped
				cmt.isMultiLine = false
				continue
			}

		}

	addJSON:
		res[i] = s
		i++
	}

	return res[:i]
}
