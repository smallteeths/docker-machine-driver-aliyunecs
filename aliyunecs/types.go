package aliyunecs

import "fmt"

var RegionalDomainServices = []string{
	"ecs",
	"vpc",
	"slb",
	"pvtz",
}

var UnitRegions = map[Region]interface{}{
	Hangzhou: Hangzhou,
}

type DescribeEndpointArgs struct {
	Id          Region
	ServiceCode string
	Type        string
}

type BusinessInfo struct {
	Pack       string `json:"pack,omitempty"`
	ActivityId string `json:"activityId,omitempty"`
}

// xml
type Endpoints struct {
	Endpoint []Endpoint `xml:"Endpoint"`
}

type Endpoint struct {
	Name      string    `xml:"name,attr"`
	RegionIds RegionIds `xml:"RegionIds"`
	Products  Products  `xml:"Products"`
}

type RegionIds struct {
	RegionId string `xml:"RegionId"`
}

type Products struct {
	Product []Product `xml:"Product"`
}

type Product struct {
	ProductName string `xml:"ProductName"`
	DomainName  string `xml:"DomainName"`
}

type IpPermission struct {
	IpProtocol string
	FromPort   int
	ToPort     int
	IpRange    string
}

// Region represents ECS region
type Region string

// Constants of region definition
const (
	Hangzhou    = Region("cn-hangzhou")
	Qingdao     = Region("cn-qingdao")
	Beijing     = Region("cn-beijing")
	Hongkong    = Region("cn-hongkong")
	Shenzhen    = Region("cn-shenzhen")
	Shanghai    = Region("cn-shanghai")
	Zhangjiakou = Region("cn-zhangjiakou")
	Huhehaote   = Region("cn-huhehaote")

	APSouthEast1 = Region("ap-southeast-1")
	APNorthEast1 = Region("ap-northeast-1")
	APSouthEast2 = Region("ap-southeast-2")
	APSouthEast3 = Region("ap-southeast-3")
	APSouthEast5 = Region("ap-southeast-5")

	APSouth1 = Region("ap-south-1")

	USWest1 = Region("us-west-1")
	USEast1 = Region("us-east-1")

	MEEast1 = Region("me-east-1")

	EUCentral1 = Region("eu-central-1")
	EUWest1    = Region("eu-west-1")

	ShenZhenFinance = Region("cn-shenzhen-finance-1")
	ShanghaiFinance = Region("cn-shanghai-finance-1")
)

type BackendServerType struct {
	ServerId string
	Weight   int
	Type     string
}

type Response struct {
	RequestId string
}

type ErrorResponse struct {
	Response
	HostId  string
	Code    string
	Message string
}

// An Error represents a custom error for Aliyun API failure response
type Error struct {
	ErrorResponse
	StatusCode int //Status Code of HTTP Response
}

func (e *Error) Error() string {
	return fmt.Sprintf("Aliyun API Error: RequestId: %s Status Code: %d Code: %s Message: %s", e.RequestId, e.StatusCode, e.Code, e.Message)
}
