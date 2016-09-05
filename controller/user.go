package controller

import (
	"fmt"
	"net/http"
)

// User is the controller for all /user endpoints.
type User struct{}

// Routes describes all endpoints handled by the User Controller.
func (user *User) Routes() []Route {
	return []Route{
		Route{
			"/user",
			[]Handler{
				Handler{"post", user.create},
				Handler{"get", user.find},
			},
		},

		Route{
			"/user/",
			[]Handler{
				Handler{"get", user.findOne},
				Handler{"patch", user.update},
				Handler{"delete", user.delete},
			},
		},
	}
}

/*
Handlers
*/

func (user *User) create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create")
}

func (user *User) find(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "find")
}

func (user *User) findOne(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "findOne")
}

func (user *User) update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "update")
}

func (user *User) delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "delete")
}
