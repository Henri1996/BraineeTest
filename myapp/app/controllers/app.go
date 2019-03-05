package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	greeting := "Brainee"

	return c.Render(greeting)
}

func (c App) Hello(id int, text string, author string) revel.Result {
	return c.Render(text)
}
