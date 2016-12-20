package convert_test

import (
	"testing"
	"github.com/gotoolkit/convert"
	"fmt"
	"log"
)

func TestClient(t *testing.T)  {
	client, err := convert.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	v := client.Version()
	fmt.Println(v)
}

