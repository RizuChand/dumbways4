package main

import (
	"context"
	"day11/config"
	"day11/middleware"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Project struct {
	ID int
	Name      string
	StartDate time.Time
	EndDate   time.Time
	Textareas string
	DataTech  []string
	Image string
	Username string
	AuthorId int
}

type Users struct {
	Id int
	Name string
	Email string
	Password string
}

// type UserLoginSession struct {
// 	isLogin bool
// 	Name string
// }


func main() {

	config.DatabaseConnection();
	e := echo.New()

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("capjari"))))

	
	e.Static("/css", "css")
	e.Static("/img", "img")
	e.Static("/javascript", "javascript")
	e.Static("uploads","uploads")

	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/myproject", myProject)
	e.GET("/testimonial", testimonial)

	e.GET("/detailproject/:id", detailProject)

	e.POST("/addmyproject", middleware.UploadFiles(addMyproject))

	e.POST("/deleteproject/:id", deleteProject)

	e.GET("/updateproject/:id",updateProject)

	e.POST("/updateprojectform",middleware.UploadFiles(updateProjectForm))

	e.GET("/form-register", formRegister)

	e.POST("/register", register)
	
	e.GET("/form-login", formLogin)

	e.POST("/login", login)

	e.POST("/logout",logout)
	

	e.Logger.Fatal(e.Start("localhost:8000"))

	
	
}

func home(c echo.Context) error {
	template, err := template.ParseFiles("./views/index.html")
	sess, errSess := session.Get("bersesion",c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	dataProject, errProject := config.Conn.Query(context.Background(),"SELECT tb_projects.id, tb_projects.name, tb_projects.start_date, tb_projects.end_date, tb_projects.description, tb_projects.technologies, tb_projects.image, tb_projects.author_id, tb_users.name FROM public.tb_projects INNER JOIN tb_users on tb_projects.author_id = tb_users.id")

	if errProject != nil {
		return c.JSON(http.StatusInternalServerError, errProject.Error())
	}

	var newDataProject [] Project

	for dataProject.Next() {
		var bucket = Project{}

		err := dataProject.Scan(&bucket.ID, &bucket.Name, &bucket.StartDate, &bucket.EndDate, &bucket.Textareas, &bucket.DataTech, &bucket.Image, &bucket.AuthorId, 
			&bucket.Username)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}



		newDataProject = append(newDataProject, bucket)


	}


	dataFlash := map[string]interface{}{
		"projects" : newDataProject,
		"flashMessage" : sess.Values["message"],
		"test" : sess.Values["name"],

	}

	delete(sess.Values, "message")
	sess.Save(c.Request(), c.Response())

	return template.Execute(c.Response(), dataFlash)
}
func contact(c echo.Context) error {
	template, err := template.ParseFiles("./views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return template.Execute(c.Response(), nil)
}

func myProject(c echo.Context) error {
	template, err := template.ParseFiles("./views/addProject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	sess, errSess := session.Get("bersesion",c)

	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	user := sess.Values["id"]


	dataProject, errProject := config.Conn.Query(context.Background(),"SELECT tb_projects.id, tb_projects.name, tb_projects.description, tb_projects.technologies, tb_projects.image, tb_projects.author_id, tb_users.name FROM public.tb_projects INNER JOIN tb_users on tb_projects.author_id = tb_users.id WHERE tb_projects.author_id =$1", user)

	if errProject != nil {
		return c.JSON(http.StatusInternalServerError, errProject.Error())
	}

	var newDataProject [] Project

	for dataProject.Next() {
		var bucket = Project{}

		
		err := dataProject.Scan(&bucket.ID, &bucket.Name, &bucket.Textareas, &bucket.DataTech, &bucket.Image, &bucket.AuthorId, &bucket.Username )
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		newDataProject = append(newDataProject, bucket)

	}

	dataFlash := map[string]interface{}{
		"flashMessage" : sess.Values["message"],
		"test" : sess.Values["name"],

	}

	data := map[string]interface{}{
		"projects": newDataProject,
		"dataFlash" : dataFlash,
	}

	

	return template.Execute(c.Response(), data)
}


func testimonial(c echo.Context) error {
	template, err := template.ParseFiles("views/testimonial.html")
	sess, errSess := session.Get("bersesion", c)

	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	dataFlash := map[string]interface{}{
		"flashMessage" : sess.Values["message"],
		"test" : sess.Values["name"],

	}
	return template.Execute(c.Response(), dataFlash)
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

	sess, err := session.Get("bersesion",c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,err.Error())
	}

	authorId := sess.Values["id"]


	
	imageDefault := c.Get("dataImage").(string)
	
	dataCheck := []string{nodejs, javascript, reactjs, typescript}
		
	dataAdd, err := config.Conn.Exec(context.Background(),"INSERT INTO tb_projects (name, start_date, end_date, description, technologies, image, author_id) VALUES ($1, $2, $3, $4, $5,$6,$7)", name, startDate, endDate,textareas,dataCheck,imageDefault,authorId)
	
	fmt.Println("row affected:", dataAdd.RowsAffected())

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	
	return c.Redirect(http.StatusMovedPermanently, "/myproject")
}


func getDataById(id int)(Project, error)  {
	
	var data = Project{}
	

	// query := "SELECT id, name, start_date, end_date, description, technologies, image FROM tb_projects WHERE id=$1"

	query2 := "SELECT tb_projects.id, tb_projects.name, tb_projects.start_date, tb_projects.end_date, tb_projects.description, tb_projects.technologies, tb_projects.image, tb_projects.author_id, tb_users.name FROM public.tb_projects INNER JOIN tb_users on tb_projects.author_id = tb_users.id WHERE tb_projects.id=$1"

	dataProject := config.Conn.QueryRow(context.Background(),query2, id)

	
	err := dataProject.Scan(&data.ID, &data.Name,&data.StartDate, &data.EndDate, &data.Textareas, &data.DataTech, &data.Image, &data.AuthorId, &data.Username)
	
		

	return data, err
}


func detailProject(c echo.Context) error {

	id := c.Param("id")

	idToInt, _ := strconv.Atoi(id)

	template, err := template.ParseFiles("./views/detailProject2.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	
	sess, errSess := session.Get("bersesion", c)
	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	data, err := getDataById(idToInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	
	dataFlash := map[string]interface{}{
		"flashMessage" : sess.Values["message"],
		"test" : sess.Values["name"],

	}


	data2 := map[string]interface{}{
		"projects": data,
		"dataFlash" : dataFlash,
	}


	return template.Execute(c.Response(), data2)
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

	template, err := template.ParseFiles("./views/updateProject.html")

	if err != nil  {
		c.JSON(http.StatusBadRequest,err.Error())
	}

	dataUpdate := map[string]interface{}{
		"ID" : id,
	}

	
	return template.Execute(c.Response(), dataUpdate)
}

func updateProjectForm(c echo.Context) error {
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

	imageDefault := c.Get("dataImage").(string)


	dataUpdate, err := config.Conn.Exec(context.Background(),"UPDATE tb_projects SET name=$1, start_date=$2, end_date=$3, description=$4, technologies=$5, image=$6 WHERE id=$7", name, startDate, endDate,textareas,dataCheck,imageDefault,ID)
	
	fmt.Println("row affected:", dataUpdate.RowsAffected())

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}


	return c.Redirect(http.StatusMovedPermanently, "/myproject")

}


func formRegister(c echo.Context) error {
	tmpl, err := template.ParseFiles("./views/formRegister.html")
	sess, sessErr := session.Get("bersesion", c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	if sessErr != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	dataFlash := map[string]interface{}{
		"flashMessage" : sess.Values["message"],

	}

	return tmpl.Execute(c.Response(),dataFlash)
}

func register(c echo.Context) error {
	name := c.FormValue("Name")
	email := c.FormValue("Email")
	password := c.FormValue("Password")


	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password),10)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	fmt.Println(name, email,password)

	query, errQuery := config.Conn.Exec(context.Background(),"INSERT INTO tb_users (name, email, password) VALUES($1, $2, $3)", name, email, hashedPass)

	fmt.Println("affected row : ", query.RowsAffected())

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message" : "register fail",
		})
	}

	return redirectWithMessage(c, "register succsess silahkan login dibawah", "/form-login")
	
	
}

