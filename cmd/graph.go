package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
)

var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "test connection to graph",
	Long:  "test connection to graph",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		// api := eos.New(viper.GetString("EosioEndpoint"))
		ctx := context.Background()

		// Query the balance for Alice and Bob.
		const q = `
		{
			var(func: has(member)){
			 members as member{
				}
			}
			members(func: uid(members)){
			  hash
			  creator
			  created_date
			  content_groups{
				expand(_all_){
				  expand(_all_)
				}
			  }
			  certificates{
				expand(_all_){
				  expand(_all_)
				}
			  }
			}
		  }
		`

		conn, err := grpc.Dial(viper.GetString("DgraphEndpoint"), grpc.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))

		txn := dgraphClient.NewReadOnlyTxn()
		defer txn.Discard(ctx)

		resp, err := txn.Query(context.Background(), q)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(resp.GetJson()))
	},
}

func newClient() *dgo.Dgraph {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	dialOpts := append([]grpc.DialOption{},
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)))
	d, err := grpc.Dial(viper.GetString("DgraphEndpoint"), dialOpts...)

	if err != nil {
		log.Fatal(err)
	}

	return dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
}

func init() {
	RootCmd.AddCommand(graphCmd)
}
