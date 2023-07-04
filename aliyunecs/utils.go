package aliyunecs

import (
	"crypto/rand"
	"errors"
	mrand "math/rand"
	"strconv"
	"strings"
	"time"
)

var (
	errInvalidRegion  = errors.New("invalid region specified")
	errNoVpcs         = errors.New("No VPCs found in region")
	errMachineFailure = errors.New("Machine failed to start")
	errNoIP           = errors.New("No IP Address associated with the instance")
	errComplete       = errors.New("Complete")
)

const defaultUbuntuImageID = "ubuntu_16_0402_64_20G_alibase_20171227.vhd"
const defaultUbuntuImagePrefix = "ubuntu_16_0402_64"

func validateECSRegion(region string) (Region, error) {
	for _, v := range validRegions {
		if v == Region(region) {
			return v, nil
		}
	}

	return "", errInvalidRegion
}

const digitals = "0123456789"
const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const specialChars = "()`~!@#$%^&*-+=|{}[]:;'<>,.?/"
const dictionary = digitals + alphabet + specialChars
const tokenDictionary = "_0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const paswordLen = 16

func randomPassword() string {
	var bytes = make([]byte, paswordLen)
	rand.Read(bytes)
	for k, v := range bytes {
		var ch byte

		switch k {
		case 0:
			ch = alphabet[v%byte(len(alphabet))]
		case 1:
			ch = digitals[v%byte(len(digitals))]
		case 2:
			ch = specialChars[v%byte(len(specialChars))]
		default:
			ch = dictionary[v%byte(len(dictionary))]
		}
		bytes[k] = ch
	}
	return string(bytes)
}

func isUbuntuImage(image string) bool {
	return strings.HasPrefix(image, "ubuntu")
}

func SplitPortProto(raw string) (port int, protocol string, err error) {
	parts := strings.SplitN(raw, "/", 2)
	out, err := strconv.Atoi(parts[0])
	if err != nil {
		return 22, "tcp", err
	}
	if len(parts) == 1 {
		return out, "tcp", nil
	}
	if parts[1] != "tcp" && parts[1] != "udp" && parts[1] != "all" && parts[1] != "icmp" && parts[1] != "gre" {
		// If the format passed in does not match the ecs communication protocol then the default is tcp
		return out, "tcp", nil
	}

	return out, parts[1], nil
}

// A PaginationResponse represents a response with pagination information
type PaginationResult struct {
	TotalCount int
	PageNumber int
	PageSize   int
}

type Pagination struct {
	PageNumber int
	PageSize   int
}

// NextPage gets the next page of the result set
func (r *PaginationResult) NextPage() *Pagination {
	if r.PageNumber*r.PageSize >= r.TotalCount {
		return nil
	}
	return &Pagination{PageNumber: r.PageNumber + 1, PageSize: r.PageSize}
}

var validRegions = []Region{
	Hangzhou, Qingdao, Beijing, Shenzhen, Hongkong, Shanghai, Zhangjiakou, Huhehaote,
	USWest1, USEast1,
	APNorthEast1, APSouthEast1, APSouthEast2, APSouthEast3, APSouthEast5,
	APSouth1,
	MEEast1,
	EUCentral1, EUWest1,
	ShenZhenFinance, ShanghaiFinance,
}

// CreateRandomString create random string
func CreateRandomString() string {
	b := make([]byte, 32)
	l := len(tokenDictionary)

	_, err := rand.Read(b)

	if err != nil {
		// fail back to insecure rand
		mrand.Seed(time.Now().UnixNano())
		for i := range b {
			b[i] = tokenDictionary[mrand.Int()%l]
		}
	} else {
		for i, v := range b {
			b[i] = dictionary[v%byte(l)]
		}
	}

	return string(b)
}
