package main

import (
	"encoding/json"
	"os"
	"strings"
	"sync"
)

type hostInfo struct {
	Idc  string `json:"idc"`
	Name string `json:"name"`
}

var HostInfo hostInfo
var once = sync.Once{}

func init() {
	once.Do(func() {
		const filename = "/Users/cipher/Downloads/host.ini"
		if bytes, err := os.ReadFile(filename); err == nil {
			m := make(map[string]string)
			str := string(bytes)
			lines := strings.Split(str, "\n")
			for _, line := range lines {
				if len(line) > 0 {
					info := strings.Split(line, "=")
					m[info[0]] = info[1]
				}
			}
			marshal, _ := json.Marshal(m)
			json.Unmarshal(marshal, &HostInfo)
		}
	})
}
