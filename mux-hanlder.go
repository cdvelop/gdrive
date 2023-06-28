package gdrive

func (g *drive) addHttpHandlers() {
	g.mux.HandleFunc("/login", g.Login)
	g.mux.HandleFunc("/callback", g.Callback)
	g.mux.HandleFunc("/drive", g.Read)
}
