package main

import (
	"context"
	"log"

	// "github.com/iki-rumondor/init-golang-service/internal/adapter/database"
	// customHTTP "github.com/iki-rumondor/init-golang-service/internal/adapter/http"
	// "github.com/iki-rumondor/init-golang-service/internal/application"
	// "github.com/iki-rumondor/init-golang-service/internal/repository"
	// "github.com/iki-rumondor/init-golang-service/internal/routes"
	aai "github.com/AssemblyAI/assemblyai-go-sdk"
)

func main() {
	// gormDB, err := database.NewMysqlDB()
	// if err != nil{
	// 	log.Fatal(err.Error())
	// 	return
	// }

	// repo := repository.NewRepository(gormDB)
	// service := application.NewService(repo)
	// handler := customHTTP.NewHandler(service)

	// var PORT = ":8080"
	// routes.StartServer(handler).Run(PORT)

	ctx := context.Background()

	client := aai.NewClient("d0f941a14eb64805ae9926b011792b2c")

	// audioURL := "https://github.com/AssemblyAI-Examples/audio-examples/raw/main/20230607_me_canadian_wildfires.mp3"
	audioURL := "https://www.youtube.com/watch?v=J-0su0hQgHE&list=RDMMJ-0su0hQgHE&start_radio=1"

	transcript, _ := client.Transcripts.TranscribeFromURL(ctx, audioURL, nil)

	// transcript, _ := client.Transcripts.TranscribeFromURL(ctx, audioURL, &aai.TranscriptOptionalParams{
	// 	DualChannel: aai.Bool(true),
	// })

	// params := &aai.TranscriptGetSubtitlesOptions{
	// 	CharsPerCaption: 32,
	// }

	// vtt, _ := client.Transcripts.GetSubtitles(ctx, aai.ToString(transcript.ID), "vtt", params)

	// file, err := os.Create("output.vtt")
	// if err != nil {
	// 	log.Fatalf("Failed to create file: %v", err)
	// }
	// defer file.Close()

	// _, err = file.WriteString(string(vtt))
	// if err != nil {
	// 	log.Fatalf("Failed to write to file: %v", err)
	// }

	// fmt.Println("VTT file saved successfully")

	log.Println(aai.ToString(transcript.Text))
}
