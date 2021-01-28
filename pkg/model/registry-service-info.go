package model

type RegisteredServices struct {
	Application RegisteredApp `json:"applications"`
}

type RegisteredApp struct {
	Version string                `json:"versions__delta"`
	Hash    string                `json:"apps__hashcode"`
	Apps    []RegistryApplication `json:"application"`
}

type RegistryApplication struct {
	Name      string             `json:"name"`
	Instances []RegistryInstance `json:"instance"`
}

type RegistryInstance struct {
	InstanceId       string       `json:"instanceId"`
	Hostname         string       `json:"hostname"`
	App              string       `json:"app"`
	IpAddress        string       `json:"ipAddr"`
	Status           string       `json:"status"`
	OverriddenStatus string       `json:"overriddenStatus"`
	Port             RegisterInfo `json:"port"`
	SecurePort       RegisterInfo `json:"securePort"`
	CountryId        int          `json:"countryId"`
}

type RegisterInfo struct {
	Value    int    `json:"$"`
	IsActive string `json:"@enabled"`
}
