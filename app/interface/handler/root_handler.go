package handler

type RootHandler struct {
	AuthHandler  *AuthHandler
	MeHandler    *MeHandler
	AliveHandler *AliveHandler
}

func NewRootHandler(
	authHandlr *AuthHandler,
	meHandler *MeHandler,
	aliveHandler *AliveHandler,
) (*RootHandler, error) {
	return &RootHandler{
		AuthHandler:  authHandlr,
		MeHandler:    meHandler,
		AliveHandler: aliveHandler,
	}, nil
}
