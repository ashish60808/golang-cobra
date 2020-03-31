package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
)

var username, password  string

var cmd = &cobra.Command{
	Use: "cobraintro",
	Short: "This tools gets a URL with basic auth",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {

			log.Fatalln("must specify the URL!")

		}
			client := http.Client{}

			req, err := http.NewRequest("GET", args[0],nil)

			if err != nil{
				log.Fatalln("unable to get  request")

			}
			if username != " " && password != " " {
				req.SetBasicAuth(username , password)
			}

			resp, err := client.Do(req)
            if err != nil{
				log.Fatalln("unable to get  response")
            }
            defer resp.Body.Close()

			content, err := ioutil.ReadAll(resp.Body)
			if err !=  nil{

				log.Fatalln("Unable to read body")
			}
			fmt.Println(string(content))

	},
}


func main() {
	cmd.PersistentFlags().StringVarP(&username,"username","u","","Username")
	cmd.PersistentFlags().StringVarP(&password,"password","p","","Password")
	cmd.Execute()
}

// curl --user foo:bar  https://httpbin.org/basic-auth/foo/bar
//  go run cobraintro_02.go -u foo -p bar  https://httpbin.org/basic-auth/foo/bar