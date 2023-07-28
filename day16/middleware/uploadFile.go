package middleware

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UploadFiles(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		file, err := c.FormFile("input-blog-image")

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		src , errFile := file.Open()

		if errFile != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		defer src.Close()
		
		tpmFile , errTmpFile := ioutil.TempFile("uploads","image-*.png")
		
		
		if errTmpFile != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		
		defer tpmFile.Close()

		writtenCopy, errCopy := io.Copy(tpmFile, src)
		
		if errCopy != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		fmt.Println("writtenCopy :", writtenCopy)

		data := tpmFile.Name()
		fmt.Println("nama File :", data)
		
		fileName := data[8:]
		
		fmt.Println("nama File terpotong :", fileName)

		c.Set("dataImage", fileName)

		return next(c)
	}
	
}