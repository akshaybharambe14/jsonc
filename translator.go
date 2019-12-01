package jsonc

const (
	quote   = 34 // ("")
	space   = 32 // ( )
	tab     = 9  // (	)
	newLine = 10 // (\n)
	star    = 42 // (*)
	slash   = 47 // (/)
)

type state int

const (
	stopped state = iota
	canStart
	started
	canStop
)

type comment struct {
	state       state
	isJSON      bool
	isMultiLine bool
}

func (cmt *comment) setState(s state) { cmt.state = s }

func (cmt *comment) checkState(s state) bool { return cmt.state == s }

func extract(s []byte) []byte {
	vj := make([]byte, len(s))
	i := 0
	cmt := &comment{
		state: stopped,
	}
	for _, s := range s {
		if cmt.checkState(stopped) {
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
				cmt.setState(canStart)
				continue
			}
		}

		if cmt.checkState(canStart) && (s == slash || s == star) {
			cmt.setState(started)
			if s == star {
				cmt.isMultiLine = true
			}
			continue
		}

		if cmt.checkState(started) {
			if s == star {
				cmt.setState(canStop)
				continue
			}

			if s == newLine {
				if cmt.isMultiLine {
					continue
				}
				cmt.setState(stopped)
			}

			continue
		}

		if cmt.checkState(canStop) && (s == slash) {
			cmt.setState(stopped)
			cmt.isMultiLine = false
			continue
		}
	addJSON:
		vj[i] = s
		i++
	}

	return vj[:i]
}
