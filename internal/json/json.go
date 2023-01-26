package json

import "time"

// Structure JSON
// Reponse POST /v6/Token
type vbo365Token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Issued       string `json:".issued"`
	Expires      string `json:".expires"`
}

// Structure JSON
// Reponse GET /v6/Organizations
type vbo365Organizations []struct {
	IsTeamsOnline          bool `json:"isTeamsOnline"`
	IsTeamsChatsOnline     bool `json:"isTeamsChatsOnline"`
	ExchangeOnlineSettings struct {
		UseApplicationOnlyAuth           bool   `json:"useApplicationOnlyAuth"`
		Account                          string `json:"account"`
		GrantAdminAccess                 bool   `json:"grantAdminAccess"`
		UseMfa                           bool   `json:"useMfa"`
		ApplicationID                    string `json:"applicationId"`
		ApplicationCertificateThumbprint string `json:"applicationCertificateThumbprint"`
	} `json:"exchangeOnlineSettings"`
	SharePointOnlineSettings struct {
		UseApplicationOnlyAuth           bool   `json:"useApplicationOnlyAuth"`
		OfficeOrganizationName           string `json:"officeOrganizationName"`
		SharePointSaveAllWebParts        bool   `json:"sharePointSaveAllWebParts"`
		Account                          string `json:"account"`
		GrantAdminAccess                 bool   `json:"grantAdminAccess"`
		UseMfa                           bool   `json:"useMfa"`
		ApplicationID                    string `json:"applicationId"`
		ApplicationCertificateThumbprint string `json:"applicationCertificateThumbprint"`
	} `json:"sharePointOnlineSettings"`
	IsExchangeOnline   bool      `json:"isExchangeOnline"`
	IsSharePointOnline bool      `json:"isSharePointOnline"`
	Type               string    `json:"type"`
	Region             string    `json:"region"`
	ID                 string    `json:"id"`
	Name               string    `json:"name"`
	OfficeName         string    `json:"officeName"`
	IsBackedup         bool      `json:"isBackedup"`
	FirstBackuptime    time.Time `json:"firstBackuptime"`
	LastBackuptime     time.Time `json:"lastBackuptime"`
	Links              struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Jobs struct {
			Href string `json:"href"`
		} `json:"jobs"`
		Groups struct {
			Href string `json:"href"`
		} `json:"groups"`
		Users struct {
			Href string `json:"href"`
		} `json:"users"`
		Sites struct {
			Href string `json:"href"`
		} `json:"sites"`
		Teams struct {
			Href string `json:"href"`
		} `json:"teams"`
		UsedRepositories struct {
			Href string `json:"href"`
		} `json:"usedRepositories"`
		RbacRoles struct {
			Href string `json:"href"`
		} `json:"rbacRoles"`
	} `json:"_links"`
	Actions struct {
	} `json:"_actions"`
}

// Structure JSON
// Reponse GET /v6/ServiceInstance
type vbo365ServiceInstance struct {
	InstallationID string `json:"installationId"`
	Version        string `json:"version"`
}

// Structure JSON
// Reponse GET /v6/Organizations/{organizationID}/LicensingInformation
type vbo365LicensingInformation struct {
	LicensedUsers int `json:"licensedUsers"`
	NewUsers      int `json:"newUsers"`
}

// Structure JSON
// Reponse GET /v6/LicensedUsers
type vbo365LicensedUsers struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Links  struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Next struct {
			Href string `json:"href"`
		} `json:"next"`
	} `json:"_links"`
	Results []struct {
		ID                     string    `json:"id"`
		Name                   string    `json:"name"`
		IsBackedUp             bool      `json:"isBackedUp"`
		LastBackupDate         time.Time `json:"lastBackupDate"`
		LicenseState           string    `json:"licenseState"`
		OrganizationID         string    `json:"organizationId"`
		BackedUpOrganizationID string    `json:"backedUpOrganizationId"`
		OrganizationName       string    `json:"organizationName"`
		Links                  struct {
			Organization struct {
				Href string `json:"href"`
			} `json:"organization"`
		} `json:"_links"`
	} `json:"results"`
	SetID string `json:"setId"`
}

