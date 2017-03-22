package commander

import "bytes"

type _HelpMessage struct {
	buffer bytes.Buffer
}

func (h *_HelpMessage) write(s string) (int, error) {
	return h.buffer.WriteString("  " + s)
}

func (h *_HelpMessage) Line() *_HelpMessage {
	h.buffer.WriteString("\n")
	return h
}

func (h *_HelpMessage) Description(s string) (int, error) {
	return h.write(s + "\n\n")
}

func (h *_HelpMessage) Title(s string) (int, error) {
	return h.write(s + ":\n")
}

func (h *_HelpMessage) Subtitle(s string) (int, error) {
	return h.write("  " + s + "\n")
}

func (h _HelpMessage) String() string {
	return h.buffer.String()
}
