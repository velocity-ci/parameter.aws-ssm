package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func getParameterFromSSM(paramName string) (export *exported, expires time.Time, _ error) {
	sess := session.Must(session.NewSession(&aws.Config{
		MaxRetries: aws.Int(3),
	}))
	svc := ssm.New(sess)

	result, err := svc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(paramName),
		WithDecryption: aws.Bool(true),
	})

	if err != nil {
		return nil, expires, err
	}

	return &exported{
		Value: *result.Parameter.Value,
	}, time.Now().UTC().Add(1 * time.Hour), nil
}

func main() {
	paramName := flag.String("name", "", "AWS SSM Parameter Name to get value of")
	flag.Parse()

	out := output{
		Secret: true,
	}
	if creds, expires, err := getParameterFromSSM(*paramName); err == nil {
		out.Exports = creds
		out.Expires = expires
		out.State = "success"
	} else {
		out.State = "critical"
		out.Error = err.Error()
		out.Exports = &exported{}
	}
	o, _ := json.Marshal(out)

	fmt.Printf("%s", o)

}

type output struct {
	Secret  bool      `json:"secret"`
	Exports *exported `json:"exports"`
	Expires time.Time `json:"expires"`
	Error   string    `json:"error"`
	State   string    `json:"state"`
}

type exported struct {
	Value string `json:"value"`
}
