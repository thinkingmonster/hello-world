package handlers

import (
	"github.com/thinkingmonster/hello-world/pkg/render"
	"net/http"
)

// Home Provides the home page for the application
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")

}

//About provides the about page for the application
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
