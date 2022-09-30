# Multitouch

This package implements reading device events from the Linux evdev interface and
turning them into simplified touch events, touch begin, touch move and touch
end.

This makes it useful for simplifying multitouch handling with popular libraries
such as Fyne.

# Note

This library was built to handle touch events for the SumUp Solo device. The
device has the display rotated 90 degrees clockwise relative to the physical
hardware. Hence, a transformation is implemented to counter this. 

As this is likely undesireable behaviour for most. You may disable this by
passing false to the initialisation.
