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
	Description   string `json:"description"`
	BackupType    string `json:"backupType"`
	RunNow        bool   `json:"runNow"`
	SelectedItems []struct {
		Type    string   `json:"type"`
		Folders []string `json:"folders"`
		ID      string   `json:"id"`
		Links   struct {
		} `json:"_links"`
		Mailbox        bool `json:"mailbox"`
		OneDrive       bool `json:"oneDrive"`
		ArchiveMailbox bool `json:"archiveMailbox"`
		Sites          bool `json:"sites"`
		Teams          bool `json:"teams"`
		TeamsChats     bool `json:"teamsChats"`
		Site           struct {
			ID                  string `json:"id"`
			URL                 string `json:"url"`
			ParentURL           string `json:"parentUrl"`
			Name                string `json:"name"`
			IsCloud             bool   `json:"isCloud"`
			IsPersonal          bool   `json:"isPersonal"`
			Title               string `json:"title"`
			SiteCollectionError string `json:"siteCollectionError"`
			Links               struct {
			} `json:"_links"`
		} `json:"site"`
		Team struct {
			ID          string `json:"id"`
			DisplayName string `json:"displayName"`
			Description string `json:"description"`
			Mail        string `json:"mail"`
			Links       struct {
			} `json:"_links"`
		} `json:"team"`
		Chats bool `json:"chats"`
		User  struct {
			ID           string `json:"id"`
			DisplayName  string `json:"displayName"`
			Name         string `json:"name"`
			Type         string `json:"type"`
			LocationType string `json:"locationType"`
			Links        struct {
			} `json:"_links"`
		} `json:"user"`
		PersonalSite bool `json:"personalSite"`
		Group        struct {
			ID           string `json:"id"`
			DisplayName  string `json:"displayName"`
			Name         string `json:"name"`
			ManagedBy    string `json:"managedBy"`
			Site         string `json:"site"`
			Type         string `json:"type"`
			LocationType string `json:"locationType"`
			Links        struct {
			} `json:"_links"`
		} `json:"group"`
		Members              bool `json:"members"`
		MemberMailbox        bool `json:"memberMailbox"`
		MemberArchiveMailbox bool `json:"memberArchiveMailbox"`
		MemberOnedrive       bool `json:"memberOnedrive"`
		MemberSite           bool `json:"memberSite"`
		GroupSite            bool `json:"groupSite"`
	} `json:"selectedItems"`
	ExcludedItems []struct {
		Type    string   `json:"type"`
		Folders []string `json:"folders"`
		ID      string   `json:"id"`
		Links   struct {
		} `json:"_links"`
		Mailbox        bool `json:"mailbox"`
		OneDrive       bool `json:"oneDrive"`
		ArchiveMailbox bool `json:"archiveMailbox"`
		Sites          bool `json:"sites"`
		Teams          bool `json:"teams"`
		TeamsChats     bool `json:"teamsChats"`
		Site           struct {
			ID                  string `json:"id"`
			URL                 string `json:"url"`
			ParentURL           string `json:"parentUrl"`
			Name                string `json:"name"`
			IsCloud             bool   `json:"isCloud"`
			IsPersonal          bool   `json:"isPersonal"`
			Title               string `json:"title"`
			SiteCollectionError string `json:"siteCollectionError"`
			Links               struct {
			} `json:"_links"`
		} `json:"site"`
		Team struct {
			ID          string `json:"id"`
			DisplayName string `json:"displayName"`
			Description string `json:"description"`
			Mail        string `json:"mail"`
			Links       struct {
			} `json:"_links"`
		} `json:"team"`
		Chats bool `json:"chats"`
		User  struct {
			ID           string `json:"id"`
			DisplayName  string `json:"displayName"`
			Name         string `json:"name"`
			Type         string `json:"type"`
			LocationType string `json:"locationType"`
			Links        struct {
			} `json:"_links"`
		} `json:"user"`
		PersonalSite bool `json:"personalSite"`
		Group        struct {
			ID           string `json:"id"`
			DisplayName  string `json:"displayName"`
			Name         string `json:"name"`
			ManagedBy    string `json:"managedBy"`
			Site         string `json:"site"`
			Type         string `json:"type"`
			LocationType string `json:"locationType"`
			Links        struct {
			} `json:"_links"`
		} `json:"group"`
		Members              bool `json:"members"`
		MemberMailbox        bool `json:"memberMailbox"`
		MemberArchiveMailbox bool `json:"memberArchiveMailbox"`
		MemberOnedrive       bool `json:"memberOnedrive"`
		MemberSite           bool `json:"memberSite"`
		GroupSite            bool `json:"groupSite"`
	} `json:"excludedItems"`
	SchedulePolicy struct {
		ScheduleEnabled      bool `json:"scheduleEnabled"`
		BackupWindowEnabled  bool `json:"backupWindowEnabled"`
		BackupWindowSettings struct {
			BackupWindow []bool `json:"backupWindow"`
			MinuteOffset int    `json:"minuteOffset"`
		} `json:"backupWindowSettings"`
		PeriodicallyWindowSettings struct {
			BackupWindow []bool `json:"backupWindow"`
			MinuteOffset int    `json:"minuteOffset"`
		} `json:"periodicallyWindowSettings"`
		PeriodicallyOffsetMinutes int    `json:"periodicallyOffsetMinutes"`
		Type                      string `json:"type"`
		PeriodicallyEvery         string `json:"periodicallyEvery"`
		DailyType                 string `json:"dailyType"`
		DailyTime                 string `json:"dailyTime"`
		RetryEnabled              bool   `json:"retryEnabled"`
		RetryNumber               int    `json:"retryNumber"`
		RetryWaitInterval         int    `json:"retryWaitInterval"`
	} `json:"schedulePolicy"`
	ID           string    `json:"id"`
	RepositoryID string    `json:"repositoryId"`
	Name         string    `json:"name"`
	LastRun      time.Time `json:"lastRun"`
	NextRun      time.Time `json:"nextRun"`
	IsEnabled    bool      `json:"isEnabled"`
	LastStatus   string    `json:"lastStatus"`
	Links        struct {
	} `json:"_links"`
}

