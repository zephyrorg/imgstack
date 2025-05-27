package main

import(
	"log"

	"github.com/hellflame/argparse"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)


func main() {
	parser := argparse.NewParser("imgstack-server", "this is a basic program", nil)

	endpoint := parser.String("e", "endpoint", nil)
	accessKeyID := parser.String("a", "accesskey", nil)
	secretAccessKey := parser.String("s", "secretkey", nil)


	if e := parser.Parse(nil); e != nil {
		log.Println(e.Error())
		return
	}

	log.Println("Endpoint:  ", *endpoint)
	log.Println("AccessKey: ", *accessKeyID)
	log.Println("SecretKey: ", *secretAccessKey)
}

