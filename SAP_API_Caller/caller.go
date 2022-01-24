package sap_api_caller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sap-api-integrations-purchasing-info-record-creates/SAP_API_Caller/requests"
	sap_api_output_formatter "sap-api-integrations-purchasing-info-record-creates/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
	sap_api_post_header_setup "github.com/latonaio/sap-api-post-header-setup"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	postClient      *sap_api_post_header_setup.SAPPostClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, postClient *sap_api_post_header_setup.SAPPostClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		postClient:      postClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncPostPurchasingInfoRecord(
	general *requests.General,
	material *requests.PurchasingOrganizationPlant,
	materialGroup *requests.PurchasingOrganizationPlant,
	accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	for _, fn := range accepter {
		switch fn {
		case "General":
			func() {
				c.General(general)
				wg.Done()
			}()
		case "Material":
			func() {
				c.Material(material)
				wg.Done()
			}()
		case "MaterialGroup":
			func() {
				c.MaterialGroup(materialGroup)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) General(general *requests.General) {
	outputDataGeneral, err := c.callPurchasingInfoRecordSrvAPIRequirementGeneral("A_PurchasingInfoRecord", general)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataGeneral)
}

func (c *SAPAPICaller) callPurchasingInfoRecordSrvAPIRequirementGeneral(api string, general *requests.General) (*sap_api_output_formatter.General, error) {
	body, err := json.Marshal(general)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_INFORECORD_PROCESS_SRV", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.postClient.POST(url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToGeneral(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Material(material *requests.PurchasingOrganizationPlant) {
	outputDataMaterial, err := c.callPurchasingInfoRecordSrvAPIRequirementMaterial("A_PurgInfoRecdOrgPlantData", material)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataMaterial)
}

func (c *SAPAPICaller) callPurchasingInfoRecordSrvAPIRequirementMaterial(api string, material *requests.PurchasingOrganizationPlant) (*sap_api_output_formatter.PurchasingOrganizationPlant, error) {
	body, err := json.Marshal(material)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_INFORECORD_PROCESS_SRV", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.postClient.POST(url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToPurchasingOrganizationPlant(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) MaterialGroup(materialGroup *requests.PurchasingOrganizationPlant) {
	outputDataMaterial, err := c.callPurchasingInfoRecordSrvAPIRequirementMaterial("A_PurgInfoRecdOrgPlantData", materialGroup)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataMaterial)
}

func (c *SAPAPICaller) callPurchasingInfoRecordSrvAPIRequirementMaterialGroup(api string, materialGroiup *requests.PurchasingOrganizationPlant) (*sap_api_output_formatter.PurchasingOrganizationPlant, error) {
	body, err := json.Marshal(materialGroiup)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_INFORECORD_PROCESS_SRV", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.postClient.POST(url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToPurchasingOrganizationPlant(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) addQuerySAPClient(params map[string]string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["sap-client"] = c.sapClientNumber
	return params
}
