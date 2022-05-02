# sap-api-integrations-contact-reads  
sap-api-integrations-contact-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API コンタクト データを取得するマイクロサービスです。  
sap-api-integrations-contact-reads には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-contact-reads は、オンプレミス版である（＝クラウド版ではない）SAPC4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/contact/overview  

## 動作環境
sap-api-integrations-contact-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。   
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。   
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須） 

## クラウド環境での利用  
sap-api-integrations-contact-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-contact-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/contact/overview 
* APIサービス名(=baseURL): c4codataapi

## 本レポジトリ に 含まれる API名
sap-api-integrations-contact-reads には、次の API をコールするためのリソースが含まれています。  

* ContactCollection（コンタクト - コンタクト）※コンタクトの詳細データを取得するために、ToContactIsContactPersonFor、ToCorporateAccountと合わせて利用されます。
* ToContactIsContactPersonFor（コンタクト - コンタクトパーソン ※To）
* ToCorpotrateAccount（コンタクト - 会社アカウント ※To）
* IndividualCustomerCollection（コンタクト - 個人顧客）※個人顧客の詳細データを取得するために、ToIndividualCustomerAddressと合わせて利用されます。
* ToIndividualCustomerAddress（コンタクト - 個人顧客住所 ※To）
* ContactIsContactPersonForCollection（コンタクト - コンタクトパーソン）
* CorporateAccountCollection（コンタクト - 会社アカウント）

## API への 値入力条件 の 初期値
sap-api-integrations-contact-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト
* inoutSDC.ContactCollection.ContactID（コンタクトID）
* inoutSDC.ContactCollection.ContactIsContactPersonFor.CorporateAccount.IndividualCustomerCollection.CustomerID（顧客ID）
* inoutSDC.ContactCollection.ContactIsContactPersonFor.AccountFormattedName（アカウント名）
* inoutSDC.ContactCollection.ContactIsContactPersonFor.CorporateAccount.AccountID（アカウントID）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"ContactCollection" が指定されています。    
  
```
	"api_schema": "ContactContactCollection",
	"accepter": ["ContactCollection"],
	"contact_code": "1000476",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "ContactContactCollection",
	"accepter": ["All"],
	"contact_code": "1000476",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
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
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP コンタクト  の コンタクトデータ が取得された結果の JSON の例です。  
以下の項目のうち、"ObjectID" ～ "CorporateAccount" は、/SAP_API_Output_Formatter/type.go 内 の Type ContactCollection {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-contact-reads/SAP_API_Caller/caller.go#L53",
	"function": "sap-api-integrations-contact-reads/SAP_API_Caller.(*SAPAPICaller).ContactCollection",
	"level": "INFO",
	"message": [
		{
			"ObjectID": "00163E038C2E1EE299C1BB0BE93B6F9B",
			"ContactID": "1000476",
			"ContactUUID": "00163E03-8C2E-1EE2-99C1-BB0BE93B6F9B",
			"ExternalID": "",
			"ExternalSystem": "",
			"StatusCode": "4",
			"StatusCodeText": "Obsolete",
			"TitleCode": "",
			"TitleCodeText": "",
			"AcademicTitleCode": "",
			"AcademicTitleCodeText": "",
			"AdditionalAcademicTitleCode": "",
			"AdditionalAcademicTitleCodeText": "",
			"NamePrefixCode": "",
			"NamePrefixCodeText": "",
			"FirstName": "Peter",
			"LastName": "Gamoff",
			"AdditionalFamilyName": "",
			"Initials": "",
			"MiddleName": "",
			"Name": "Peter Gamoff",
			"GenderCode": "0",
			"GenderCodeText": "Gender not known",
			"MaritalStatusCode": "",
			"MaritalStatusCodeText": "",
			"LanguageCode": "",
			"LanguageCodeText": "",
			"NickName": "",
			"BirthDate": "",
			"BirthName": "",
			"ContactPermissionCode": "",
			"ContactPermissionCodeText": "",
			"ProfessionCode": "",
			"ProfessionCodeText": "",
			"PerceptionOfCompanyCode": "",
			"PerceptionOfCompanyCodeText": "",
			"DeviatingFullName": "",
			"AccountID": "1001",
			"AccountUUID": "00163E03-8C2E-1EE2-99C1-BB0BE9398F9B",
			"AccountFormattedName": "Porter LLC",
			"Building": "",
			"Floor": "",
			"Room": "",
			"JobTitle": "",
			"FunctionCode": "0002",
			"FunctionCodeText": "Purchasing Manager",
			"DepartmentCode": "0001",
			"DepartmentCodeText": "Purchasing Dept",
			"Department": "",
			"VIPContactCode": "",
			"VIPContactCodeText": "",
			"Phone": "+1 (865) 813-0643",
			"Mobile": "+1 (712) 242-5596",
			"Fax": "+1 (573) 598-7485",
			"Email": "doug.gamoff@3M.com",
			"EmailInvalidIndicator": false,
			"BestReachedByCode": "",
			"BestReachedByCodeText": "",
			"FormattedPostalAddressDescription": "5 Centerstage / St. Paul MN 55144 / US",
			"BusinessAddressCountryCode": "US",
			"BusinessAddressCountryCodeText": "United States",
			"BusinessAddressStateCodeTextUpdatable": "Minnesota",
			"BusinessAddressHouseNumber": "5",
			"BusinessAddressStreet": "Centerstage",
			"BusinessAddressCity": "St. Paul",
			"BusinessAddressStreetPostalCode": "55144",
			"BusinessAddressStateCode": "MN",
			"BusinessAddressStateCodeText": "Minnesota",
			"CreationOn": "2013-01-24T19:09:19+09:00",
			"CreatedBy": "Eddie Smoke",
			"CreatedByIdentityUUID": "00163E03-A070-1EE2-88BA-39BD20F290B5",
			"ChangedOn": "2020-04-09T01:11:03+09:00",
			"ChangedBy": "Robert Mark",
			"ChangedByIdentityUUID": "00163E0C-84E2-1ED7-80DF-DF0660BB2577",
			"ContactOwnerID": "",
			"ContactOwnerUUID": "",
			"NormalisedPhone": "+18658130643",
			"NormalisedMobile": "+17122425596",
			"EntityLastChangedOn": "2020-04-09T01:11:03+09:00",
			"ContactIsContactPersonFor": "https://sandbox.api.sap.com/sap/c4c/odata/v1/c4codataapi/ContactCollection('00163E038C2E1EE299C1BB0BE93B6F9B')/ContactIsContactPersonFor",
			"CorporateAccount": "https://sandbox.api.sap.com/sap/c4c/odata/v1/c4codataapi/ContactCollection('00163E038C2E1EE299C1BB0BE93B6F9B')/CorporateAccount"
		}
	],
	"time": "2022-05-02T17:39:37+09:00"
}

```