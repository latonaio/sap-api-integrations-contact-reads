package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-contact-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToContactCollection(raw []byte, l *logger.Logger) ([]ContactCollection, error) {
	pm := &responses.ContactCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ContactCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	contactCollection := make([]ContactCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		contactCollection = append(contactCollection, ContactCollection{
			ObjectID:                              data.ObjectID,
			ContactID:                             data.ContactID,
			ContactUUID:                           data.ContactUUID,
			ExternalID:                            data.ExternalID,
			ExternalSystem:                        data.ExternalSystem,
			StatusCode:                            data.StatusCode,
			StatusCodeText:                        data.StatusCodeText,
			TitleCode:                             data.TitleCode,
			TitleCodeText:                         data.TitleCodeText,
			AcademicTitleCode:                     data.AcademicTitleCode,
			AcademicTitleCodeText:                 data.AcademicTitleCodeText,
			AdditionalAcademicTitleCode:           data.AdditionalAcademicTitleCode,
			AdditionalAcademicTitleCodeText:       data.AdditionalAcademicTitleCodeText,
			NamePrefixCode:                        data.NamePrefixCode,
			NamePrefixCodeText:                    data.NamePrefixCodeText,
			FirstName:                             data.FirstName,
			LastName:                              data.LastName,
			AdditionalFamilyName:                  data.AdditionalFamilyName,
			Initials:                              data.Initials,
			MiddleName:                            data.MiddleName,
			Name:                                  data.Name,
			GenderCode:                            data.GenderCode,
			GenderCodeText:                        data.GenderCodeText,
			MaritalStatusCode:                     data.MaritalStatusCode,
			MaritalStatusCodeText:                 data.MaritalStatusCodeText,
			LanguageCode:                          data.LanguageCode,
			LanguageCodeText:                      data.LanguageCodeText,
			NickName:                              data.NickName,
			BirthDate:                             data.BirthDate,
			BirthName:                             data.BirthName,
			ContactPermissionCode:                 data.ContactPermissionCode,
			ContactPermissionCodeText:             data.ContactPermissionCodeText,
			ProfessionCode:                        data.ProfessionCode,
			ProfessionCodeText:                    data.ProfessionCodeText,
			PerceptionOfCompanyCode:               data.PerceptionOfCompanyCode,
			PerceptionOfCompanyCodeText:           data.PerceptionOfCompanyCodeText,
			DeviatingFullName:                     data.DeviatingFullName,
			AccountID:                             data.AccountID,
			AccountUUID:                           data.AccountUUID,
			AccountFormattedName:                  data.AccountFormattedName,
			Building:                              data.Building,
			Floor:                                 data.Floor,
			Room:                                  data.Room,
			JobTitle:                              data.JobTitle,
			FunctionCode:                          data.FunctionCode,
			FunctionCodeText:                      data.FunctionCodeText,
			DepartmentCode:                        data.DepartmentCode,
			DepartmentCodeText:                    data.DepartmentCodeText,
			Department:                            data.Department,
			VIPContactCode:                        data.VIPContactCode,
			VIPContactCodeText:                    data.VIPContactCodeText,
			Phone:                                 data.Phone,
			Mobile:                                data.Mobile,
			Fax:                                   data.Fax,
			Email:                                 data.Email,
			EmailInvalidIndicator:                 data.EmailInvalidIndicator,
			BestReachedByCode:                     data.BestReachedByCode,
			BestReachedByCodeText:                 data.BestReachedByCodeText,
			FormattedPostalAddressDescription:     data.FormattedPostalAddressDescription,
			BusinessAddressCountryCode:            data.BusinessAddressCountryCode,
			BusinessAddressCountryCodeText:        data.BusinessAddressCountryCodeText,
			BusinessAddressStateCodeTextUpdatable: data.BusinessAddressStateCodeTextUpdatable,
			BusinessAddressHouseNumber:            data.BusinessAddressHouseNumber,
			BusinessAddressStreet:                 data.BusinessAddressStreet,
			BusinessAddressCity:                   data.BusinessAddressCity,
			BusinessAddressStreetPostalCode:       data.BusinessAddressStreetPostalCode,
			BusinessAddressStateCode:              data.BusinessAddressStateCode,
			BusinessAddressStateCodeText:          data.BusinessAddressStateCodeText,
			CreationOn:                            data.CreationOn,
			CreatedBy:                             data.CreatedBy,
			CreatedByIdentityUUID:                 data.CreatedByIdentityUUID,
			ChangedOn:                             data.ChangedOn,
			ChangedBy:                             data.ChangedBy,
			ChangedByIdentityUUID:                 data.ChangedByIdentityUUID,
			ContactOwnerID:                        data.ContactOwnerID,
			ContactOwnerUUID:                      data.ContactOwnerUUID,
			NormalisedPhone:                       data.NormalisedPhone,
			NormalisedMobile:                      data.NormalisedMobile,
			EntityLastChangedOn:                   data.EntityLastChangedOn,
			ToContactIsContactPersonFor:           data.ContactIsContactPersonFor.Deferred.URI,
			ToCorporateAccount:                    data.CorporateAccount.Deferred.URI,
		})
	}

	return contactCollection, nil
}

