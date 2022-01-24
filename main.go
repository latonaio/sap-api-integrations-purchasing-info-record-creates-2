package main

import (
	sap_api_caller "sap-api-integrations-purchasing-info-record-creates/SAP_API_Caller"
	sap_api_input_reader "sap-api-integrations-purchasing-info-record-creates/SAP_API_Input_Reader"
	"sap-api-integrations-purchasing-info-record-creates/config"

	"github.com/latonaio/golang-logging-library/logger"
	sap_api_post_header_setup "github.com/latonaio/sap-api-post-header-setup"
)

func main() {
	l := logger.NewLogger()
	conf := config.NewConf()
	fr := sap_api_input_reader.NewFileReader()
	pc := sap_api_post_header_setup.NewSAPPostClientWithOption(conf.SAP)
	caller := sap_api_caller.NewSAPAPICaller(
		conf.SAP.BaseURL(),
		"100",
		pc,
		l,
	)
	inputSDC := fr.ReadSDC("./Inputs/SDC_Purchasing_Info_Record_General_sample.json")
	accepter := getAccepter(inputSDC)
	general := inputSDC.ConvertToGenerl()
	material := inputSDC.ConvertToMaterial()
	materialGroup := inputSDC.ConvertToMaterialGroup()

	caller.AsyncPostPurchasingInfoRecord(
		general,
		material,
		materialGroup,
		accepter,
	)
}

func getAccepter(sdc sap_api_input_reader.SDC) []string {
	accepter := sdc.Accepter
	if len(sdc.Accepter) == 0 {
		accepter = []string{"All"}
	}

	if accepter[0] == "All" {
		accepter = []string{
			"General", "Material", "MaterialGroup",
		}
	}
	return accepter
}
