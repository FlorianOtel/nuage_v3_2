package nuage_v3_2

type Enterprise struct {
	AllowAdvancedQOSConfiguration         bool     `json:"allowAdvancedQOSConfiguration,omitempty"`
	AllowedForwardingClasses              []string `json:"allowedForwardingClasses,omitempty"`
	AllowGatewayManagement                bool     `json:"allowGatewayManagement,omitempty"`
	AllowTrustedForwardingClass           bool     `json:"allowTrustedForwardingClass,omitempty"`
	AssociatedEnterpriseSecurityID        string   `json:"associatedEnterpriseSecurityID,omitempty"`
	AssociatedGroupKeyEncryptionProfileID string   `json:"associatedGroupKeyEncryptionProfileID,omitempty"`
	AssociatedKeyServerMonitorID          string   `json:"associatedKeyServerMonitorID,omitempty"`
	AvatarData                            string   `json:"avatarData,omitempty"`
	AvatarType                            []string `json:"avatarType,omitempty"`
	CustomerID                            int64    `json:"customerID,omitempty"`
	Description                           string   `json:"description"`
	DHCPLeaseinterval                     int      `json:"DHCPLeaseinterval,omitempty"`
	EncryptionManagementMode              string   `json:"encryptionManagementMode,omitempty"`
	EnterpriseProfileID                   string   `json:"enterpriseProfileID,omitempty"`
	FloatingIPsQuota                      int      `json:"floatingIPsQuota,omitempty"`
	FloatingIPsUsed                       int      `json:"floatingIPsUsed,omitempty"`
	LDAPAuthorizationEnabled              bool     `json:"LDAPAuthorizationEnabled,omitempty"`
	LDAPEnabled                           bool     `json:"LDAPEnabled,omitempty"`
	Name                                  string   `json:"name"`
	ReceiveMultiCastListID                string   `json:"receiveMultiCastListID,omitempty"`
	SendMultiCastListID                   string   `json:"sendMultiCastListID,omitempty"`
	CreationDate                          int64    `json:"creationDate,omitempty"`
	LastUpdatedBy                         string   `json:"lastUpdatedBy,omitempty"`
	LastUpdatedDate                       int64    `json:"lastUpdatedDate,omitempty"`
	Owner                                 string   `json:"owner,omitempty"`
	EntityScope                           string   `json:"entityScope,omitempty"`
	ExternalID                            string   `json:"externalID,omitempty"`
	ID                                    string   `json:"ID,omitempty"`
	ParentID                              string   `json:"parentID,omitempty"`
	ParentType                            string   `json:"parentType,omitempty"`
}

type EnterpriseSlice []Enterprise