// Structure JSON
// Reponse GET /v6/JobSessions
type vbo365JobSessions struct {
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
	TotalCount int `json:"totalCount"`
	Links      struct {
	} `json:"_links"`
	Results []struct {
		ID           string    `json:"id"`
		Details      string    `json:"details"`
		CreationTime time.Time `json:"creationTime"`
		EndTime      time.Time `json:"endTime"`
		RetryCount   int       `json:"retryCount"`
		Progress     int       `json:"progress"`
		Status       string    `json:"status"`
		Statistics   struct {
			ProcessingRateBytesPS int    `json:"processingRateBytesPS"`
			ProcessingRateItemsPS int    `json:"processingRateItemsPS"`
			ReadRateBytesPS       int    `json:"readRateBytesPS"`
			WriteRateBytesPS      int    `json:"writeRateBytesPS"`
			TransferredDataBytes  int    `json:"transferredDataBytes"`
			ProcessedObjects      int    `json:"processedObjects"`
			Bottleneck            string `json:"bottleneck"`
		} `json:"statistics"`
		Links struct {
		} `json:"_links"`
	} `json:"results"`
	SetID string `json:"setId"`
}

// Structure JSON
// Reponse GET /v6/RestorePortalSettings
type vbo365RestorePortalSettings struct {
	ApplicationID    string `json:"applicationId"`
	MsalAuthorityURI string `json:"msalAuthorityUri"`
	IsEnabled        bool   `json:"isEnabled"`
}

// Structure JSON
// Reponse GET /v6/RbacRoles
type vbo365RbacRoles []struct {
	ID             string `json:"id"`
	OrganizationID string `json:"organizationId"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	RoleType       string `json:"roleType"`
	Links          struct {
	} `json:"_links"`
}

// Structure JSON
// Reponse GET /v6/RbacRoles/{roleId}/operators
type vbo365RbacRoleOperators []struct {
	Type string `json:"type"`
	User struct {
		ID          string `json:"id"`
		DisplayName string `json:"displayName"`
		Name        string `json:"name"`
		Type        string `json:"type"`
	} `json:"user"`
	ID    string `json:"id"`
	Links struct {
	} `json:"_links"`
	Site struct {
		ID         string `json:"id"`
		URL        string `json:"url"`
		Title      string `json:"title"`
		ParentURL  string `json:"parentUrl"`
		IsCloud    bool   `json:"isCloud"`
		IsPersonal bool   `json:"isPersonal"`
	} `json:"site"`
	Group struct {
		ID          string `json:"id"`
		DisplayName string `json:"displayName"`
		Name        string `json:"name"`
		Type        string `json:"type"`
	} `json:"group"`
}

// Structure JSON
// Reponse GET /v6/RbacRoles/{roleId}/selectedItems
type vbo365RbacSelectedItems []struct {
	Type string `json:"type"`
	User struct {
		ID          string `json:"id"`
		DisplayName string `json:"displayName"`
		Name        string `json:"name"`
		Type        string `json:"type"`
	} `json:"user"`
	ID    string `json:"id"`
	Links struct {
	} `json:"_links"`
	Site struct {
		ID         string `json:"id"`
		URL        string `json:"url"`
		Title      string `json:"title"`
		ParentURL  string `json:"parentUrl"`
		IsCloud    bool   `json:"isCloud"`
		IsPersonal bool   `json:"isPersonal"`
	} `json:"site"`
	Group struct {
		ID          string `json:"id"`
		DisplayName string `json:"displayName"`
		Name        string `json:"name"`
		Type        string `json:"type"`
	} `json:"group"`
}
