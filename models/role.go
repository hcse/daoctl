package models

import (
	"context"
	"fmt"
	"log"

	"github.com/hypha-dao/document-graph/docgraph"
	"github.com/spf13/viper"

	"github.com/eoscanada/eos-go"
)

// Role is an approved or proposed role for the DAO
type Role struct {
	docgraph.Document
	BallotName       eos.Name
	AnnualUSDSalary  eos.Asset
	MinTime          float64
	MinDeferred      float64
	FullTimeCapacity float64
}

// Roles provides the set of active approved roles
func Roles(ctx context.Context, api *eos.API) ([]Role, error) {
	dho, err := docgraph.LoadDocument(ctx, api, eos.AN(viper.GetString("DAOContract")), viper.GetString("DHONode"))
	if err != nil {
		return []Role{}, fmt.Errorf("error retrieving DHO document %v", err)
	}

	roleDocs, err := docgraph.GetDocumentsWithEdge(ctx, api, eos.AN(viper.GetString("DAOContract")), dho, eos.Name("role"))
	if err != nil {
		return []Role{}, fmt.Errorf("error retrieving role documents %v", err)
	}

	roles := make([]Role, len(roleDocs))
	for index, roleDoc := range roleDocs {
		roles[index], err = toRole(&roleDoc)
		if err != nil {
			return []Role{}, fmt.Errorf("cannot convert document to role %v %v", roleDoc.Hash.String(), err)
		}
	}

	return roles, nil
}

// ToRole converts a Document to a Role instance
func toRole(d *docgraph.Document) (Role, error) {
	role := Role{Document: *d}

	details, err := d.GetContentGroup("details")
	if err != nil {
		return Role{}, fmt.Errorf("unable to retrieve details group from document %v %v", d.Hash.String(), err)
	}

	role.AnnualUSDSalary = getSalary(details)
	role.MinTime = float64(getInt(details, "min_time_share_x100")) / 100
	role.MinDeferred = float64(getInt(details, "min_deferred_x100")) / 100
	role.FullTimeCapacity = float64(getInt(details, "fulltime_capacity_x100")) / 100

	system, err := d.GetContentGroup("system")
	if err != nil {
		return Role{}, fmt.Errorf("unable to retrieve system group from document %v %v", d.Hash.String(), err)
	}
	role.BallotName = getBallot(system)

	return role, nil
}

func getSalary(details *docgraph.ContentGroup) eos.Asset {
	var annualUsdSalary eos.Asset
	annualUsdSalaryContent, err := details.GetContent("annual_usd_salary")
	if err != nil {
		log.Printf("Default to 0.00 USD: unable to retrieve annual_usd_salary from details group of document: %v", err)
		annualUsdSalary, _ = eos.NewAssetFromString("0.00 USD")
	} else {
		annualUsdSalary, err = annualUsdSalaryContent.Asset()
		if err != nil {
			annualUsdSalary, _ = eos.NewAssetFromString("0.00 USD")
		}
	}
	return annualUsdSalary
}

func getInt(details *docgraph.ContentGroup, label string) int64 {
	var intValue int64
	intValueContent, err := details.GetContent(label)
	if err != nil {
		log.Printf("Defaulting to zero: unable to retrieve %v from details group of document %v :: ", label, err)
		intValue = 0
	} else {
		intValue, err = intValueContent.Int64()
		if err != nil {
			intValue = 0
		}
	}
	return intValue
}

func getBallot(system *docgraph.ContentGroup) eos.Name {
	var ballotName eos.Name
	ballotNameContent, err := system.GetContent("ballot_id")
	if err != nil {
		log.Printf("Default to zero: unable to retrieve ballot_id from system group of document %v :: ", err)
		ballotName = eos.Name("")
	} else {
		ballotName, err = ballotNameContent.Name()
		if err != nil {
			ballotName = eos.Name("")
		}
	}
	return ballotName
}
