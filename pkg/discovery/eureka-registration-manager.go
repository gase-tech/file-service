package discovery

import (
	"encoding/json"
	"github.com/carlescere/scheduler"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/helper"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/model"
	"github.com/imminoglobulin/e-commerce-backend/file-service/pkg/store"
	"github.com/rs/zerolog/log"
	"os"
	"runtime"
	"strconv"
	"time"
)

type AppRegistrationBody struct {
	Instance InstanceDetails `json:"instance"`
}

type InstanceDetails struct {
	HostName         string         `json:"hostName"`
	App              string         `json:"app"`
	VipAddress       string         `json:"vipAddress"`
	SecureVipAddress string         `json:"secureVipAddress"`
	IpAddr           string         `json:"ipAddr"`
	Status           string         `json:"status"`
	Port             Port           `json:"port"`
	SecurePort       Port           `json:"securePort"`
	HealthCheckUrl   string         `json:"healthCheckUrl"`
	StatusPageUrl    string         `json:"statusPageUrl"`
	HomePageUrl      string         `json:"homePageUrl"`
	DataCenterInfo   DataCenterInfo `json:"dataCenterInfo"`
}
type Port struct {
	Port    string `json:"$"`
	Enabled string `json:"@enabled"`
}

type DataCenterInfo struct {
	Class string `json:"@class"`
	Name  string `json:"name"`
}

// This struct shall be responsible for manager to manage registration with Eureka
type EurekaRegistrationManager struct {
}

func (erm EurekaRegistrationManager) RegisterWithServiceRegistry(eurekaConfigs RegistrationVariables) {
	log.Info().Msg("Registering service with status : STARTING")
	body := erm.getBodyForEureka("STARTING")
	url := eurekaConfigs.ServiceRegistryURL + "apps/" + store.AppConfig.ApplicationName
	_, _ = helper.MakePostCall(url, body, nil, true)
	log.Info().Msg("Waiting for 10 seconds for application to start properly")
	time.Sleep(10 * time.Second)
	log.Info().Msg("Updating the status to : UP")
	bodyUP := erm.getBodyForEureka("UP")
	url = eurekaConfigs.ServiceRegistryURL + "apps/" + store.AppConfig.ApplicationName
	_, _ = helper.MakePostCall(url, bodyUP, nil, true)
}

func (erm EurekaRegistrationManager) SendHeartBeat(eurekaConfigs RegistrationVariables) {
	log.Info().Msg("In SendHeartBeat!")
	hostname, err := os.Hostname()
	if err != nil {
		log.Print("Error while getting hostname which shall be used as APP ID")
	}
	job := func() {
		url := eurekaConfigs.ServiceRegistryURL + "apps/" + store.AppConfig.ApplicationName + "/" + hostname
		_, _ = helper.MakePutCall(url, nil, nil, true)

		erm.StoreOtherMSInfo(eurekaConfigs)
	}
	// Run every 25 seconds but not now.
	_, _ = scheduler.Every(25).Seconds().Run(job)
	runtime.Goexit()
}

func (erm EurekaRegistrationManager) DeRegisterFromServiceRegistry(configs RegistrationVariables) {
	_, _ = helper.MakePostCall(configs.ServiceRegistryURL, nil, nil, true)
}

func (erm EurekaRegistrationManager) getBodyForEureka(status string) *AppRegistrationBody {
	httpPort := store.AppConfig.Port
	hostname, err := os.Hostname()
	if err != nil {
		log.Print("Enable to find hostname form OS, sending appname as host name")
	}

	ipAddress, err := helper.ExternalIP()
	if err != nil {
		log.Print("Enable to find IP address form OS")
	}

	port := Port{httpPort, "true"}
	securePort := Port{httpPort, strconv.FormatBool(store.AppConfig.RegistryConfig.UseSSL)}
	dataCenterInfo := DataCenterInfo{"com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo", "MyOwn"} // TODO: will be remove

	homePageUrl := "http://" + hostname + ":" + httpPort
	statusPageUrl := "http://" + hostname + ":" + httpPort + "/status"
	healthCheckUrl := "http://" + hostname + ":" + httpPort + "/healthcheck"

	instance := InstanceDetails{hostname, store.AppConfig.ApplicationName, store.AppConfig.ApplicationName, store.AppConfig.ApplicationName,
		ipAddress, status, port, securePort, healthCheckUrl, statusPageUrl, homePageUrl, dataCenterInfo}

	body := &AppRegistrationBody{instance}
	return body
}

func (erm EurekaRegistrationManager) StoreOtherMSInfo(configs RegistrationVariables) {
	url := configs.ServiceRegistryURL + "apps"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err, resp := helper.MakeGetCall(url, headers, nil, true)

	if err != nil {
		log.Error().Err(err).Msg("Can not send update apps info request")
	}

	var serviceInfo model.RegisteredServices
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&serviceInfo)

	if err != nil {
		log.Err(err).Msg("Error update apps info response convert to json ")
	}

	registeredServices := make([]store.ServiceApp, len(serviceInfo.Application.Apps))
	for i, app := range serviceInfo.Application.Apps {
		instanceUrls := make([]string, len(app.Instances))
		for j, instance := range app.Instances {
			instanceUrls[j] = "http://" + instance.Hostname + ":" + strconv.Itoa(instance.Port.Value)
		}

		registeredServices[i] = store.ServiceApp{
			Name:             app.Name,
			InstanceBaseUrls: instanceUrls,
		}
	}

	store.RegisteredServices = &store.ServiceInfos{
		Services: registeredServices,
	}
}
