package aliyunecs

import (
	"crypto/rand"
	"errors"
	"strconv"
	"strings"

	"github.com/denverdino/aliyungo/common"
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

func validateECSRegion(region string) (common.Region, error) {
	for _, v := range common.ValidRegions {
		if v == common.Region(region) {
			return v, nil
		}
	}

	return "", errInvalidRegion
}

const digitals = "0123456789"
const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const specialChars = "()`~!@#$%^&*-+=|{}[]:;'<>,.?/"
const dictionary = digitals + alphabet + specialChars
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