func ConvertToIndividualCustomerCollection(raw []byte, l *logger.Logger) ([]IndividualCustomerCollection, error) {
	pm := &responses.IndividualCustomerCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to IndividualCustomerCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	individualCustomerCollection := make([]IndividualCustomerCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		individualCustomerCollection = append(individualCustomerCollection, IndividualCustomerCollection{
			ObjectID:                          data.ObjectID,
			CustomerID:                        data.CustomerID,
			UUID:                              data.UUID,
			ExternalID:                        data.ExternalID,
			ExternalSystem:                    data.ExternalSystem,
			RoleCode:                          data.RoleCode,
			RoleCodeText:                      data.RoleCodeText,
			LifeCycleStatusCode:               data.LifeCycleStatusCode,
			LifeCycleStatusCodeText:           data.LifeCycleStatusCodeText,
			CustomerABCClassificationCode:     data.CustomerABCClassificationCode,
			CustomerABCClassificationCodeText: data.CustomerABCClassificationCodeText,
			ContactPermissionCode:             data.ContactPermissionCode,
			ContactPermissionCodeText:         data.ContactPermissionCodeText,
			TitleCode:                         data.TitleCode,
			TitleCodeText:                     data.TitleCodeText,
			AcademicTitleCode:                 data.AcademicTitleCode,
			AcademicTitleCodeText:             data.AcademicTitleCodeText,
			FirstName:                         data.FirstName,
			MiddleName:                        data.MiddleName,
			LastName:                          data.LastName,
			AdditionalLastName:                data.AdditionalLastName,
			Initials:                          data.Initials,
			NickName:                          data.NickName,
			GenderCode:                        data.GenderCode,
			GenderCodeText:                    data.GenderCodeText,
			NamePrefixCode:                    data.NamePrefixCode,
			NamePrefixCodeText:                data.NamePrefixCodeText,
			MaritalStatusCode:                 data.MaritalStatusCode,
			MaritalStatusCodeText:             data.MaritalStatusCodeText,
			LanguageCode:                      data.LanguageCode,
			LanguageCodeText:                  data.LanguageCodeText,
			BirthName:                         data.BirthName,
			BirthDate:                         data.BirthDate,
			NationalityCountryCode:            data.NationalityCountryCode,
			NationalityCountryCodeText:        data.NationalityCountryCodeText,
			ProfessionCode:                    data.ProfessionCode,
			ProfessionCodeText:                data.ProfessionCodeText,
			FormattedName:                     data.FormattedName,
			FormattedPostalAddressDescription: data.FormattedPostalAddressDescription,
			CountryCode:                       data.CountryCode,
			CountryCodeText:                   data.CountryCodeText,
			StateCode:                         data.StateCode,
			StateCodeText:                     data.StateCodeText,
			CareOfName:                        data.CareOfName,
			AddressLine1:                      data.AddressLine1,
			AddressLine2:                      data.AddressLine2,
			HouseNumber:                       data.HouseNumber,
			AdditionalHouseNumber:             data.AdditionalHouseNumber,
			Street:                            data.Street,
			AddressLine4:                      data.AddressLine4,
			AddressLine5:                      data.AddressLine5,
			District:                          data.District,
			City:                              data.City,
			DifferentCity:                     data.DifferentCity,
			StreetPostalCode:                  data.StreetPostalCode,
			County:                            data.County,
			POBoxIndicator:                    data.POBoxIndicator,
			POBox:                             data.POBox,
			POBoxPostalCode:                   data.POBoxPostalCode,
			POBoxDeviatingCountryCode:         data.POBoxDeviatingCountryCode,
			POBoxDeviatingCountryCodeText:     data.POBoxDeviatingCountryCodeText,
			POBoxDeviatingStateCode:           data.POBoxDeviatingStateCode,
			POBoxDeviatingStateCodeText:       data.POBoxDeviatingStateCodeText,
			POBoxDeviatingCity:                data.POBoxDeviatingCity,
			TimeZoneCode:                      data.TimeZoneCode,
			TimeZoneCodeText:                  data.TimeZoneCodeText,
			TaxJurisdictionCode:               data.TaxJurisdictionCode,
			TaxJurisdictionCodeText:           data.TaxJurisdictionCodeText,
			Building:                          data.Building,
			Floor:                             data.Floor,
			Room:                              data.Room,
			Phone:                             data.Phone,
			NormalisedPhone:                   data.NormalisedPhone,
			Mobile:                            data.Mobile,
			NormalisedMobile:                  data.NormalisedMobile,
			Fax:                               data.Fax,
			Email:                             data.Email,
			EmailInvalidIndicator:             data.EmailInvalidIndicator,
			WebSite:                           data.WebSite,
			BestReachedByCode:                 data.BestReachedByCode,
			BestReachedByCodeText:             data.BestReachedByCodeText,
			OrderBlockingReasonCode:           data.OrderBlockingReasonCode,
			OrderBlockingReasonCodeText:       data.OrderBlockingReasonCodeText,
			DeliveryBlockingReasonCode:        data.DeliveryBlockingReasonCode,
			DeliveryBlockingReasonCodeText:    data.DeliveryBlockingReasonCodeText,
			BillingBlockingReasonCode:         data.BillingBlockingReasonCode,
			BillingBlockingReasonCodeText:     data.BillingBlockingReasonCodeText,
			SalesSupportBlockingIndicator:     data.SalesSupportBlockingIndicator,
			RecommendedVisitingFrequency:      data.RecommendedVisitingFrequency,
			VisitDuration:                     data.VisitDuration,
			LastVisitingDate:                  data.LastVisitingDate,
			NextVisitingDate:                  data.NextVisitingDate,
			LatestRecommendedVisitingDate:     data.LatestRecommendedVisitingDate,
			OwnerID:                           data.OwnerID,
			OwnerUUID:                         data.OwnerUUID,
			CreationOn:                        data.CreationOn,
			CreatedBy:                         data.CreatedBy,
			CreatedByIdentityUUID:             data.CreatedByIdentityUUID,
			ChangedOn:                         data.ChangedOn,
			ChangedBy:                         data.ChangedBy,
			ChangedByIdentityUUID:             data.ChangedByIdentityUUID,
			EntityLastChangedOn:               data.EntityLastChangedOn,
			ToIndividualCustomerAddress:       data.IndividualCustomerAddress.Deferred.URI,
		})
	}

	return individualCustomerCollection, nil
}

