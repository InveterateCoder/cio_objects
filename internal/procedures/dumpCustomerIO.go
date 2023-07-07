package procedures

import (
	awshelp "cio_objects/internal/aws_help"
	"cio_objects/internal/helpers"
	"cio_objects/internal/paths"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
)

type dumpRet struct {
	Name  string `json:"name"`
	Links struct {
		File      string `json:"file"`
		Execution string `json:"execution"`
	} `json:"links"`
	Status struct {
		Metadata struct {
			HttpStatusCode  int    `json:"httpStatusCode"`
			RequestID       string `json:"requestId"`
			Attempts        int    `json:"attempts"`
			TotalRetryDelay int    `json:"totalRetryDelay"`
		} `json:"$metadata"`
		ExecutionArn string    `json:"executionArn"`
		StartDate    time.Time `json:"startDate"`
	} `json:"status"`
}

func dump_customerio(path *paths.Paths) {
	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("https://33mhwoi3jt6xzzntk2xohzq7fm0lmpnt.lambda-url.sa-east-1.on.aws/%s/customers/dump",
			strings.ToLower(string(path.Region))),
		strings.NewReader(`{ "filter": {} }`))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Custom-Auth", "WTg1OTnZ1qKomI1kPPQrAqOL9uPupSdk")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	var ret dumpRet
	if err = json.NewDecoder(res.Body).Decode(&ret); err != nil {
		panic(err)
	}
	res.Body.Close()
	if ret.Status.Metadata.HttpStatusCode != 200 {
		panic(fmt.Sprintf("%+v\n", ret))
	}
	fmt.Println("Created Dump SF execution")
	sf := awshelp.NewSF("arn:aws:states:sa-east-1:787732066160:stateMachine:DumpCustomersToS3FileStateMachine-m1CjbgzcK55Q")
	secs := 0
	for {
		desc := sf.DescribeExecution(ret.Status.ExecutionArn)
		if *desc.Status == "SUCCEEDED" {
			break
		}
		if *desc.Status != "RUNNING" {
			panic(fmt.Sprintf("%s\n%s\n", ret.Status.ExecutionArn, *desc.Status))
		}
		time.Sleep(5 * time.Second)
		secs += 5
		fmt.Printf("\rSeconds past %d", secs)
	}
	fmt.Println()
	s3 := awshelp.NewS3("cio-api-process-bucket")
	bts := make([]byte, 300_000_000)
	buf := aws.NewWriteAtBuffer(bts)
	var n int64
	if n, err = s3.Download(ret.Name, buf); err != nil {
		panic(err)
	}
	var vals []helpers.Customer
	if err = json.Unmarshal(bts[:n], &vals); err != nil {
		panic(err)
	}
	helpers.SaveJson(vals, path.CustomerioPath)
	if _, err = s3.Delete(ret.Name); err != nil {
		panic(err)
	}
}
