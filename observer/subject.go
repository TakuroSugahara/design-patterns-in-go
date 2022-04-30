package observer

type subject interface {
	register(observer)
	deregister(observer)
	notify()
}
