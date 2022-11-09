package models

type Button struct {
	EventListeners map[string][]chan string
}

func MakeButton() *Button {
	button := new(Button)
	button.EventListeners = make(map[string][]chan string)
	return button
}

func (button *Button) AddEventListener(event string, listener chan string) {
	for _, eventListener := range button.EventListeners[event] {

		if eventListener == listener {
			return
		}
	}

	button.EventListeners[event] = append(button.EventListeners[event], listener)
}

func (button *Button) RemoveEventListener(event string, listener chan string) {
	for i, eventListener := range button.EventListeners[event] {
		if eventListener == listener {
			button.EventListeners[event] = append(button.EventListeners[event][:i], button.EventListeners[event][i+1:]...)
			break
		}
	}
}

func (button *Button) DispatchEvent(event string, data string) {
	for _, eventListener := range button.EventListeners[event] {
		go func(handler chan string) {
			handler <- data
		}(eventListener)
	}
}

func (button *Button) TriggerEvent(event string, data string) {
	button.DispatchEvent(event, data)
}
