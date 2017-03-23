package commander

import (
	"bytes"
	"strings"
)

const _HelpMessageSpace string = "  "

type _HelpMessage struct {
	buffer bytes.Buffer
}

func (h *_HelpMessage) write(s string, spaces int) (int, error) {
	for ; spaces > 0; spaces-- {
		s = _HelpMessageSpace + s
	}
	return h.buffer.WriteString(s)
}

func (h *_HelpMessage) Line() *_HelpMessage {
	h.buffer.WriteString("\n")
	return h
}

func (h *_HelpMessage) Description(s string) error {
	_, err := h.write(s+"\n\n", 1)
	return err
}

func (h *_HelpMessage) Title(s string) error {
	_, err := h.write(s+":\n", 1)
	return err
}

func (h *_HelpMessage) Subtitle(s string) error {
	if strings.Contains(s, "\n") {
		strs := strings.Split(s, "\n")
		for _, str := range strs {
			if len(strings.TrimSpace(str)) == 0 {
			} else if _, err := h.write(str+"\n", 2); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := h.write(s+"\n", 2)
	return err
}

func (h _HelpMessage) String() string {
	return h.buffer.String()
}
