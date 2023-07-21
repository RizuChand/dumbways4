package main

import (
	"fmt"
	"html/template"
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Name      string
	StartDate string
	EndDate   string
	Textareas string
	DataTech  []string
}

var dataProject = []Project{
	{
		Name:      "user1",
		StartDate: "02/11/2022",
		EndDate:   "03/12/2022",
		Textareas: "Percobaan user satu",
		DataTech:  []string{"nodeJS", "VueJS"},
	},
	{
		Name:      "user2",
		StartDate: "02/11/2022",
		EndDate:   "03/12/2022",
		Textareas: "Percobaan user sdua",
		DataTech:  []string{"nodeJS", "VueJS"},
	},
}

func main() {
	e := echo.New()

	e.Static("/css", "css")
	e.Static("/img", "img")
	e.Static("/javascript", "javascript")

	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/myproject", myProject)
	e.GET("/testimonial", testimonial)

	e.GET("/detailproject/:id", detailProject)

	e.POST("/addmyproject", addMyproject)

	e.POST("/deleteproject/:id", deleteProject)

	e.GET("/updateproject/:id", updateProject)
	e.POST("updateprojectform/:id", updateProjectForm)

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

	data := map[string]interface{}{
		"projects": dataProject,
	}

	return template.Execute(c.Response(), data)
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

	nodejs := c.FormValue("nodejs")
	nextjs := c.FormValue("nextjs")
	reactjs := c.FormValue("reactjs")
	typescript := c.FormValue("typescript")

	dataCheck := []string{nodejs, nextjs, reactjs, typescript}

	newProject := Project{
		Name:      name,
		StartDate: startDate,
		EndDate:   endDate,
		Textareas: textareas,
		DataTech:  dataCheck,
	}

	dataProject = append(dataProject, newProject)

	fmt.Println(dataProject)

	return c.Redirect(http.StatusMovedPermanently, "/myproject")
}

func detailProject(c echo.Context) error {

	id := c.Param("id")

	template, err := template.ParseFiles("./detailProject.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	idToInt, _ := strconv.Atoi(id)

	projectDetail := Project{}

	for index, data := range dataProject {
		if index == idToInt {
			projectDetail = Project{
				Name:      data.Name,
				StartDate: data.StartDate,
				EndDate:   data.EndDate,
				Textareas: data.Textareas,
				DataTech:  data.DataTech,
			}

		}
	}

	apalah := map[string]interface{}{
		"ID":      id,
		"project": projectDetail,
	}

	return template.Execute(c.Response(), apalah)
}

func deleteProject(c echo.Context) error {
	id := c.Param("id")
	idToInt, _ := strconv.Atoi(id)

	dataProject = append(dataProject[:idToInt], dataProject[idToInt+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/myproject")
}

func updateProject(c echo.Context) error {
	id := c.Param("id")

	template, _ := template.ParseFiles("./updateProject.html")

	dataID := map[string]interface{}{
		"ID": id,
	}
	return template.Execute(c.Response(), dataID)
}

func updateProjectForm(c echo.Context) error {
	id := c.Param("id")
	idToInt, _ := strconv.Atoi(id)

	name := c.FormValue("name")
	startDate := c.FormValue("start-date")
	endDate := c.FormValue("end-date")
	textareas := c.FormValue("text-area")

	nodejs := c.FormValue("nodejs")
	nextjs := c.FormValue("nextjs")
	reactjs := c.FormValue("reactjs")
	typescript := c.FormValue("typescript")

	dataCheck := []string{nodejs, nextjs, reactjs, typescript}

	dataProject[idToInt].Name = name
	dataProject[idToInt].StartDate = startDate
	dataProject[idToInt].EndDate = endDate
	dataProject[idToInt].Textareas = textareas
	dataProject[idToInt].DataTech = dataCheck

	fmt.Println(dataProject[idToInt])

	return c.Redirect(http.StatusMovedPermanently, "/myproject")

}
