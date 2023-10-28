package io

type Read interface {
	Scan()
	Input()
}

type Write interface {
	Echo()
	Output()
}

type ReaderWriter interface {
	Read
	Write
}

func ReadFile(r Read) {
	r.Scan()
}

func WriteFile(w Write) {
	w.Echo()
	w.Output()
}

func WriteReadFile(rw ReaderWriter) {
	rw.Scan()
	rw.Input()
	rw.Echo()
	rw.Output()
}
