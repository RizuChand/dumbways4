package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/css", "css")
	e.Static("/img", "img")
	e.Static("/javascript", "javascript")

	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/myproject", myProject)
	e.GET("/testimonial", testimonial)

	e.POST("/addmyproject", addMyproject)

	fmt.Println("hai bro")

	e.Logger.Fatal(e.Start("localhost:8080"))

}

func home(c echo.Context) error {
	template, err := template.ParseFiles("./index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return template.Execute(c.Response(), nil)
}
func contact(c echo.Context) error {
	template, err := template.ParseFiles("./contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return template.Execute(c.Response(), nil)
}

func myProject(c echo.Context) error {
	template, err := template.ParseFiles("./addProject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return template.Execute(c.Response(), nil)
}

func testimonial(c echo.Context) error {
	template, err := template.ParseFiles("./testimonial.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return template.Execute(c.Response(), nil)
}

func addMyproject(c echo.Context) error {
	name := c.FormValue("name")
	startDate := c.FormValue("start-date")
	endDate := c.FormValue("end-date")
	textareas := c.FormValue("text-area")

	fmt.Println("nama :", name)
	fmt.Println("start-date :", startDate)
	fmt.Println("end-date :", endDate)
	fmt.Println("text area :", textareas)

	return c.Redirect(http.StatusMovedPermanently, "/")
}
