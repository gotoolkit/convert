package convert_test

import (
	"testing"
	"github.com/gotoolkit/convert"
	"fmt"
	"log"
)

func TestClientVersion(t *testing.T)  {
	client, err := convert.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	v := client.Version()
	fmt.Println(v)
}


func TestPdfConvert(t *testing.T)  {
	client, err := convert.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	out, err := client.Src("sample.pdf").Digest("sample.jpg").Out()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
}