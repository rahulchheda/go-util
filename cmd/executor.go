package cmd

import (
	"context"
	"fmt"

	"github.com/litmuschaos/go-util/internal/grpc/domain"
	"github.com/litmuschaos/go-util/internal/grpc/service"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var executeCmd = &cobra.Command{
	Use:   "execute",
	Short: "Execute the file",
	Long: `This CMD is used for executing the file with a name

	If no blog post is found for the ID it will return a 'Not Found' error`,
	RunE: func(cmd *cobra.Command, Args []string) error {
		filename, err := cmd.Flags().GetString("filename")
		if err != nil {
			return err
		}
		req := domain.Name{
			FileName: filename,
		}
		serverAddress := "localhost:7000"
		conn, e := grpc.Dial(serverAddress, grpc.WithInsecure())
		if e != nil {
			panic(e)
		}
		client := service.NewExecutorServiceClient(conn)
		res, err := client.Execute(context.Background(), &req)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

func init() {
	executeCmd.Flags().String("filename", "", "The name of the file")
	executeCmd.MarkFlagRequired("filename")
	rootCmd.AddCommand(executeCmd)
}
