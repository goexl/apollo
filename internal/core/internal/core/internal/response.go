package internal

type Response struct {
	Appid          string `json:"appId"`
	Cluster        string `json:"cluster"`
	Namespace      string `json:"namespaceName"`
	Configurations struct {
		Content string `json:"content"`
	} `json:"configurations"`
	Key string `json:"releaseKey"`
}
