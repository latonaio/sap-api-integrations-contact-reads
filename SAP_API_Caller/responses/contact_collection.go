package responses

type ContactCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                              string `json:"ObjectID"`
			ContactID                             string `json:"ContactID"`
			ContactUUID                           string `json:"ContactUUID"`
			ExternalID                            string `json:"ExternalID"`
			ExternalSystem                        string `json:"ExternalSystem"`
			StatusCode                            string `json:"StatusCode"`
			StatusCodeText                        string `json:"StatusCodeText"`
			TitleCode                             string `json:"TitleCode"`
			TitleCodeText                         string `json:"TitleCodeText"`
			AcademicTitleCode                     string `json:"AcademicTitleCode"`
			AcademicTitleCodeText                 string `json:"AcademicTitleCodeText"`
			AdditionalAcademicTitleCode           string `json:"AdditionalAcademicTitleCode"`
			AdditionalAcademicTitleCodeText       string `json:"AdditionalAcademicTitleCodeText"`
			NamePrefixCode                        string `json:"NamePrefixCode"`
			NamePrefixCodeText                    string `json:"NamePrefixCodeText"`
			FirstName                             string `json:"FirstName"`
			LastName                              string `json:"LastName"`
			AdditionalFamilyName                  string `json:"AdditionalFamilyName"`
			Initials                              string `json:"Initials"`
			MiddleName                            string `json:"MiddleName"`
			Name                                  string `json:"Name"`
			GenderCode                            string `json:"GenderCode"`
			GenderCodeText                        string `json:"GenderCodeText"`
			MaritalStatusCode                     string `json:"MaritalStatusCode"`
			MaritalStatusCodeText                 string `json:"MaritalStatusCodeText"`
			LanguageCode                          string `json:"LanguageCode"`
			LanguageCodeText                      string `json:"LanguageCodeText"`
			NickName                              string `json:"NickName"`
			BirthDate                             string `json:"BirthDate"`
			BirthName                             string `json:"BirthName"`
			ContactPermissionCode                 string `json:"ContactPermissionCode"`
			ContactPermissionCodeText             string `json:"ContactPermissionCodeText"`
			ProfessionCode                        string `json:"ProfessionCode"`
			ProfessionCodeText                    string `json:"ProfessionCodeText"`
			PerceptionOfCompanyCode               string `json:"PerceptionOfCompanyCode"`
			PerceptionOfCompanyCodeText           string `json:"PerceptionOfCompanyCodeText"`
			DeviatingFullName                     string `json:"DeviatingFullName"`
			AccountID                             string `json:"AccountID"`
			AccountUUID                           string `json:"AccountUUID"`
			AccountFormattedName                  string `json:"AccountFormattedName"`
			Building                              string `json:"Building"`
			Floor                                 string `json:"Floor"`
			Room                                  string `json:"Room"`
			JobTitle                              string `json:"JobTitle"`
			FunctionCode                          string `json:"FunctionCode"`
			FunctionCodeText                      string `json:"FunctionCodeText"`
			DepartmentCode                        string `json:"DepartmentCode"`
			DepartmentCodeText                    string `json:"DepartmentCodeText"`
			Department                            string `json:"Department"`
			VIPContactCode                        string `json:"VIPContactCode"`
			VIPContactCodeText                    string `json:"VIPContactCodeText"`
			Phone                                 string `json:"Phone"`
			Mobile                                string `json:"Mobile"`
			Fax                                   string `json:"Fax"`
			Email                                 string `json:"Email"`
			EmailInvalidIndicator                 bool   `json:"EmailInvalidIndicator"`
			BestReachedByCode                     string `json:"BestReachedByCode"`
			BestReachedByCodeText                 string `json:"BestReachedByCodeText"`
			FormattedPostalAddressDescription     string `json:"FormattedPostalAddressDescription"`
			BusinessAddressCountryCode            string `json:"BusinessAddressCountryCode"`
			BusinessAddressCountryCodeText        string `json:"BusinessAddressCountryCodeText"`
			BusinessAddressStateCodeTextUpdatable string `json:"BusinessAddressStateCodeTextUpdatable"`
			BusinessAddressHouseNumber            string `json:"BusinessAddressHouseNumber"`
			BusinessAddressStreet                 string `json:"BusinessAddressStreet"`
			BusinessAddressCity                   string `json:"BusinessAddressCity"`
			BusinessAddressStreetPostalCode       string `json:"BusinessAddressStreetPostalCode"`
			BusinessAddressStateCode              string `json:"BusinessAddressStateCode"`
			BusinessAddressStateCodeText          string `json:"BusinessAddressStateCodeText"`
			CreationOn                            string `json:"CreationOn"`
			CreatedBy                             string `json:"CreatedBy"`
			CreatedByIdentityUUID                 string `json:"CreatedByIdentityUUID"`
			ChangedOn                             string `json:"ChangedOn"`
			ChangedBy                             string `json:"ChangedBy"`
			ChangedByIdentityUUID                 string `json:"ChangedByIdentityUUID"`
			ContactOwnerID                        string `json:"ContactOwnerID"`
			ContactOwnerUUID                      string `json:"ContactOwnerUUID"`
			NormalisedPhone                       string `json:"NormalisedPhone"`
			NormalisedMobile                      string `json:"NormalisedMobile"`
			EntityLastChangedOn                   string `json:"EntityLastChangedOn"`
			ETag                                  string `json:"ETag"`
			ContactIsContactPersonFor             struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"ContactIsContactPersonFor"`
			CorporateAccount struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"CorporateAccount"`
		} `json:"results"`
	} `json:"d"`
}
