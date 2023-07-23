package main

import (
	"context"
	"day11/config"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Project struct {
	ID int
	Name      string
	StartDate string
	EndDate   string
	Textareas string
	DataTech  []string
	Image string
}

var dataProject = []Project{
	{
		Name:      "user1",
		StartDate: "02/11/2022",
		EndDate:   "03/12/2022",
		Textareas: "Percobaan user satu",
		DataTech:  []string{"nodeJS", "VueJS"},
		Image: "percobaan.jpg",
	},
	{
		Name:      "user2",
		StartDate: "02/11/2022",
		EndDate:   "03/12/2022",
		Textareas: "Percobaan user sdua",
		DataTech:  []string{"nodeJS", "VueJS"},
		Image: "percobaan.jpg",
	},
}

func main() {

	config.DatabaseConnection();
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

	e.Logger.Fatal(e.Start("localhost:8000"))

	
	
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


	dataProject, errProject := config.Conn.Query(context.Background(),"select id, name, description, technologies, image FROM tb_projects")

	if errProject != nil {
		return c.JSON(http.StatusInternalServerError, errProject.Error())
	}

	var newDataProject [] Project

	for dataProject.Next() {
		var bucket = Project{}



		err := dataProject.Scan(&bucket.ID, &bucket.Name, &bucket.Textareas, &bucket.DataTech, &bucket.Image )
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}


		newDataProject = append(newDataProject, bucket)


	}

	data := map[string]interface{}{
		"projects": newDataProject,
	}

	

	return template.Execute(c.Response(), data)
}



// func convertDate(EndTime, StartTime string) string {
// 	timeEnd, err := time.Parse(time.DateOnly, EndTime)
// 	if err != nil {
// 		return "Invalid EndTime format"
// 	}

// 	timeStart, err := time.Parse(time.DateOnly, StartTime)
// 	if err != nil {
// 		return "Invalid StartTime format"
// 	}

// 	timeDistance := timeEnd.Sub(timeStart)

// 	distanceSecond := int(timeDistance.Seconds())
// 	distanceMinute := distanceSecond / 60
// 	distanceHour := distanceMinute / 60
// 	distanceDay := distanceHour / 24
// 	distanceWeek := distanceDay / 7
// 	distanceMonth := distanceDay / 30
// 	distanceYear := distanceMonth / 12

// 	if distanceHour >= 24 && distanceDay <= 7 {
// 		return fmt.Sprintf("%d day of durations", distanceDay)
// 	} else if distanceDay >= 7 && distanceWeek <= 4 {
// 		return fmt.Sprintf("%d Weeks of durations", distanceWeek)
// 	} else if distanceWeek >= 4 && distanceMonth <= 12 {
// 		return fmt.Sprintf("%d months of durations", distanceMonth)
// 	} else if distanceMonth >= 12 {
// 		return fmt.Sprintf("%d years of durations", distanceYear)
// 	}

// 	return "Invalid duration"
// }





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

	image := c.FormValue("image")

	dataCheck := []string{nodejs, nextjs, reactjs, typescript}

	newProject := Project{
		Name:      name,
		StartDate: startDate,
		EndDate:   endDate,
		Textareas: textareas,
		DataTech:  dataCheck,
		Image : image,
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
				Image: data.Image,
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
