package mboxparser

import (
	bmbox "github.com/blabber/mbox"
	"io"
	"net/mail"
	"os"
)

func Read(r io.Reader) (*Mbox, error) {
	var messages []*Message

	msgs := bmbox.NewReader(r)
	r, err := msgs.NextMessage()
	if err != nil {
		return nil, err
	}
	for err != io.EOF && r != nil {
		msg, err := mail.ReadMessage(r)
		if err != nil {
			return nil, err
		}
		messages = append(messages, Decode(msg))
		r, err = msgs.NextMessage()
	}

	return &Mbox{
		Messages: messages,
	}, nil
}

func ReadFile(filename string) (*Mbox, error) {
	fp, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return Read(fp)
}
