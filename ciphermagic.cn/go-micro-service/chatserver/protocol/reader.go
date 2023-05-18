package protocol

import (
	"bufio"
	"io"
	"log"
)

type Reader struct {
	reader *bufio.Reader
}

func NewReader(reader io.Reader) *Reader {
	return &Reader{reader: bufio.NewReader(reader)}
}

func (r *Reader) Read() (interface{}, error) {
	cmd, err := r.reader.ReadString(' ')
	if err != nil {
		return nil, err
	}

	switch cmd {
	case "MESS ":
		name, err := r.reader.ReadString(' ')
		if err != nil {
			return nil, err
		}
		message, err := r.reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		return MessCmd{name[:len(name)-1], message[:len(message)-1]}, nil
	case "SEND ":
		message, err := r.reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		return SendCmd{message[:len(message)-1]}, nil
	case "NAME ":
		name, err := r.reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		return NameCmd{name[:len(name)-1]}, nil
	default:
		log.Printf("Unknown command: %v", cmd)
	}
	return nil, UnknownCommand
}
