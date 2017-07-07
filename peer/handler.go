package peer

import "github.com/bft/protos"

type Handler struct {

}

func NewPeerHandler(coord MessageHandlerCoordinator, stream ChatStream, initiatedStream bool) (MessageHandler, error) {

	d := &Handler{}

	return d, nil
}

func (h *Handler) HandleMessage(msg *protos.Message) error {
	return nil
}

func (h *Handler) SendMessage(msg *protos.Message) error {
	return nil
}

func (h *Handler) To() (protos.PeerEndpoint, error) {
	return protos.PeerEndpoint{}, nil
}

func (h *Handler) Stop() error {
	return nil
}
