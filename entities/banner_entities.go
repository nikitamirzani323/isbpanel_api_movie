package entities

type Model_banner struct {
	Banner_url        string `json:"banner_url"`
	Banner_urlwebsite string `json:"banner_urlwebsite"`
	Banner_posisi     string `json:"banner_posisi"`
	Banner_device     string `json:"banner_device"`
}
type Controller_banner struct {
	Client_Device string `json:"client_device"`
}
