package main

import (
	_ "context"
	_ "encoding/json"
	"flag"
	_ "flag"
	"fmt"
	_ "fmt"
	"github.com/olivere/elastic"
	_ "github.com/olivere/elastic"
	"log"
	_ "log"
	_ "reflect"
	"time"
	_ "time"
)

type doc struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"@timestamp"`
}
type esDefaults struct {
	url   string
	index string
	typ   string
	sniff bool
}

var esDefs = &esDefaults{"http://127.0.0.1:9200", "cloudwatch", "log", false}

func main() {

	var (
		url   = flag.String("url", esDefs.url, "Elasticsearch URL")
		sniff = flag.Bool("sniff", esDefs.sniff, "Enable or disable sniffing")
		index = flag.String("index", esDefs.index, "Elasticsearch index name")
		typ   = flag.String("type", esDefs.typ, "Elasticsearch type name")
	)

	flag.Parse()
	log.SetFlags(0)

	fmt.Printf("\nInitiating connection with parameters:\n")
	fmt.Printf("  [Url: \"%s\", Sniff: \"%v\", Index: \"%s\", Type: \"%s\"]\n\n", *url, *sniff, *index, *typ)

	// Create Elasticsearch client
	client, err := elastic.NewClient(elastic.SetURL(*url), elastic.SetSniff(*sniff))
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Connection succeeded... ")
	}
	// Fetch Elasticsearch version
	esversion, err := client.ElasticsearchVersion(*url)
	if err != nil {
		// Handle error
		panic(err)
	} else {
		fmt.Printf("Elasticsearch version is: %s\n\n", esversion)
	}

}
