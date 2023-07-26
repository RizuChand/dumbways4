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

	e.POST("/updateprojectform", updateProjectForm)

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
	javascript := c.FormValue("javascript")
	reactjs := c.FormValue("reactjs")
	typescript := c.FormValue("typescript")
	
	imageDefault := "https://cdn.shopify.com/s/files/1/0003/9573/9145/files/2_large.jpg?v=1552901583"
	
	dataCheck := []string{nodejs, javascript, reactjs, typescript}
		
	dataAdd, err := config.Conn.Exec(context.Background(),"INSERT INTO tb_projects (name, start_date, end_date, description, technologies, image) VALUES ($1, $2, $3, $4,$5,$6)", name, startDate, endDate,textareas,dataCheck,imageDefault)
	
	fmt.Println("row affected:", dataAdd.RowsAffected())

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	
	return c.Redirect(http.StatusMovedPermanently, "/myproject")
}


func getDataById(id int)(Project, error)  {
	
	var data = Project{}
	

	query := "SELECT id, name, description, technologies, image FROM tb_projects WHERE id=$1"

	dataProject := config.Conn.QueryRow(context.Background(),query, id)

	
	err := dataProject.Scan(&data.ID, &data.Name,&data.Textareas, &data.DataTech, &data.Image )
	
		

	return data, err
}


func detailProject(c echo.Context) error {

	id := c.Param("id")

	idToInt, _ := strconv.Atoi(id)

	template, err := template.ParseFiles("./detailProject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	

	data, err := getDataById(idToInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	fmt.Println(data)

	return template.Execute(c.Response(), data)
}


func deleteProject(c echo.Context) error {
	id := c.Param("id")
	idToInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	query := "DELETE FROM tb_projects WHERE id=$1"

	config.Conn.Exec(context.Background(),query,idToInt)


	return c.Redirect(http.StatusMovedPermanently, "/myproject")
}



func updateProject(c echo.Context) error {
	id := c.Param("id")

	

	template, _ := template.ParseFiles("./updateProject.html")


	dataUpdate := map[string]interface{}{
		"ID" : id,
	}

	
	return template.Execute(c.Response(), dataUpdate)
}

func updateProjectForm(c echo.Context) error {
	// id := c.Param("id")
	// idToInt, errorAtoi := strconv.Atoi(id)

	// if errorAtoi != nil {
	// 	c.JSON(http.StatusInternalServerError, errorAtoi.Error())
	// }
	ID := c.FormValue("ID")
	name := c.FormValue("name")
	startDate := c.FormValue("start-date")
	endDate := c.FormValue("end-date")
	textareas := c.FormValue("text-area")

	nodejs := c.FormValue("nodejs")
	javascript := c.FormValue("javascript")
	reactjs := c.FormValue("reactjs")
	typescript := c.FormValue("typescript")

	dataCheck := []string{nodejs, javascript, reactjs, typescript}

	imageDefault := "https://cdn.shopify.com/s/files/1/0003/9573/9145/files/2_large.jpg?v=1552901583"


	dataUpdate, err := config.Conn.Exec(context.Background(),"UPDATE tb_projects SET name=$1, start_date=$2, end_date=$3, description=$4, technologies=$5, image=$6 WHERE id=$7", name, startDate, endDate,textareas,dataCheck,imageDefault,ID)
	
	fmt.Println("row affected:", dataUpdate.RowsAffected())

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}


	return c.Redirect(http.StatusMovedPermanently, "/myproject")

}
