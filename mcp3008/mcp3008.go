package mcp3008

// #include "lib/wiringPiSPI.c"
import "C"
import (
	"errors"
	"unsafe"
)

const (
	SINGLE_ENDED = 128
	DIFFERENTIAL = 0

	CHAN_0 = 0
	CHAN_1 = 1
)

var (
	buffer [3]C.uchar
	ret    C.int
	recv   uint16
)

func Setup(channel int, speed int) error {
	fd := C.wiringPiSPISetup(C.int(channel), C.int(speed))
	if fd <= -1 {
		return errors.New("Failed to connect to SPI device, are you running this as root?")
	}
	return nil
}

func DigitalToAnalog(d uint16) float32 {
	return float32(d) / 10.23 * 0.05
}

func controlBits(readType int, channel int) byte {
	return byte(SINGLE_ENDED | ((channel&7)<<4)&0xff)
}

func ReadADC(channel int) (uint16, error) {
	buffer[0] = 1                                           // start bit
	buffer[1] = C.uchar(controlBits(SINGLE_ENDED, channel)) // control bits for selected channel
	buffer[2] = 0                                           // don't care byte

	ret = C.wiringPiSPIDataRW(0, (*C.uchar)(unsafe.Pointer(&buffer)), C.int(len(buffer)))
	if ret == -1 {
		return 0, errors.New("Failed to read from SPI device")
	}

	recv = (uint16(buffer[1]) << 8) & 0x300
	recv |= (uint16(buffer[2]) & 0xff)

	return recv, nil
}
