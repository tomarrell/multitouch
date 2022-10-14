package multitouch

type Action int

func (a Action) String() string {
	switch a {
	case ActionUnknown:
		return "ActionUnknown"
	case ActionBegin:
		return "ActionBegin"
	case ActionMove:
		return "ActionMove"
	case ActionEnd:
		return "ActionEnd"
	}

	return "ActionUnknown"
}

const (
	screenWidth  = 480
	screenHeight = screenWidth
)

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

type slots [10]*TouchEvent

type Multitouch struct {
	dev       *inputDevice
	out       chan TouchEvent
	slots     slots
	transform bool
}

// NewMultitouch creates a new multitouch device, interpreting evdev events from
// the kernel. Path is the string path to the evdevice.
func NewMultitouch(path string, transform bool) (*Multitouch, error) {
	dev, err := open(path)
	if err != nil {
		return nil, err
	}

	m := &Multitouch{
		dev:       dev,
		out:       make(chan TouchEvent),
		slots:     [10]*TouchEvent{},
		transform: transform,
	}

	return m, nil
}

func (m *Multitouch) Begin() {
	m.processInput()
}

func (m *Multitouch) Out() chan TouchEvent {
	return m.out
}

func (m *Multitouch) Next() TouchEvent {
	return <-m.out
}

func (m *Multitouch) processInput() {
	var (
		modified = map[int]bool{}
		slot     = 0
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

		case EV_ABS:
			switch e.Code {
			case ABS_MT_TRACKING_ID:
				if e.Value == -1 { // Slot can be freed
					m.slots[slot].Action = ActionEnd
					modified[slot] = true
				}

			case ABS_MT_SLOT:
				slot = int(e.Value)
				if s, exists := getSlot(&m.slots, slot); exists && s.Action == ActionBegin {
					m.slots[slot].Action = ActionMove
				}

			case ABS_MT_POSITION_X:
				modified[slot] = true
				s, _ := getSlot(&m.slots, slot)
				s.X = int(e.Value)
				break

			case ABS_MT_POSITION_Y:
				modified[slot] = true
				s, _ := getSlot(&m.slots, slot)
				s.Y = int(e.Value)
				break
			}

		case EV_SYN:
			for k := range modified {
				o := *m.slots[k]

				if m.transform {
					o.X, o.Y = transformPoint(o.X, o.Y)
				}

				m.out <- o
				delete(modified, k)
				if m.slots[k].Action == ActionEnd {
					m.slots[k] = nil
				}
			}
		}
	}
}

func transformPoint(x, y int) (xp int, yp int) {
	return screenWidth - y, x
}

func getSlot(slots *slots, slot int) (event *TouchEvent, exists bool) {
	if slots[slot] != nil {
		return slots[slot], true
	}

	slots[slot] = &TouchEvent{
		ID:     slot,
		Action: ActionBegin,
	}

	return slots[slot], false
}
