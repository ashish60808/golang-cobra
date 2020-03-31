package main

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var username, password , content string

var cmd = &cobra.Command{
	Use: "cobraintro",
	Short: "This tools gets a URL with basic auth",
	Run: func(cmd *cobra.Command, args []string) {

log.Fatalln("must use a subcommand")

	},
}

//get subcommand code
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a URL",
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

var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Post a URL",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {

			log.Fatalln("must specify the URL!")

		}
		client := http.Client{}
		var contentReader io.Reader
		if content != "" {
			contentReader = bytes.NewReader([]byte(content))
		}

		req, err := http.NewRequest("POST", args[0], contentReader)

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
	cmd.AddCommand(getCmd)
	cmd.AddCommand(postCmd)
	postCmd.PersistentFlags().StringVarP(&content,"content","c","","Content")
	cmd.Execute()
}
// go run cobraintro_03.go get -u foo -p bar https://httpbin.org/basic-auth/foo/bar
// go run cobraintro_03.go post  -c "this is my content" https://httpbin.org/post
