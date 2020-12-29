// +build !windows

package term

import (
	"github.com/pkg/term/termios"
	"golang.org/x/sys/unix"
)

// SetRaw put terminal into a raw mode
func SetRaw(fd int) error {
	n, err := getOriginalTermios(fd)
	if err != nil {
		return err
	}

	n.Iflag &^= unix.IGNBRK | unix.BRKINT | unix.PARMRK |
		unix.ISTRIP | unix.INLCR | unix.IGNCR |
		unix.ICRNL | unix.IXON
	n.Lflag &^= unix.ECHO | unix.ICANON | unix.IEXTEN | unix.ISIG | unix.ECHONL
	n.Cflag &^= unix.CSIZE | unix.PARENB
	n.Cflag |= unix.CS8 // Set to 8-bit wide.  Typical value for displaying characters.
	n.Cc[unix.VMIN] = 1
	n.Cc[unix.VTIME] = 0

	return termios.Tcsetattr(uintptr(fd), termios.TCSANOW, &n)
}