func ConvertToContactIsContactPersonFor(raw []byte, l *logger.Logger) ([]ContactIsContactPersonFor, error) {
	pm := &responses.ContactIsContactPersonFor{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ContactIsContactPersonFor. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	contactIsContactPersonFor := make([]ContactIsContactPersonFor, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		contactIsContactPersonFor = append(contactIsContactPersonFor, ContactIsContactPersonFor{
			ObjectID:                data.ObjectID,
			ParentObjectID:          data.ParentObjectID,
			ETag:                    data.ETag,
			ContactID:               data.ContactID,
			AccountID:               data.AccountID,
			AccountFormattedName:    data.AccountFormattedName,
			ReverseMainIndicator:    data.ReverseMainIndicator,
			DepartmentCode:          data.DepartmentCode,
			DepartmentCodeText:      data.DepartmentCodeText,
			FunctionCode:            data.FunctionCode,
			FunctionCodeText:        data.FunctionCodeText,
			VIPReasonCode:           data.VIPReasonCode,
			VIPReasonCodeText:       data.VIPReasonCodeText,
			JobTitle:                data.JobTitle,
			Department:              data.Department,
			Building:                data.Building,
			Floor:                   data.Floor,
			Room:                    data.Room,
			Email:                   data.Email,
			EmailInvalidIndicator:   data.EmailInvalidIndicator,
			Fax:                     data.Fax,
			Mobile:                  data.Mobile,
			Phone:                   data.Phone,
			BestReachedByCode:       data.BestReachedByCode,
			BestReachedByCodeText:   data.BestReachedByCodeText,
			OrganisationAddressUUID: data.OrganisationAddressUUID,
			EntityLastChangedOn:     data.EntityLastChangedOn,
		})
	}

	return contactIsContactPersonFor, nil
}

func ConvertToToCorporateAccount(raw []byte, l *logger.Logger) (*ToCorporateAccount, error) {
	pm := &responses.ToCorporateAccount{}
	err := json.Unmarshal(raw, &pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToCorporateAccount. unmarshal error: %w", err)
	}

	return &ToCorporateAccount{
		ObjectID:                          pm.D.Results.ObjectID,
		AccountID:                         pm.D.Results.AccountID,
		UUID:                              pm.D.Results.UUID,
		ExternalID:                        pm.D.Results.ExternalID,
		ExternalSystem:                    pm.D.Results.ExternalSystem,
		RoleCode:                          pm.D.Results.RoleCode,
		RoleCodeText:                      pm.D.Results.RoleCodeText,
		LifeCycleStatusCode:               pm.D.Results.LifeCycleStatusCode,
		LifeCycleStatusCodeText:           pm.D.Results.LifeCycleStatusCodeText,
		Dunsid:                            pm.D.Results.DunsId,
		LegalFormCode:                     pm.D.Results.LegalFormCode,
		LegalFormCodeText:                 pm.D.Results.LegalFormCodeText,
		CustomerABCClassificationCode:     pm.D.Results.CustomerABCClassificationCode,
		CustomerABCClassificationCodeText: pm.D.Results.CustomerABCClassificationCodeText,
		NielsenRegionCode:                 pm.D.Results.NielsenRegionCode,
		NielsenRegionCodeText:             pm.D.Results.NielsenRegionCodeText,
		IndustrialSectorCode:              pm.D.Results.IndustrialSectorCode,
		IndustrialSectorCodeText:          pm.D.Results.IndustrialSectorCodeText,
		ContactPermissionCode:             pm.D.Results.ContactPermissionCode,
		ContactPermissionCodeText:         pm.D.Results.ContactPermissionCodeText,
		BusinessPartnerFormattedName:      pm.D.Results.BusinessPartnerFormattedName,
		Name:                              pm.D.Results.Name,
		AdditionalName:                    pm.D.Results.AdditionalName,
		AdditionalName2:                   pm.D.Results.AdditionalName2,
		AdditionalName3:                   pm.D.Results.AdditionalName3,
		CurrentDefaultAddressUUID:         pm.D.Results.CurrentDefaultAddressUUID,
		FormattedPostalAddressDescription: pm.D.Results.FormattedPostalAddressDescription,
		CountryCode:                       pm.D.Results.CountryCode,
		CountryCodeText:                   pm.D.Results.CountryCodeText,
		StateCode:                         pm.D.Results.StateCode,
		StateCodeText:                     pm.D.Results.StateCodeText,
		CareOfName:                        pm.D.Results.CareOfName,
		AddressLine1:                      pm.D.Results.AddressLine1,
		AddressLine2:                      pm.D.Results.AddressLine2,
		HouseNumber:                       pm.D.Results.HouseNumber,
		AdditionalHouseNumber:             pm.D.Results.AdditionalHouseNumber,
		Street:                            pm.D.Results.Street,
		AddressLine4:                      pm.D.Results.AddressLine4,
		AddressLine5:                      pm.D.Results.AddressLine5,
		District:                          pm.D.Results.District,
		City:                              pm.D.Results.City,
		DifferentCity:                     pm.D.Results.DifferentCity,
		StreetPostalCode:                  pm.D.Results.StreetPostalCode,
		County:                            pm.D.Results.County,
		CompanyPostalCode:                 pm.D.Results.CompanyPostalCode,
		POBoxIndicator:                    pm.D.Results.POBoxIndicator,
		POBox:                             pm.D.Results.POBox,
		POBoxPostalCode:                   pm.D.Results.POBoxPostalCode,
		POBoxDeviatingCountryCode:         pm.D.Results.POBoxDeviatingCountryCode,
		POBoxDeviatingCountryCodeText:     pm.D.Results.POBoxDeviatingCountryCodeText,
		POBoxDeviatingRegionCode:          pm.D.Results.POBoxDeviatingRegionCode,
		POBoxDeviatingRegionCodeText:      pm.D.Results.POBoxDeviatingRegionCodeText,
		POBoxDeviatingCity:                pm.D.Results.POBoxDeviatingCity,
		TimeZoneCode:                      pm.D.Results.TimeZoneCode,
		TimeZoneCodeText:                  pm.D.Results.TimeZoneCodeText,
		Building:                          pm.D.Results.Building,
		Floor:                             pm.D.Results.Floor,
		Room:                              pm.D.Results.Room,
		Phone:                             pm.D.Results.Phone,
		NormalisedPhone:                   pm.D.Results.NormalisedPhone,
		Mobile:                            pm.D.Results.Mobile,
		NormalisedMobile:                  pm.D.Results.NormalisedMobile,
		Fax:                               pm.D.Results.Fax,
		Email:                             pm.D.Results.Email,
		WebSite:                           pm.D.Results.WebSite,
		LanguageCode:                      pm.D.Results.LanguageCode,
		LanguageCodeText:                  pm.D.Results.LanguageCodeText,
		BestReachedByCode:                 pm.D.Results.BestReachedByCode,
		BestReachedByCodeText:             pm.D.Results.BestReachedByCodeText,
		OrderBlockingReasonCode:           pm.D.Results.OrderBlockingReasonCode,
		OrderBlockingReasonCodeText:       pm.D.Results.OrderBlockingReasonCodeText,
		DeliveryBlockingReasonCode:        pm.D.Results.DeliveryBlockingReasonCode,
		DeliveryBlockingReasonCodeText:    pm.D.Results.DeliveryBlockingReasonCodeText,
		BillingBlockingReasonCode:         pm.D.Results.BillingBlockingReasonCode,
		BillingBlockingReasonCodeText:     pm.D.Results.BillingBlockingReasonCodeText,
		SalesSupportBlockingIndicator:     pm.D.Results.SalesSupportBlockingIndicator,
		LegalCompetenceIndicator:          pm.D.Results.LegalCompetenceIndicator,
		RecommendedVisitingFrequency:      pm.D.Results.RecommendedVisitingFrequency,
		VisitDuration:                     pm.D.Results.VisitDuration,
		LastVisitingDate:                  pm.D.Results.LastVisitingDate,
		NextVisitingDate:                  pm.D.Results.NextVisitingDate,
		LatestRecommendedVisitingDate:     pm.D.Results.LatestRecommendedVisitingDate,
		OwnerID:                           pm.D.Results.OwnerID,
		OwnerUUID:                         pm.D.Results.OwnerUUID,
		ParentAccountID:                   pm.D.Results.ParentAccountID,
		CreationOn:                        pm.D.Results.CreationOn,
		CreatedBy:                         pm.D.Results.CreatedBy,
		CreatedByIdentityUUID:             pm.D.Results.CreatedByIdentityUUID,
		ChangedOn:                         pm.D.Results.ChangedOn,
		ChangedBy:                         pm.D.Results.ChangedBy,
		ChangedByIdentityUUID:             pm.D.Results.ChangedByIdentityUUID,
		EntityLastChangedOn:               pm.D.Results.EntityLastChangedOn,
	}, nil
}

func ConvertToIndividualCustomerAddress(raw []byte, l *logger.Logger) ([]IndividualCustomerAddress, error) {
	pm := &responses.ToIndividualCustomerAddress{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to IndividualCustomerAddress. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	individualCustomerAddress := make([]IndividualCustomerAddress, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		individualCustomerAddress = append(individualCustomerAddress, IndividualCustomerAddress{
			ObjectID:                              data.ObjectID,
			ParentObjectID:                        data.ParentObjectID,
			CustomerID:                            data.CustomerID,
			MainIndicator:                         data.MainIndicator,
			ShipTo:                                data.ShipTo,
			DefaultShipTo:                         data.DefaultShipTo,
			BillTo:                                data.BillTo,
			DefaultBillTo:                         data.DefaultBillTo,
			FormattedPostalAddressDescription:     data.FormattedPostalAddressDescription,
			FormattedAddressFirstLineDescription:  data.FormattedAddressFirstLineDescription,
			FormattedAddressSecondLineDescription: data.FormattedAddressSecondLineDescription,
			FormattedAddressThirdLineDescription:  data.FormattedAddressThirdLineDescription,
			FormattedAddressFourthLineDescription: data.FormattedAddressFourthLineDescription,
			FormattedPostalAddressFirstLineDescription:  data.FormattedPostalAddressFirstLineDescription,
			FormattedPostalAddressSecondLineDescription: data.FormattedPostalAddressSecondLineDescription,
			FormattedPostalAddressThirdLineDescription:  data.FormattedPostalAddressThirdLineDescription,
			CountryCode:                   data.CountryCode,
			CountryCodeText:               data.CountryCodeText,
			StateCode:                     data.StateCode,
			StateCodeText:                 data.StateCodeText,
			CareOfName:                    data.CareOfName,
			AddressLine1:                  data.AddressLine1,
			AddressLine2:                  data.AddressLine2,
			HouseNumber:                   data.HouseNumber,
			AdditionalHouseNumber:         data.AdditionalHouseNumber,
			Street:                        data.Street,
			AddressLine4:                  data.AddressLine4,
			AddressLine5:                  data.AddressLine5,
			District:                      data.District,
			City:                          data.City,
			DifferentCity:                 data.DifferentCity,
			StreetPostalCode:              data.StreetPostalCode,
			County:                        data.County,
			POBoxIndicator:                data.POBoxIndicator,
			POBox:                         data.POBox,
			POBoxPostalCode:               data.POBoxPostalCode,
			POBoxDeviatingCountryCode:     data.POBoxDeviatingCountryCode,
			POBoxDeviatingCountryCodeText: data.POBoxDeviatingCountryCodeText,
			POBoxDeviatingStateCode:       data.POBoxDeviatingStateCode,
			POBoxDeviatingStateCodeText:   data.POBoxDeviatingStateCodeText,
			POBoxDeviatingCity:            data.POBoxDeviatingCity,
			TimeZoneCode:                  data.TimeZoneCode,
			TimeZoneCodeText:              data.TimeZoneCodeText,
			Latitude:                      data.Latitude,
			Longitude:                     data.Longitude,
			Building:                      data.Building,
			Floor:                         data.Floor,
			Room:                          data.Room,
			Phone:                         data.Phone,
			NormalisedPhone:               data.NormalisedPhone,
			Mobile:                        data.Mobile,
			NormalisedMobile:              data.NormalisedMobile,
			Fax:                           data.Fax,
			Email:                         data.Email,
			EmailInvalidIndicator:         data.EmailInvalidIndicator,
			WebSite:                       data.WebSite,
			BestReachedByCode:             data.BestReachedByCode,
			BestReachedByCodeText:         data.BestReachedByCodeText,
		})
	}

	return individualCustomerAddress, nil
}

func ConvertToCorporateAccount(raw []byte, l *logger.Logger) ([]CorporateAccount, error) {
	pm := &responses.CorporateAccount{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to CorporateAccount. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	corporateAccount := make([]CorporateAccount, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		corporateAccount = append(corporateAccount, CorporateAccount{
			ObjectID:                          data.ObjectID,
			AccountID:                         data.AccountID,
			UUID:                              data.UUID,
			ExternalID:                        data.ExternalID,
			ExternalSystem:                    data.ExternalSystem,
			RoleCode:                          data.RoleCode,
			RoleCodeText:                      data.RoleCodeText,
			LifeCycleStatusCode:               data.LifeCycleStatusCode,
			LifeCycleStatusCodeText:           data.LifeCycleStatusCodeText,
			Dunsid:                            data.Dunsid,
			LegalFormCode:                     data.LegalFormCode,
			LegalFormCodeText:                 data.LegalFormCodeText,
			CustomerABCClassificationCode:     data.CustomerABCClassificationCode,
			CustomerABCClassificationCodeText: data.CustomerABCClassificationCodeText,
			NielsenRegionCode:                 data.NielsenRegionCode,
			NielsenRegionCodeText:             data.NielsenRegionCodeText,
			IndustrialSectorCode:              data.IndustrialSectorCode,
			IndustrialSectorCodeText:          data.IndustrialSectorCodeText,
			ContactPermissionCode:             data.ContactPermissionCode,
			ContactPermissionCodeText:         data.ContactPermissionCodeText,
			BusinessPartnerFormattedName:      data.BusinessPartnerFormattedName,
			Name:                              data.Name,
			AdditionalName:                    data.AdditionalName,
			AdditionalName2:                   data.AdditionalName2,
			AdditionalName3:                   data.AdditionalName3,
			CurrentDefaultAddressUUID:         data.CurrentDefaultAddressUUID,
			FormattedPostalAddressDescription: data.FormattedPostalAddressDescription,
			CountryCode:                       data.CountryCode,
			CountryCodeText:                   data.CountryCodeText,
			StateCode:                         data.StateCode,
			StateCodeText:                     data.StateCodeText,
			CareOfName:                        data.CareOfName,
			AddressLine1:                      data.AddressLine1,
			AddressLine2:                      data.AddressLine2,
			HouseNumber:                       data.HouseNumber,
			AdditionalHouseNumber:             data.AdditionalHouseNumber,
			Street:                            data.Street,
			AddressLine4:                      data.AddressLine4,
			AddressLine5:                      data.AddressLine5,
			District:                          data.District,
			City:                              data.City,
			DifferentCity:                     data.DifferentCity,
			StreetPostalCode:                  data.StreetPostalCode,
			County:                            data.County,
			CompanyPostalCode:                 data.CompanyPostalCode,
			POBoxIndicator:                    data.POBoxIndicator,
			POBox:                             data.POBox,
			POBoxPostalCode:                   data.POBoxPostalCode,
			POBoxDeviatingCountryCode:         data.POBoxDeviatingCountryCode,
			POBoxDeviatingCountryCodeText:     data.POBoxDeviatingCountryCodeText,
			POBoxDeviatingRegionCode:          data.POBoxDeviatingRegionCode,
			POBoxDeviatingRegionCodeText:      data.POBoxDeviatingRegionCodeText,
			POBoxDeviatingCity:                data.POBoxDeviatingCity,
			TimeZoneCode:                      data.TimeZoneCode,
			TimeZoneCodeText:                  data.TimeZoneCodeText,
			Building:                          data.Building,
			Floor:                             data.Floor,
			Room:                              data.Room,
			Phone:                             data.Phone,
			NormalisedPhone:                   data.NormalisedPhone,
			Mobile:                            data.Mobile,
			NormalisedMobile:                  data.NormalisedMobile,
			Fax:                               data.Fax,
			Email:                             data.Email,
			WebSite:                           data.WebSite,
			LanguageCode:                      data.LanguageCode,
			LanguageCodeText:                  data.LanguageCodeText,
			BestReachedByCode:                 data.BestReachedByCode,
			BestReachedByCodeText:             data.BestReachedByCodeText,
			OrderBlockingReasonCode:           data.OrderBlockingReasonCode,
			OrderBlockingReasonCodeText:       data.OrderBlockingReasonCodeText,
			DeliveryBlockingReasonCode:        data.DeliveryBlockingReasonCode,
			DeliveryBlockingReasonCodeText:    data.DeliveryBlockingReasonCodeText,
			BillingBlockingReasonCode:         data.BillingBlockingReasonCode,
			BillingBlockingReasonCodeText:     data.BillingBlockingReasonCodeText,
			SalesSupportBlockingIndicator:     data.SalesSupportBlockingIndicator,
			LegalCompetenceIndicator:          data.LegalCompetenceIndicator,
			RecommendedVisitingFrequency:      data.RecommendedVisitingFrequency,
			VisitDuration:                     data.VisitDuration,
			LastVisitingDate:                  data.LastVisitingDate,
			NextVisitingDate:                  data.NextVisitingDate,
			LatestRecommendedVisitingDate:     data.LatestRecommendedVisitingDate,
			OwnerID:                           data.OwnerID,
			OwnerUUID:                         data.OwnerUUID,
			ParentAccountID:                   data.ParentAccountID,
			CreationOn:                        data.CreationOn,
			CreatedBy:                         data.CreatedBy,
			CreatedByIdentityUUID:             data.CreatedByIdentityUUID,
			ChangedOn:                         data.ChangedOn,
			ChangedBy:                         data.ChangedBy,
			ChangedByIdentityUUID:             data.ChangedByIdentityUUID,
			EntityLastChangedOn:               data.EntityLastChangedOn,
		})
	}

	return corporateAccount, nil
}
