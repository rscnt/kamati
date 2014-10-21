package kamatis

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"labix.org/v2/mgo"
)

/** Martini Server struct wrapper */
type Server *martini.ClassicMartini

/*
 * Create a new *martini.ClassicMartini server.
 * Using JSON to render mongodb results.
 * ---
 * GET     /users
 * POST    /users
 * GET     /users/:id
 * PUT     /users/:id
 * DELETE  /users/:id
 * ---
 * GET     /todolist
 * POST    /todolist
 * GET     /todolist/:id
 * PUT     /todolist/:id
 * DELETE  /todolist/:id
 */
func NewServer(session *DataBaseSession) Server {
	m := Server(martini.Classic())
	m.Use(render.Renderer(render.Options{
		IndentJSON: true,
	}))
	m.Use(session.DataBase())

	// Define "GET /users"
	m.Get("/users", func(r render.Render, db *mgo.Database) {
		r.JSON(200, fetchAllUsers(db))
	})

	m.Post("/users", binding.Json(User{}),
		func(user User, r render.Render, db *mgo.Database) {
			if user.valid() {
				err := db.C("users").Insert(user)
				if err == nil {
					r.JSON(201, user)
				} else {
					r.JSON(400, map[string]string{"error": err.Error()})
				}
			} else {
				r.JSON(400, map[string]string{"error": "Not a valid structure"})
			}
		})

	return m
}
