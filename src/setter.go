package terminal

type setter interface{
	save()
	setStyle(codes string,t *writer)
	resetStyle(t *writer)
	setTitle(t *writer)
}
