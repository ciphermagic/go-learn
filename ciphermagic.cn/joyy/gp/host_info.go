package main

import (
	"encoding/json"
	"os"
	"strings"
)

var HostInfo hostInfo

type hostInfo struct {
	Idc  string `json:"idc"`
	Name string `json:"name"`
}

func init() {
	const filename = "/Users/cipher/Downloads/host.ini"
	if b, e := os.ReadFile(filename); e == nil {
		m := make(map[string]string)
		for _, line := range strings.Split(string(b), "\n") {
			if info := strings.Split(line, "="); len(info) == 2 {
				m[strings.TrimSpace(info[0])] = strings.TrimSpace(info[1])
			}
		}
		marshal, _ := json.Marshal(m)
		json.Unmarshal(marshal, &HostInfo)
	}
}
