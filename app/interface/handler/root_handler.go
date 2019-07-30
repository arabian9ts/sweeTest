package handler

type RootHandler struct {
	AuthHandler *AuthHandler
	MeHandler   *MeHandler
}

func NewRootHandler(
	authHandlr *AuthHandler,
	meHandler *MeHandler,
) (*RootHandler, error) {
	return &RootHandler{
		AuthHandler: authHandlr,
		MeHandler:   meHandler,
	}, nil
}
