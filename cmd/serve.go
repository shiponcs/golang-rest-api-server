package cmd

import (
	"fmt"
	"github.com/bookstore-rest-api-server/api"
	"github.com/spf13/cobra"
)

var (
	port string
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "starts the server",
	Long:  `It takes the port value and starts the server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serving at port " + port)
		api.ServeEndpoints(port)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringVarP(&port, "port", "p", "", "This flag sets the port of the server")
	err := serveCmd.MarkFlagRequired("port")
	if err != nil {
		fmt.Println(err)
		return
	}
	//serveCmd.PersistentFlags().BoolVarP(&auth, "auth", "a", true, "This flag will impose/bypass authentication to API server")
}