func formLogin(c echo.Context) error {
	tmpl, err := template.ParseFiles("./views/formLogin.html")
	
	sess, errSess := session.Get("bersesion", c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	if errSess != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	dataFlash := map[string]interface{}{
		"flashMessage" : sess.Values["message"],

	}

	delete(sess.Values, "message")
	sess.Save(c.Request(), c.Response())

	return tmpl.Execute(c.Response(),dataFlash)
}


func login(c echo.Context) error {

	email := c.FormValue("Email")
	password := c.FormValue("Password")

	user := Users{} 

	//checked Email is exist in DB

	err := config.Conn.QueryRow(context.Background(),"SELECT id, name, email, password FROM tb_users WHERE email=$1",email).Scan(&user.Id,&user.Name,&user.Email,&user.Password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message" : "gagal login",
		})
	}

	errPass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	
	if errPass != nil {
		return redirectWithMessage(c, "login failed", "/form-login")
	}

	sess, errSess := session.Get("bersesion", c)
	if errSess != nil {
		return c.JSON(http.StatusInternalServerError,errSess.Error())
	}

	sess.Options.MaxAge = 10800 
	sess.Values["name"] = user.Name
	sess.Values["email"] = user.Email
	sess.Values["id"] = user.Id
	
	sess.Save(c.Request(), c.Response())


	
	return redirectWithMessage(c, "berhasil login bos", "/")
	
}

func logout(c echo.Context) error {

	sess, err := session.Get("bersesion", c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,err.Error())
	}
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	return redirectWithMessage(c, "Logout berhasil!", "/")
}


//fungsi ngeredirect dan pesan


func redirectWithMessage(c echo.Context, message string, redirectUrl string) error {
	sess, errSess := session.Get("bersesion", c)

	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	sess.Values["message"] = message
	sess.Save(c.Request(), c.Response())

	fmt.Println("Pesan :",sess.Values["message"])

	return c.Redirect(http.StatusMovedPermanently, redirectUrl)
}