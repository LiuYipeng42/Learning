package main

type Logger struct {
	flags int
	prefix string
}

func (l *Logger) Flags() int {
	return l.flags
}

func (l *Logger) SetFlags(flag int) {
	l.flags = flag
}

func (l *Logger) Prefix() string {
	return l.prefix
}

func (l *Logger) SetPrefix(prefix string) {
	l.prefix = prefix
}

