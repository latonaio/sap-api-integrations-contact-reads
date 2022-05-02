package responses

type ToIndividualCustomerAddress struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                                    string `json:"ObjectID"`
			ParentObjectID                              string `json:"ParentObjectID"`
			CustomerID                                  string `json:"CustomerID"`
			MainIndicator                               bool   `json:"MainIndicator"`
			ShipTo                                      bool   `json:"ShipTo"`
			DefaultShipTo                               bool   `json:"DefaultShipTo"`
			BillTo                                      bool   `json:"BillTo"`
			DefaultBillTo                               bool   `json:"DefaultBillTo"`
			FormattedPostalAddressDescription           string `json:"FormattedPostalAddressDescription"`
			FormattedAddressFirstLineDescription        string `json:"FormattedAddressFirstLineDescription"`
			FormattedAddressSecondLineDescription       string `json:"FormattedAddressSecondLineDescription"`
			FormattedAddressThirdLineDescription        string `json:"FormattedAddressThirdLineDescription"`
			FormattedAddressFourthLineDescription       string `json:"FormattedAddressFourthLineDescription"`
			FormattedPostalAddressFirstLineDescription  string `json:"FormattedPostalAddressFirstLineDescription"`
			FormattedPostalAddressSecondLineDescription string `json:"FormattedPostalAddressSecondLineDescription"`
			FormattedPostalAddressThirdLineDescription  string `json:"FormattedPostalAddressThirdLineDescription"`
			CountryCode                                 string `json:"CountryCode"`
			CountryCodeText                             string `json:"CountryCodeText"`
			StateCode                                   string `json:"StateCode"`
			StateCodeText                               string `json:"StateCodeText"`
			CareOfName                                  string `json:"CareOfName"`
			AddressLine1                                string `json:"AddressLine1"`
			AddressLine2                                string `json:"AddressLine2"`
			HouseNumber                                 string `json:"HouseNumber"`
			AdditionalHouseNumber                       string `json:"AdditionalHouseNumber"`
			Street                                      string `json:"Street"`
			AddressLine4                                string `json:"AddressLine4"`
			AddressLine5                                string `json:"AddressLine5"`
			District                                    string `json:"District"`
			City                                        string `json:"City"`
			DifferentCity                               string `json:"DifferentCity"`
			StreetPostalCode                            string `json:"StreetPostalCode"`
			County                                      string `json:"County"`
			POBoxIndicator                              bool   `json:"POBoxIndicator"`
			POBox                                       string `json:"POBox"`
			POBoxPostalCode                             string `json:"POBoxPostalCode"`
			POBoxDeviatingCountryCode                   string `json:"POBoxDeviatingCountryCode"`
			POBoxDeviatingCountryCodeText               string `json:"POBoxDeviatingCountryCodeText"`
			POBoxDeviatingStateCode                     string `json:"POBoxDeviatingStateCode"`
			POBoxDeviatingStateCodeText                 string `json:"POBoxDeviatingStateCodeText"`
			POBoxDeviatingCity                          string `json:"POBoxDeviatingCity"`
			TimeZoneCode                                string `json:"TimeZoneCode"`
			TimeZoneCodeText                            string `json:"TimeZoneCodeText"`
			Latitude                                    string `json:"Latitude"`
			Longitude                                   string `json:"Longitude"`
			Building                                    string `json:"Building"`
			Floor                                       string `json:"Floor"`
			Room                                        string `json:"Room"`
			Phone                                       string `json:"Phone"`
			NormalisedPhone                             string `json:"NormalisedPhone"`
			Mobile                                      string `json:"Mobile"`
			NormalisedMobile                            string `json:"NormalisedMobile"`
			Fax                                         string `json:"Fax"`
			Email                                       string `json:"Email"`
			EmailInvalidIndicator                       bool   `json:"EmailInvalidIndicator"`
			WebSite                                     string `json:"WebSite"`
			BestReachedByCode                           string `json:"BestReachedByCode"`
			BestReachedByCodeText                       string `json:"BestReachedByCodeText"`
		} `json:"results"`
	} `json:"d"`
}
