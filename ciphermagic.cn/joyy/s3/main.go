package main

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"log"
	"os"
)

const bucket = "drc-cdc-data-log-by-fs"

var sess *session.Session
var svc *s3.S3

func init() {
	accessKey := "xxx"
	secretKey := "xxx+Uk4"
	endPoint := "xxx" //endpoint设置，不要动
	_sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Endpoint:    aws.String(endPoint),
		Region:      aws.String("xxx"),
	})
	if err != nil {
		panic(err)
	}
	sess = _sess
	svc = s3.New(_sess)
}

func main() {
	//ListObjects()
	//Upload()
	//GetObject()
	SelectObject()
	//DeleteObject()
}

func Upload() {
	filename := "/var/folders/2g/ysdkpd6j4d1f060h48qq7mbw0000gn/T/local/testId/testDb_testTable_0.log"
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	uploader := s3manager.NewUploader(sess)
	output, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String("/archive/" + filename),
		Body:   file,
	})
	if err != nil {
		log.Fatalf("Unable to upload %q to %q, %v", filename, bucket, err)
	}
	log.Printf("Successfully uploaded %q to %q, output: %+v\n", filename, bucket, output)
}

func GetObject() {
	filename := "archive/local/testId/testDb/testTable/0.log"
	object, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		log.Fatalf("Unable to get %q to %q, %v", filename, bucket, err)
	}
	defer object.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(object.Body)
	data := buf.String()
	log.Printf("output: %s\n", data)
}

func ListObjects() {
	result, err := svc.ListObjects(&s3.ListObjectsInput{
		Prefix: aws.String("archive/local/testId/testDb/testTable/"),
		Bucket: aws.String(bucket),
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range result.Contents {
		log.Println("Name:          ", *item.Key)
		log.Println("Size:          ", *item.Size)
		log.Println("Last modified: ", *item.LastModified)
		log.Println("Storage class: ", *item.StorageClass)
		fmt.Println("--------------------------------------------")
	}
}

func DeleteObject() {
	filename := "archive/local/testId/testDb/testTable/batch_0.log"
	output, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		log.Fatalf("Unable to delete %q to %q, %v", filename, bucket, err)
	}
	log.Printf("Successfully delete %q to %q, output: %+v\n", filename, bucket, output)
}

func SelectObject() {
	filename := "archive/local/testId/testDb/testTable/0.log"
	resp, err := svc.SelectObjectContent(&s3.SelectObjectContentInput{
		Bucket:         aws.String(bucket),
		Key:            aws.String(filename),
		Expression:     aws.String("SELECT S3Object.name FROM S3Object WHERE S3Object.id = '94'"),
		ExpressionType: aws.String("SQL"),
		InputSerialization: &s3.InputSerialization{
			JSON: &s3.JSONInput{
				Type: aws.String("Document"),
			},
		},
		OutputSerialization: &s3.OutputSerialization{
			JSON: &s3.JSONOutput{},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer resp.EventStream.Close()
	results, resultWriter := io.Pipe()
	go func() {
		defer resultWriter.Close()
		for event := range resp.EventStream.Events() {
			switch e := event.(type) {
			case *s3.RecordsEvent:
				resultWriter.Write(e.Payload)
			case *s3.StatsEvent:
				fmt.Printf("Processed %d bytes\n", *e.Details.BytesProcessed)
			}
		}
	}()
	buf := new(bytes.Buffer)
	buf.ReadFrom(results)
	data := buf.String()
	log.Printf("output: %s\n", data)
	log.Printf("size: %d\n", len(buf.Bytes())/1024/1024)
}
