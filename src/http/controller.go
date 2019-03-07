package http

import "github.com/Madredix/theboats/src/models"

func Autocomplete(h Handler) {
	s := models.NewSearchRepo(h.GetDB())
	h.Ok(s.Autocomplite(h.GetParam(`q`)))
}

func Search(h Handler) {
	s := models.NewSearchRepo(h.GetDB())
	h.Ok(s.Search(h.GetParam(`q`)))
}

func Update(h Handler) {
	h.Ok(`not ready`)
}
