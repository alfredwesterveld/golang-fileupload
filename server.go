package main

import (
	"fmt"
	"io"
	"os"
	"net/http"
	"github.com/GeertJohan/go.rice"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func upload(c echo.Context) error {
	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create("public/" + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully</p>", file.Filename))
}

func main() {
	addr := ":1323"

	if os.Getenv("ADDRESS") != "" {
		addr = os.Getenv("ADDRESS")
	}


	e := echo.New()

	assetHandler := http.FileServer(rice.MustFindBox("public").HTTPBox())

	e.Get("/", standard.WrapHandler(assetHandler))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Static("public"))
	e.Post("/upload", upload)

	e.Run(standard.New(addr))

	fmt.Printf("Server running at", addr)
}
