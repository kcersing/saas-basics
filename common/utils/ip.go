// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"errors"
	"net"
)

func GetLocalIPv4() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, iface := range interfaces {
		if iface.Flags&net.FlagLoopback != net.FlagLoopback && iface.Flags&net.FlagUp != 0 {
			addrs, err := iface.Addrs()
			if err != nil {
				continue
			}

			for _, addr := range addrs {
				if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
					return ipNet.IP.String(), nil
				}
			}
		}
	}
	return "", errors.New("get local IP error")
}

func MustGetLocalIPv4() string {
	ipv4, err := GetLocalIPv4()
	if err != nil {
		panic("get local IP error")
	}
	return ipv4
}

// GetLocalIPv4Address 单网卡
func GetLocalIPv4Address() (string, error) {
	addr, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, add := range addr {
		if ipNet, ok := add.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			ipv4 := ipNet.IP.To4()
			if ipv4 != nil {
				return ipv4.String(), nil
			}
		}
	}
	return "", errors.New("not found ipv4 address")
}

// GetLocalIPv4AddressArr 多网卡模式
func GetLocalIPv4AddressArr() (ipAdds []string, err error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return ipAdds, errors.New("not found ipv4 address")
	}
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addr, _ := netInterfaces[i].Addrs()
			for _, address := range addr {
				if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
					if ipNet.IP.To4() != nil {
						ipAdds = append(ipAdds, ipNet.IP.String())
					}
				}
			}
		}
	}
	return ipAdds, errors.New("not found ipv4 address")
}
