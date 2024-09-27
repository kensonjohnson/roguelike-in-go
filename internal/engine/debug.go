package engine

type debug struct {
	on bool
}

var Debug = &debug{
	on: false,
}

func (d *debug) TurnOn() {
	d.on = true
}

func (d *debug) TurnOff() {
	d.on = false
}

func (d *debug) On() bool {
	return d.on
}
