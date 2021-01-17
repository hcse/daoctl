package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/alexeyco/simpletable"
	eos "github.com/eoscanada/eos-go"
	"github.com/hypha-dao/daoctl/models"
	"github.com/hypha-dao/daoctl/views"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var getRolesCmd = &cobra.Command{
	Use:   "roles",
	Short: "retrieve roles",
	Long:  "retrieve all active roles For a json dump, append the argument --json.",
	Run: func(cmd *cobra.Command, args []string) {
		api := eos.New(viper.GetString("EosioEndpoint"))
		ctx := context.Background()

		if viper.GetBool("global-active") == true {
			printRolesTable(ctx, api, "Current Roles")
		}

		// if viper.GetBool("global-include-proposals") == true {
		// 	printRolesTable(ctx, api, periods, "Current Role Proposals", "proposal")
		// }

		// if viper.GetBool("global-failed-proposals") == true {
		// 	printRolesTable(ctx, api, periods, "Failed Role Proposals", "failedprops")
		// }

		// if viper.GetBool("global-include-archive") == true {
		// 	printRolesTable(ctx, api, periods, "Archive of Role Proposals", "proparchive")
		// }
	},
}

func printRolesTable(ctx context.Context, api *eos.API, title string) {
	fmt.Println("\n", title)
	roles, err := models.Roles(ctx, api)
	if err != nil {
		log.Printf("Unable to retrieve roles %v", err)
		return
	}
	rolesTable := views.RoleTable(&roles)
	rolesTable.SetStyle(simpletable.StyleCompactLite)

	fmt.Println("\n" + rolesTable.String() + "\n\n")
}

func init() {
	getCmd.AddCommand(getRolesCmd)
}
