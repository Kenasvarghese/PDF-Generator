package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Kenasvarghese/PDF-Generator/config"
	"github.com/Kenasvarghese/PDF-Generator/pdfgenerator"
	"github.com/Kenasvarghese/PDF-Generator/utils"
	"github.com/go-chi/chi/v5"
)

func main() {
	config.LoadConfig()
	err := utils.Validate(config.Configuration)
	if err != nil {
		panic(err)
	}
	r := chi.NewRouter()
	router := chi.NewRouter()
	router.Post("/generate", handler)
	r.Mount(config.Configuration.BasePath, router)
	server := &http.Server{Addr: ":" + config.Configuration.Port, Handler: r}
	server.ListenAndServe()

}

func handler(w http.ResponseWriter, r *http.Request) {
	detailsMap := make(map[string]any, 0)
	err := json.NewDecoder(r.Body).Decode(&detailsMap)
	if err != nil {
		log.Print(fmt.Sprintf("json.NewDecoder returned err :%v", err))
		utils.WriteApiResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	generator := pdfgenerator.NewPDF(config.Configuration.BrowserURL)
	err = generator.GenerateHtml(pdfgenerator.SampleTemplate, detailsMap)
	if err != nil {
		log.Print(fmt.Sprintf("generator.GenerateHtml returned err :%v", err))
		utils.WriteApiResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	bytes, err := generator.GeneratePDF(context.Background())
	if err != nil {
		log.Print(fmt.Sprintf("generator.GeneratePDF returned err :%v", err))
		utils.WriteApiResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	fo, err := os.Create("output.pdf")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	fo.Write(bytes)
	utils.WriteApiResponse(w, http.StatusOK, "pdf successfully generated")
	log.Print("Success")
	return
}