// Structure JSON
// Reponse GET /v6/BackupRepositories
type vbo365BackupRepositories []struct {
	ObjectStorageID                string `json:"objectStorageId,omitempty"`
	ObjectStorageCachePath         string `json:"objectStorageCachePath,omitempty"`
	ObjectStorageEncryptionEnabled bool   `json:"objectStorageEncryptionEnabled,omitempty"`
	EncryptionKeyID                string `json:"encryptionKeyId,omitempty"`
	IsOutOfSync                    bool   `json:"isOutOfSync"`
	CapacityBytes                  int64  `json:"capacityBytes"`
	FreeSpaceBytes                 int64  `json:"freeSpaceBytes"`
	ID                             string `json:"id"`
	Name                           string `json:"name"`
	Description                    string `json:"description"`
	RetentionType                  string `json:"retentionType"`
	RetentionPeriodType            string `json:"retentionPeriodType"`
	YearlyRetentionPeriod          string `json:"yearlyRetentionPeriod"`
	RetentionFrequencyType         string `json:"retentionFrequencyType"`
	MonthlyTime                    string `json:"monthlyTime,omitempty"`
	MonthlyDaynumber               string `json:"monthlyDaynumber,omitempty"`
	MonthlyDayofweek               string `json:"monthlyDayofweek,omitempty"`
	ProxyID                        string `json:"proxyId"`
	IsLongTerm                     bool   `json:"isLongTerm,omitempty"`
	Links                          struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Proxy struct {
			Href string `json:"href"`
		} `json:"proxy"`
	} `json:"_links"`
	Path      string `json:"path,omitempty"`
	DailyTime string `json:"dailyTime,omitempty"`
	DailyType string `json:"dailyType,omitempty"`
}

// Structure JSON
// Reponse GET /v6/objectstoragerepositories/{repositoryId}
type vbo365ObjectStorageRepositories struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	AccountID        string `json:"accountId"`
	SizeLimitEnabled bool   `json:"sizeLimitEnabled"`
	UsedSpaceBytes   int64  `json:"usedSpaceBytes"`
	Type             string `json:"type"`
	Links            struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Account struct {
			Href string `json:"href"`
		} `json:"account"`
		Container struct {
			Href string `json:"href"`
		} `json:"container"`
		Folder struct {
			Href string `json:"href"`
		} `json:"folder"`
	} `json:"_links"`
	UseArchiverAppliance bool `json:"useArchiverAppliance"`
	AzureContainer       struct {
		Name       string `json:"name"`
		RegionType string `json:"regionType"`
	} `json:"azureContainer"`
	AzureFolder string `json:"azureFolder"`
}

// Structure JSON
// Reponse GET /v6/Proxies
type vbo365Proxies []struct {
	Type                    string `json:"type"`
	UseInternetProxy        bool   `json:"useInternetProxy"`
	InternetProxyType       string `json:"internetProxyType"`
	ID                      string `json:"id"`
	HostName                string `json:"hostName"`
	Description             string `json:"description"`
	Port                    int    `json:"port"`
	ThreadsNumber           int    `json:"threadsNumber"`
	EnableNetworkThrottling bool   `json:"enableNetworkThrottling"`
	Status                  string `json:"status"`
	Links                   struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Repositories struct {
			Href string `json:"href"`
		} `json:"repositories"`
	} `json:"_links"`
}

// Structure JSON
// Reponse GET /v6/Jobs
type vbo365Jobs []struct {
	Description    string `json:"description"`
	BackupType     string `json:"backupType"`
	SchedulePolicy struct {
		ScheduleEnabled     bool   `json:"scheduleEnabled"`
		BackupWindowEnabled bool   `json:"backupWindowEnabled"`
		Type                string `json:"type"`
		DailyType           string `json:"dailyType"`
		DailyTime           string `json:"dailyTime"`
		RetryEnabled        bool   `json:"retryEnabled"`
		RetryNumber         int    `json:"retryNumber"`
		RetryWaitInterval   int    `json:"retryWaitInterval"`
	} `json:"schedulePolicy,omitempty"`
	ID           string    `json:"id"`
	RepositoryID string    `json:"repositoryId"`
	Name         string    `json:"name"`
	LastRun      time.Time `json:"lastRun"`
	NextRun      time.Time `json:"nextRun,omitempty"`
	IsEnabled    bool      `json:"isEnabled"`
	LastStatus   string    `json:"lastStatus"`
	Links        struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		CopyJob struct {
			Href string `json:"href"`
		} `json:"copyJob"`
		Organization struct {
			Href string `json:"href"`
		} `json:"organization"`
		BackupRepository struct {
			Href string `json:"href"`
		} `json:"backupRepository"`
		Jobsessions struct {
			Href string `json:"href"`
		} `json:"jobsessions"`
		ExcludedItems struct {
			Href string `json:"href"`
		} `json:"excludedItems"`
		SelectedItems struct {
			Href string `json:"href"`
		} `json:"selectedItems"`
	} `json:"_links"`
	SchedulePolicy0 struct {
		ScheduleEnabled     bool   `json:"scheduleEnabled"`
		BackupWindowEnabled bool   `json:"backupWindowEnabled"`
		Type                string `json:"type"`
		DailyType           string `json:"dailyType"`
		DailyTime           string `json:"dailyTime"`
		RetryEnabled        bool   `json:"retryEnabled"`
	} `json:"schedulePolicy,omitempty"`
	SchedulePolicy1 struct {
		ScheduleEnabled     bool `json:"scheduleEnabled"`
		BackupWindowEnabled bool `json:"backupWindowEnabled"`
		RetryEnabled        bool `json:"retryEnabled"`
		RetryNumber         int  `json:"retryNumber"`
		RetryWaitInterval   int  `json:"retryWaitInterval"`
	} `json:"schedulePolicy,omitempty"`
}
