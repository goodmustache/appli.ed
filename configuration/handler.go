package configuration

type Handler struct{}

func NewHandler() *Handler {
	return new(Handler)
}

func (Handler) ReadConfig(path string) (WatchConfiguration, error) {
	return Config{}, nil
}
