package db

import (
	"time"

	"github.com/sociam/xray-archiver/pipeline/util"
)

// Range represents a range
type Range struct {
	Min int64 `json:"min"`
	Max int64 `json:"max"`
}

// App represents an app from the database
type App struct {
	ID   string  `json:"id"`
	Vers []int64 `json:"vers"`
}

// AppVersion represents all the information about an app version contained in
// the database
type AppVersion struct {
	ID              int64     `json:"id"`
	App             string    `json:"app"`
	Store           string    `json:"string"`
	Region          string    `json:"region"`
	Ver             string    `json:"ver"`
	ScreenFlags     int64     `json:"screenFlags"`
	StoreInfo       StoreInfo `json:"storeinfo"`
	Icon            string    `json:"icon"`
	Dev             Developer `json:"developer"`
	Hosts           []string  `json:"hosts"`
	Perms           []string  `json:"perms"`
	Packages        []string  `json:"packages"`
	IsAnalyzed      bool      `json:"isAnalyzed"`
	APKLocationUUID string    `json:"apkLocationUUID"`
	APKLocationPath string    `json:"apkLocationPath"`
	APKLocationRoot string    `json:"apkLocationRoot"`
}

// GenreStats represents a row from the Genre average statistics table
type GenreStats struct {
	Category  string  `json:"category"`
	HostCount int64   `json:"hostCount"`
	AppCount  int64   `json:"appCount"`
	GenreAvg  float64 `json:"genreAvg"`
}

// CompanyCoverage represents a row from the company_app_coverage statistics
// table found in the XRay DB.
type CompanyCoverage struct {
	Company     string  `json:"company"`
	Type        string  `json:"type"`
	AppCount    int64   `json:"appCount"`
	TotalApps   int64   `json:"totalApps"`
	CompanyFreq float64 `json:"companyFreq"`
}

// CompanyGenreCoverage Stats for company Genre Coverage.
type CompanyGenreCoverage struct {
	Company      string  `json:"company"`
	CompanyCount int64   `json:"companyCount"`
	Genre        string  `json:"genre"`
	GenreTotal   int64   `json:"genreTotal"`
	CoveragePct  float64 `json:"companyPct"`
}

// CompanyTypeCoverage represents a row from the app_type_coverage statistics
// table found in the XRay DB
type CompanyTypeCoverage struct {
	Type      string  `json:"type"`
	AppCount  int64   `json:"appCount"`
	TotalApps int64   `json:"totalApps"`
	TypeFreq  float64 `json:"typeFreq"`
}

// AltApp represents Alternative App title and playstore url. Used in API
type AltApp struct {
	AltAppTitle     string `json:"altAppTitle"`
	AltToURL        string `json:"altToURL"`
	GPlayURL        string `json:"gPlayURL"`
	GPlayID         string `json:"gPlayID"`
	IconURL         string `json:"iconURL"`
	OfficialSiteURL string `json:"officialSiteURL"`
	IsScraped       bool   `json:"isScraped"`
}

// StoreInfo represents the information contained about an app in its respective
// store
type StoreInfo interface{}

// PlayStoreInfo represents the data contained in the google play store
type PlayStoreInfo struct {
	Title         string    `json:"title"`
	Summary       string    `json:"summary"` //TODO:Omitempty for null vlaues?
	Description   string    `json:"description"`
	StoreURL      string    `json:"storeURL"`
	Price         string    `json:"price"`
	Free          bool      `json:"free"`
	Rating        string    `json:"rating"`
	NumReviews    int64     `json:"numReviews"`
	Genre         string    `json:"genre"`
	FamilyGenre   string    `json:"familyGenre"`
	Installs      Range     `json:"installs"`
	Developer     int64     `json:"developer"`
	Updated       time.Time `json:"updated"`
	AndroidVer    string    `json:"androidVer"`
	ContentRating string    `json:"contentRating"`
	Screenshots   []string  `json:"screenshots"`
	Video         string    `json:"video"`
	RecentChanges []string  `json:"recentChanges"`
	CrawlDate     time.Time `json:"crawlDate"`
	Permissions   []string  `json:"permissions"`
}

// Developer represents a developer from the database
type Developer struct {
	Emails    []string `json:"emails"`
	Name      string   `json:"name"`
	StoreSite string   `json:"storeSite"`
	Site      string   `json:"site"`
}

// AppStub represents the 'stub' of app data located in the Database.
// it is formed of only the title given to the app by the developer
// and the App ID that it can be uniquely identified as.
type AppStub struct {
	Title string `json:"title"`
	App   string `json:"app"`
}

// Company represents a company from the database
type Company struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Hosts        []string `json:"nosts"`
	Founded      string   `json:"founded"`
	Acquired     string   `json:"acquired"`
	Type         []string `json:"type"`
	TypeTag      string   `json:"typeTag"`
	Jurisdiction string   `json:"jurisdiction"`
	Parent       string   `json:"parent"`
	Capital      string   `json:"capital"`
	Equity       string   `json:"equity"`
	Size         Range    `json:"size"`
	DataSources  []string `json:"dataSources"`
	Description  []string `json:"description"`
}

// UtilApp creates a *util.App from an AppVersion
func (a AppVersion) UtilApp() *util.App {
	return util.NewApp(a.ID, a.App, a.Store, a.Region, a.Ver, a.APKLocationPath, a.APKLocationRoot, a.APKLocationUUID)
}

// AppHostRecord holds app_host data from the xray DB
type AppHostRecord struct {
	ID        int64    `json:"id"`
	HostNames []string `json:"hostnames"`
}

// TrackerMapperRequest holds the data used in requests to the OxfordHCC TrackerMapper API.
type TrackerMapperRequest struct {
	HostNames []string `json:"host_names"`
}

// TrackerMapperCompany holds the data requested from the OxfordHCC TrackerMapper API.
type TrackerMapperCompany struct {
	HostName    string   `json:"hostName"`
	HostID      int64    `json:"hostID"`
	CompanyName string   `json:"companyName"`
	CompanyID   int64    `json:"companyID"`
	Locale      string   `json:"locale"`
	Categories  []string `json:"categories"`
}

// CompanyNames represents the json structure used to send company names selected from the DB via the rest API.
type CompanyNames struct {
	CompanyNames []string `json:"companyNames"`
}

// CompanyAssociations holds a record of from the DB describing the associations between a single company and various apps/devices/websites.
type CompanyAssociations struct {
	CompanyName            string  `json:"companyName"`
	AssociatedAppIDs       []int64 `json:"associatedAppIDs"`
	AssociatedIoTDeviceIDs []int64 `json:"associatedIoTDeviceIDs"`
	AssociatedWebsiteIDs   []int64 `json:"associatedWebsiteIDs"`
}

// AssociatedCompany holds the name of a company and the number of associations that it has.
type AssociatedCompany struct {
	CompanyName          string `json:"companyName"`
	NumberOfAssociations int64  `json:"numberOfAssociations"`
}

// AppAssociations holds the result of a query to the DB for all company names with an app association.
type AppAssociations struct {
	AppVersionID        int64               `json:"appVersionID"`
	AssociatedCompanies []AssociatedCompany `json:"associatedCompanies"`
}

// APIRequestError holds Error information sent via the API should na endpoint fail.
type APIRequestError struct {
	ErrorType    string `json:"ErrorType"`
	ErrorMessage string `json:"ErrorMsg"`
	APIRequest   string `json:"APIRequest"`
}
