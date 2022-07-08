package multitouch

type EventType uint16

// Event Types
const (
	EV_SYN       EventType = 0x00 // EV_SYN event values are undefined
	EV_KEY       EventType = 0x01 // Reporting of key events, including touch
	EV_REL       EventType = 0x02 // Relative position, prefer EV_ABS
	EV_ABS       EventType = 0x03 // Absolute position
	EV_MSC       EventType = 0x04
	EV_SW        EventType = 0x05
	EV_LED       EventType = 0x11
	EV_SND       EventType = 0x12
	EV_REP       EventType = 0x14
	EV_FF        EventType = 0x15
	EV_PWR       EventType = 0x16
	EV_FF_STATUS EventType = 0x17
	EV_MAX       EventType = 0x1f
	EV_CNT       EventType = (EV_MAX + 1)
)

type EventCode uint16

// Event Codes
const (
	SYN_REPORT EventCode = 0x0 // Sync report, defines event boundaries

	BTN_TOUCH EventCode = 0x14a // Dec: 330 - Touch bool

	ABS_MT_SLOT        EventCode = 0x2f // Dec: 47 - MT slot being modified
	ABS_MT_TOUCH_MAJOR EventCode = 0x30 // Dec: 48 - Major axis of touching ellipse
	ABS_MT_TOUCH_MINOR EventCode = 0x31 // Dec: 49 - Minor axis (omit if circular)
	ABS_MT_WIDTH_MAJOR EventCode = 0x32 // Dec: 50 - Major axis of approaching ellipse
	ABS_MT_WIDTH_MINOR EventCode = 0x33 // Dec: 51 - Minor axis (omit if circular)
	ABS_MT_ORIENTATION EventCode = 0x34 // Dec: 52 - Ellipse orientation
	ABS_MT_POSITION_X  EventCode = 0x35 // Dec: 53 - Center X touch position
	ABS_MT_POSITION_Y  EventCode = 0x36 // Dec: 54 - Center Y touch position
	ABS_MT_TOOL_TYPE   EventCode = 0x37 // Dec: 55 - Type of touching device
	ABS_MT_BLOB_ID     EventCode = 0x38 // Dec: 56 - Group a set of packets as a blob
	ABS_MT_TRACKING_ID EventCode = 0x39 // Dec: 57 - Unique ID of initiated contact
	ABS_MT_PRESSURE    EventCode = 0x3a // Dec: 58 - Pressure on contact area
	ABS_MT_DISTANCE    EventCode = 0x3b // Dec: 59 - Contact hover distance
	ABS_MT_TOOL_X      EventCode = 0x3c // Dec: 60 - Center X tool position
	ABS_MT_TOOL_Y      EventCode = 0x3d // Dec: 61 - Center Y tool positio
)
