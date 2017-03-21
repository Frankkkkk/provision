package main

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/ghodss/yaml"

	"github.com/VictorLowther/jsonpatch/utils"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	apiclient "github.com/rackn/rocket-skates/client"
	"github.com/spf13/cobra"
)

var (
	version            = "1.1.1"
	debug              = false
	endpoint           = "https://127.0.0.1:8092"
	username, password string
	format             = "json"
	app                = &cobra.Command{
		Use:   "rscli",
		Short: "A CLI application for interacting with the Rocket-Skates API",
	}
	session   *apiclient.RocketSkates
	basicAuth runtime.ClientAuthInfoWriter
)

func safeMergeJSON(target, toMerge []byte) ([]byte, error) {
	targetObj := make(map[string]interface{})
	toMergeObj := make(map[string]interface{})
	if err := json.Unmarshal(target, &targetObj); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(toMerge, &toMergeObj); err != nil {
		return nil, err
	}
	outObj, ok := utils.Merge(targetObj, toMergeObj).(map[string]interface{})
	if !ok {
		return nil, errors.New("Cannot happen in safeMergeJSON")
	}
	keys := make([]string, 0)
	for k := range outObj {
		if _, ok := targetObj[k]; !ok {
			keys = append(keys, k)
		}
	}
	for _, k := range keys {
		delete(outObj, k)
	}
	return json.Marshal(outObj)
}

func d(msg string, args ...interface{}) {
	if debug {
		log.Printf(msg, args...)
	}
}

func pretty(o interface{}) (res string) {
	var buf []byte
	var err error
	switch format {
	case "json":
		buf, err = json.MarshalIndent(o, "", "  ")
	case "yaml":
		buf, err = yaml.Marshal(o)
	default:
		log.Fatalf("Unknown pretty format %s", format)
	}
	if err != nil {
		log.Fatalf("Failed to unmarshal returned object!")
	}
	return string(buf)
}

func init() {
	if ep := os.Getenv("RS_ENDPOINT"); ep != "" {
		endpoint = ep
	}
	if kv := os.Getenv("RS_KEY"); kv != "" {
		key := strings.SplitN(kv, ":", 2)
		if len(key) < 2 {
			log.Fatal("RS_KEY does not contain a username:password pair!")
		}
		if key[0] == "" || key[1] == "" {
			log.Fatal("RS_KEY contains an invalid username:password pair!")
		}
		username = key[0]
		password = key[1]
	}
	app.PersistentFlags().StringVarP(&endpoint,
		"endpoint", "E", endpoint,
		"The Rocket-Skates API endpoint to talk to")
	app.PersistentFlags().StringVarP(&username,
		"username", "U", username,
		"Name of the Rocket-Skates user to talk to")
	app.PersistentFlags().StringVarP(&password,
		"password", "P", password,
		"Password of the Rocket-Skates user")
	app.PersistentFlags().BoolVarP(&debug,
		"debug", "d", false,
		"Whether the CLI should run in debug mode")
	app.PersistentFlags().StringVarP(&format,
		"format", "F", "json",
		`The serialzation we expect for output.  Can be "json" or "yaml"`)
}

func main() {
	app.PersistentPreRun = func(c *cobra.Command, a []string) {
		var err error
		d("Talking to Rocket-Skates with %v (%v:%v)", endpoint, username, password)
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		hc := &http.Client{Transport: tr}
		epURL, err := url.Parse(endpoint)
		if err != nil {
			log.Fatalf("Error handling endpoint %s: %v", endpoint, err)
		}
		transport := httptransport.NewWithClient(epURL.Host, "/api/v3", []string{epURL.Scheme}, hc)
		session = apiclient.New(transport, strfmt.Default)
		basicAuth = httptransport.BasicAuth(username, password)

		if err != nil {
			if c.Use != "version" {
				log.Fatalf("Could not connect to Rocket-Skates: %v\n", err.Error())
			}
		}
	}
	app.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Rocket-Skates CLI Command Version",
		Run: func(cmd *cobra.Command, args []string) {
			log.Printf("Version: %v", version)
		},
	})
	app.Execute()
}
