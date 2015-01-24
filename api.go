package main

// The main Resource of this API
// It abstracts all the REST service
// through an idiomatic Go way
type Api struct {
	Version int
	Message string `json:"welcome"` // Rename this property

	Categories *Categories `json:"-"` // This Resource will be accessible on /api/categories
	DB         *DB         `json:"-"` // This Resource will be injected when needed
	// The DB Resource has no methods accessible by the client
}

// [GET] /api
// Receives the initial state of this resource and return it
// The initial state of all Resources is passed
// in the creation of the Resource in the main file
func (a *Api) GET() *Api {
	return a
}
