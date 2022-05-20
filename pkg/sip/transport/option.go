package transport

// Listen method options
type ListenOption interface {
	ApplyListen(opts *ListenOptions)
}

type ListenOptions struct {
}
