package responses

type ToContactIsContactPersonFor struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                string `json:"ObjectID"`
			ParentObjectID          string `json:"ParentObjectID"`
			ETag                    string `json:"ETag"`
			ContactID               string `json:"ContactID"`
			AccountID               string `json:"AccountID"`
			AccountFormattedName    string `json:"AccountFormattedName"`
			ReverseMainIndicator    bool   `json:"ReverseMainIndicator"`
			DepartmentCode          string `json:"DepartmentCode"`
			DepartmentCodeText      string `json:"DepartmentCodeText"`
			FunctionCode            string `json:"FunctionCode"`
			FunctionCodeText        string `json:"FunctionCodeText"`
			VIPReasonCode           string `json:"VIPReasonCode"`
			VIPReasonCodeText       string `json:"VIPReasonCodeText"`
			JobTitle                string `json:"JobTitle"`
			Department              string `json:"Department"`
			Building                string `json:"Building"`
			Floor                   string `json:"Floor"`
			Room                    string `json:"Room"`
			Email                   string `json:"Email"`
			EmailInvalidIndicator   bool   `json:"EmailInvalidIndicator"`
			Fax                     string `json:"Fax"`
			Mobile                  string `json:"Mobile"`
			Phone                   string `json:"Phone"`
			BestReachedByCode       string `json:"BestReachedByCode"`
			BestReachedByCodeText   string `json:"BestReachedByCodeText"`
			OrganisationAddressUUID string `json:"OrganisationAddressUUID"`
			EntityLastChangedOn     string `json:"EntityLastChangedOn"`
		} `json:"results"`
	} `json:"d"`
}
