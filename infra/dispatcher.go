package infra

import (
	"fmt"
	"sync"
)

// Dispatcher is to dispatch the response to the corresponding channel.
type Dispatcher struct {
	channels map[uint32]*ProtobufChan // key is serial number
	mu       sync.Mutex
}

// NewDispatcher creates a new dispatcher.
func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		channels: make(map[uint32]*ProtobufChan),
		mu:       sync.Mutex{},
	}
}

// Register registers a channel with a serial number.
func (d *Dispatcher) Register(sn uint32, ch *ProtobufChan) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if sn > 0 && d.channels[sn] != nil {
		d.channels[sn].Close() // close the old channel
	}

	d.channels[sn] = ch
}

// Dispatch dispatches the response to the corresponding channel.
func (d *Dispatcher) Dispatch(sn uint32, body []byte) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	ch := d.channels[sn]
	if ch == nil {
		ch = d.channels[0] // use the default channel for notifications
	}

	if ch == nil {
		return fmt.Errorf("channel not found for sn %d", sn)
	}

	if err := ch.Send(body); err != nil {
		return err
	}

	// deregister the channel after sending the response.
	// we do not close the channel here, it's the responsibility of the caller.
	// the caller can reuse the channel for next serial number if needed.
	if sn != 0 {
		delete(d.channels, sn)
	}

	return nil
}

// Close closes all channels.
func (d *Dispatcher) Close() {
	d.mu.Lock()
	defer d.mu.Unlock()

	for sn, ch := range d.channels {
		ch.Close()
		delete(d.channels, sn)
	}
}

// DispatcherHub is to store dispatchers for different proto id, key is proto id.
type DispatcherHub struct {
	dispatchers map[uint32]*Dispatcher
	mu          sync.Mutex
}

// NewDispatcherHub creates a new dispatcher hub.
func NewDispatcherHub() *DispatcherHub {
	return &DispatcherHub{
		dispatchers: make(map[uint32]*Dispatcher),
		mu:          sync.Mutex{},
	}
}

// Register registers a dispatcher with a proto id and serial number.
func (h *DispatcherHub) Register(id uint32, sn uint32, ch *ProtobufChan) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.dispatchers[id] == nil {
		h.dispatchers[id] = NewDispatcher()
	}

	h.dispatchers[id].Register(sn, ch)
}

// Dispatch dispatches the response to the corresponding channel.
func (h *DispatcherHub) Dispatch(id uint32, sn uint32, body []byte) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.dispatchers[id] == nil {
		return fmt.Errorf("dispatcher not found for proto id %d", id)
	}

	return h.dispatchers[id].Dispatch(sn, body)
}

// Close closes all dispatchers.
func (h *DispatcherHub) Close() {
	h.mu.Lock()
	defer h.mu.Unlock()

	for _, dispatcher := range h.dispatchers {
		dispatcher.Close()
	}
}
