package source

// 从CMDB或者其他的数据源中获取

import (
	"devops-dns-server/config"
	"github.com/levigross/grequests"
	"log"
)

type Data struct {
	Data       string `json:"data"`
	EntryType  string `json:"entryType,omitempty"`
	Error      string `json:"error,omitempty"`
	StatusCode int    `json:"statusCode,omitempty"`
}

func FromAPI(hostname string) string {
	resp, err := grequests.Get(config.GetConfig().String("fromAPI::url"), &grequests.RequestOptions{
		Params: map[string]string{"name": hostname},
	})
	if err != nil {
		log.Printf("访问API地址出错了: %s\n", err)
		return ""
	}
	var data Data
	err = resp.JSON(&data)
	if err != nil {
		log.Printf("解析API JSON出错了: %s\n", err)
		return ""
	}

	return data.Data
}
