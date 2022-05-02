package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-contact-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetContact(contactID, customerID, accountFormattedName, accountID string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "ContactCollection":
			func() {
				c.ContactCollection(contactID)
				wg.Done()
			}()
		case "IndividualCustomerCollection":
			func() {
				c.IndividualCustomerCollection(customerID)
				wg.Done()
			}()
		case "ContactIsContactPersonFor":
			func() {
				c.ContactIsContactPersonFor(accountFormattedName)
				wg.Done()
			}()
		case "CorporateAccount":
			func() {
				c.CorporateAccount(accountID)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) ContactCollection(contactID string) {
	contactCollectionData, err := c.callContactSrvAPIRequirementContactCollection("ContactCollection", contactID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(contactCollectionData)

	contactIsContactPersonForData, err := c.callContactIsContactPersonFor(contactCollectionData[0].ToContactIsContactPersonFor)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(contactIsContactPersonForData)

	corporateAccountData, err := c.callCorporateAccount(contactCollectionData[0].ToCorporateAccount)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(corporateAccountData)

}

func (c *SAPAPICaller) callContactSrvAPIRequirementContactCollection(api, contactID string) ([]sap_api_output_formatter.ContactCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithContactCollection(req, contactID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToContactCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callContactIsContactPersonFor(url string) ([]sap_api_output_formatter.ContactIsContactPersonFor, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToContactIsContactPersonFor(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callCorporateAccount(url string) (*sap_api_output_formatter.ToCorporateAccount, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToCorporateAccount(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) IndividualCustomerCollection(customerID string) {
	individualCustomerCollectionData, err := c.callContactSrvAPIRequirementIndividualCustomerCollection("IndividualCustomerCollection", customerID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(individualCustomerCollectionData)

	individualCustomerAddressData, err := c.callIndividualCustomerAddress(individualCustomerCollectionData[0].ToIndividualCustomerAddress)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(individualCustomerAddressData)

}

func (c *SAPAPICaller) callContactSrvAPIRequirementIndividualCustomerCollection(api, customerID string) ([]sap_api_output_formatter.IndividualCustomerCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithIndividualCustomerCollection(req, customerID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToIndividualCustomerCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callIndividualCustomerAddress(url string) ([]sap_api_output_formatter.IndividualCustomerAddress, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToIndividualCustomerAddress(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ContactIsContactPersonFor(accountFormattedName string) {
	contactIsContactPersonForData, err := c.callContactSrvAPIRequirementContactIsContactPersonFor("ContactIsContactPersonForCollection", accountFormattedName)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(contactIsContactPersonForData)
}

func (c *SAPAPICaller) callContactSrvAPIRequirementContactIsContactPersonFor(api, accountFormattedName string) ([]sap_api_output_formatter.ContactIsContactPersonFor, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithContactIsContactPersonFor(req, accountFormattedName)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToContactIsContactPersonFor(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) CorporateAccount(accountID string) {
	data, err := c.callContactSrvAPIRequirementCorporateAccount("CorporateAccountCollection", accountID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callContactSrvAPIRequirementCorporateAccount(api, accountID string) ([]sap_api_output_formatter.CorporateAccount, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithCorporateAccount(req, accountID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToCorporateAccount(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithContactCollection(req *http.Request, contactID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("ContactID eq '%s'", contactID))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithIndividualCustomerCollection(req *http.Request, customerID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("CustomerID eq '%s'", customerID))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithContactIsContactPersonFor(req *http.Request, accountFormattedName string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("substringof('%s', AccountFormattedName)", accountFormattedName))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithCorporateAccount(req *http.Request, accountID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("AccountID eq '%s'", accountID))
	req.URL.RawQuery = params.Encode()
}
