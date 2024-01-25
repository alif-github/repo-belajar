package main

import (
	"context"
	"fmt"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"github.com/Azure/azure-storage-blob-go/azblob"
	"io/ioutil"
	"log"
	"net/url"
	"os"
)

var azurePipeline pipeline.Pipeline

func main() {
	args := os.Args
	connection()
	switch args[0] {

	}
	errorS := upload()
	if errorS != nil {
		log.Println(fmt.Sprintf(`error msg : `), errorS.Error())
	}
}

func connection() {
	credential, _ := azblob.NewSharedKeyCredential(fmt.Sprintf(`nextest`), fmt.Sprintf(`2ltj8uR1Xu1nEwf2YNbN7lIsYrsIYCOcaWQjST2lTS+FQCLEUqwejdQt7gM2wCTqwhcTTx9IeLaI+AStKPmm6Q==`))
	azurePipeline = azblob.NewPipeline(credential, azblob.PipelineOptions{})
}

func upload() (errorS error) {
	var fileByte []byte

	fileByte, errorS = ioutil.ReadFile(fmt.Sprintf(`C:\repo-belajar\repo-eksperimen-code\src\azure-import\file\sesi.pdf`))
	if errorS != nil {
		return
	}

	host := fmt.Sprintf(`https://nextest.blob.core.windows.net`)
	suffix := fmt.Sprintf(`/nextestdisk/nextrac/`)
	URL, _ := url.Parse(host + suffix)
	containerURL := azblob.NewContainerURL(*URL, azurePipeline)

	ctx := context.Background()

	path := fmt.Sprintf(`import/todo_list/done/2022/07/21/sesi1_9999994e33cb41d0a47144b16c0a441d.pdf`)

	blobURL := containerURL.NewBlockBlobURL(path)
	_, errorS = azblob.UploadBufferToBlockBlob(ctx, fileByte, blobURL, azblob.UploadToBlockBlobOptions{BlockSize: 4 * 1024 * 1024, Parallelism: 16})
	if errorS != nil {
		return
	}

	log.Println("link : ", host+suffix+path)
	return nil
}
