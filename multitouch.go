package multitouch

type Action int

const (
	ActionUnknown Action = iota
	ActionBegin
	ActionMove
	ActionEnd
)

type TouchEvent struct {
	ID     int
	Action Action
	X      int
	Y      int
}

type Multitouch struct {
	dev   *inputDevice
	out   chan *TouchEvent
	slots [10]*TouchEvent
}

// NewMultitouch creates a new multitouch device, interpreting evdev events from
// the kernel. Path is the string path to the evdevice.
func NewMultitouch(path string) (*Multitouch, error) {
	dev, err := open(path)
	if err != nil {
		return nil, err
	}

	m := &Multitouch{
		dev:   dev,
		out:   make(chan *TouchEvent),
		slots: [10]*TouchEvent{},
	}

	return m, nil
}

func (m *Multitouch) Begin() {
	m.processInput()
}

func (m *Multitouch) Next() *TouchEvent {
	return <-m.out
}

func (m *Multitouch) processInput() {
	var (
		modified = []int{}
		slot     = 0
		// state    = 0
	)

	for {
		e, err := m.dev.Read()
		if err != nil {
			panic(err)
		}

		switch e.Type {
		case EV_KEY:
			if e.Code != BTN_TOUCH {
				panic("unknown even code for EV_KEY")
			}
			// state = int(e.Value)

		case EV_ABS:
			switch e.Code {
			case ABS_MT_TRACKING_ID:
				if e.Value == -1 { // Slot can be freed
					m.slots[slot].Action = ActionEnd
				}

			case ABS_MT_SLOT:
				slot = int(e.Value)
				if m.slots[slot] == nil {
					m.slots[slot] = &TouchEvent{
						ID:     slot,
						Action: ActionBegin,
					}
				} else if m.slots[slot].Action == ActionBegin {
					m.slots[slot].Action = ActionMove
				}

			case ABS_MT_POSITION_X:
				modified = append(modified, slot)
				m.slots[slot].X = int(e.Value)
				break
			case ABS_MT_POSITION_Y:
				modified = append(modified, slot)
				m.slots[slot].Y = int(e.Value)
				break
			}

		case EV_SYN:
			for _, s := range modified {
				m.out <- m.slots[s]
				if m.slots[s].Action == ActionEnd {
					m.slots[s] = nil
				}
			}
		}
	}
}
