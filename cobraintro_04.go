package main

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		username := viper.GetString("username")
		password := viper.GetString("password")
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

		username := viper.GetString("username")
		password := viper.GetString("password")
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
	cmd.PersistentFlags().StringP("username","u",viper.GetString("credentials.username"),"Username")
	cmd.PersistentFlags().StringP("password","p",viper.GetString("credentials.password"),"Password")
	viper.BindPFlag("username",cmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("password",cmd.PersistentFlags().Lookup("password"))
	cmd.AddCommand(getCmd)
	cmd.AddCommand(postCmd)
	postCmd.PersistentFlags().StringVarP(&content,"content","c","","Content")
	viper.BindPFlag("content",cmd.PersistentFlags().Lookup("content"))
	cmd.Execute()
}

func init(){
	viper.AddConfigPath(".")
	viper.SetConfigName("cobra")
	viper.ReadInConfig()
}
// go run cobraintro_04.go get  https://httpbin.org/basic-auth/foo/bar
//  ../../../../bin/cobraintro_04 get  -u foo  -p bar https://httpbin.org/basic-auth/foo/bar
