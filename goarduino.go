package goarduino

import (
    "github.com/tarm/goserial"
    "fmt"
    "strings"
)

type Arduino struct {
  serialDev string
  baudrate  int
  conn      *io.ReadWriteCloser
}

func NewArduino(dev string, baud int) (board *Arduino, err error) {
	var conn io.ReadWriteCloser

	c := &serial.Config{Name: dev, Baud: baud}
	conn, err = serial.OpenPort(c)
	if err != nil {
		panic(err)
		return
	}
	client = &Arduino{
    serialDev: dev,
    baud:      baud,
    conn:      &conn
  }

	return
}

func (c *Arduino) Close() {
  (*c.conn).Close()
}