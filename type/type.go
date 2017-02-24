package eureka

import "github.com/ibrokethecloud/go-utils"
import "os"

type RegisterationInfo struct {

  Instance InstanceInfo `json:"instance"`
}

type InstanceInfo struct {
  Hostname string `json:"hostName"`
  App string `json:"app"`
  VipAddress  string `json:"vipAddress"`
  SecureVipAddress  string  `json:"secureVipAddress"`
  IpAddress string `json:"ipAddr"`
  Status string `json:"status"`
  Port  string  `json:"port"`
  SecurePort string `json:"securePort"`
  HealthCheckUrl string `json:"healthCheckUrl"`
  StatusPageUrl string  `json:"statusPageUrl"`
  HomePageUrl  string `json:"homePageUrl"`
  DataCenter DataCenterInfo `json:"dataCenterInfo"`
}

type DataCenterInfo struct{
  Class string  `json:"@class"`
  Name string `json:"name"`
}

var ContainerIp = utils.GetIP()

var DefaultDataCenterInfo = DataCenterInfo {
  Class: "com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo",
  Name: "Internal",
}

var DefaultInstanceInfo = InstanceInfo {
  Hostname: utils.GetHostName(),
  App:  os.Getenv("APP_NAME"),
  VipAddress: ContainerIp,
  SecureVipAddress: ContainerIp,
  IpAddress:  ContainerIp,
  Status: "STARTING",
  Port: os.Getenv("PORT"),
  SecurePort: "8443",
  HealthCheckUrl: "http://"+ContainerIp+":"+os.Getenv("PORT")+"/health",
  StatusPageUrl: "http://"+ContainerIp+":"+os.Getenv("PORT")+"/status",
  HomePageUrl: "http://"+ContainerIp+":"+os.Getenv("PORT")+"/",
  DataCenter: DefaultDataCenterInfo,
}

var DefaultRegisterationInfo = RegisterationInfo {
  Instance: DefaultInstanceInfo,
}


/*  {
      "instance": {
          "hostName": "WKS-SOF-L011",
          "app": "com.automationrhapsody.eureka.app",
          "vipAddress": "com.automationrhapsody.eureka.app",
          "secureVipAddress": "com.automationrhapsody.eureka.app"
          "ipAddr": "10.0.0.10",
          "status": "STARTING",
          "port": {"$": "8080", "@enabled": "true"},
          "securePort": {"$": "8443", "@enabled": "true"},
          "healthCheckUrl": "http://WKS-SOF-L011:8080/healthcheck",
          "statusPageUrl": "http://WKS-SOF-L011:8080/status",
          "homePageUrl": "http://WKS-SOF-L011:8080",
          "dataCenterInfo": {
              "@class": "com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo",
              "name": "MyOwn"
          },
      }
  } */