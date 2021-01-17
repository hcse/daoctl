package views

import (
	"math/big"
	"strconv"

	"github.com/hypha-dao/daoctl/models"
	"github.com/hypha-dao/daoctl/util"

	"github.com/alexeyco/simpletable"
	"github.com/eoscanada/eos-go"
)

func roleHeader() *simpletable.Header {
	return &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Title"},
			{Align: simpletable.AlignCenter, Text: "Created Date"},
			{Align: simpletable.AlignCenter, Text: "Creator"},
			{Align: simpletable.AlignCenter, Text: "Annual USD"},
			{Align: simpletable.AlignCenter, Text: "FTE Cap"},
			{Align: simpletable.AlignCenter, Text: "Extended"},
			{Align: simpletable.AlignCenter, Text: "Min Time %"},
			{Align: simpletable.AlignCenter, Text: "Min Def %"},
			{Align: simpletable.AlignCenter, Text: "Ballot"},
			{Align: simpletable.AlignCenter, Text: "Hash"},
		},
	}
}

// RoleTable returns a string representing an output table for a Role array
func RoleTable(roles *[]models.Role) *simpletable.Table {

	table := simpletable.New()
	table.Header = roleHeader()

	usdFteTotal, _ := eos.NewAssetFromString("0.00 USD")

	for index, role := range *roles {

		usdFte := util.AssetMult(role.AnnualUSDSalary, big.NewFloat(role.FullTimeCapacity))
		usdFteTotal = usdFteTotal.Add(usdFte)

		r := []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: strconv.Itoa(index)},
			{Align: simpletable.AlignLeft, Text: role.GetNodeLabel()},
			{Align: simpletable.AlignRight, Text: role.CreatedDate.Time.Format("2006 Jan 02")},
			{Align: simpletable.AlignRight, Text: string(role.Creator)},
			{Align: simpletable.AlignRight, Text: util.FormatAsset(&role.AnnualUSDSalary, 0)},
			{Align: simpletable.AlignRight, Text: strconv.FormatFloat(role.FullTimeCapacity, 'f', 1, 64)},
			{Align: simpletable.AlignRight, Text: util.FormatAsset(&usdFte, 0)},
			{Align: simpletable.AlignRight, Text: strconv.FormatFloat(role.MinTime*100, 'f', -1, 64)},
			{Align: simpletable.AlignRight, Text: strconv.FormatFloat(role.MinDeferred*100, 'f', -1, 64)},
			{Align: simpletable.AlignRight, Text: string(role.BallotName)},
			{Align: simpletable.AlignRight, Text: role.Hash.String()},
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{}, {}, {}, {}, {},
			{Align: simpletable.AlignRight, Text: "Subtotal"},
			{Align: simpletable.AlignRight, Text: util.FormatAsset(&usdFteTotal, 0)},
			{}, {}, {}, {},
		},
	}

	return table
}
