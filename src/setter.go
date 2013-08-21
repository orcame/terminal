package terminal

type setter interface{
	setStyle(code uint,t *writer)
	resetStyle(t *writer)
	setTitle(t *writer)
}
