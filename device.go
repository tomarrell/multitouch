package multitouch

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"unsafe"

	"golang.org/x/sys/unix"
)

type evdevEvent struct {
	Time  unix.Timeval // time in seconds since epoch at which event occurred
	Type  EventType    // event type - one of ecodes.EV_*
	Code  EventCode    // event code related to the event type
	Value int32        // event value related to the event type
}

const eventsize = int(unsafe.Sizeof(evdevEvent{}))

type inputDevice struct {
	File *os.File
}

func open(path string) (*inputDevice, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return &inputDevice{File: f}, nil
}

// Read a single event.
func (idev *inputDevice) Read() (*evdevEvent, error) {
	buffer := make([]byte, eventsize)
	_, err := idev.File.Read(buffer)
	if err != nil {
		return nil, fmt.Errorf("failed to read buffer: %v", err)
	}

	event := evdevEvent{}
	b := bytes.NewBuffer(buffer)
	if err := binary.Read(b, binary.LittleEndian, &event); err != nil {
		return nil, fmt.Errorf("failed to read binary: %v", err)
	}

	return &event, err
}
