package clog

type Formatter interface {
	Format(entry *Entry) error
}
