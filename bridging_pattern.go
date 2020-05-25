package desgin_patterns

type Device interface {
	isEnabled() bool
	enable()
	disable()
	getVolume() int
	setVolume(percent int)
	getChannel() int
	setChannel(channel int) int
}

type RemoteControl struct {
	device Device
}

func NewRemoteControl(d Device) RemoteControl {
	return RemoteControl{device: d}
}

func (r *RemoteControl) TogglePower() {
	if r.device.isEnabled() {
		r.device.disable()
	} else {
		r.device.enable()
	}
}

func (r *RemoteControl) VolumeDown() {
	r.device.setVolume(r.device.getVolume() - 10)
}

func (r *RemoteControl) VolumeUp() {
	r.device.setVolume(r.device.getVolume() + 10)
}

func (r *RemoteControl) ChannelDown() {
	r.device.setChannel(r.device.getChannel() - 1)
}

func (r *RemoteControl) ChannelUp() {
	r.device.setChannel(r.device.getChannel() + 1)
}
