package pages

import (
	"net/http"

	"GoTenancy/libs/render"
)

// Controller home controller
type Controller struct {
	View *render.Render
}

// Index home index page
func (ctrl Controller) Index(w http.ResponseWriter, req *http.Request) {
	ctrl.View.Execute("index", map[string]interface{}{}, req, w)
}
