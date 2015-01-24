package main

// The main resource of this API
type Api struct {
	Version    int
	Env        Env
	Categories *Categories
	DB         *DB
}

func (a *Api) GET() string {
	return "This is the REST API for a book store"
}
