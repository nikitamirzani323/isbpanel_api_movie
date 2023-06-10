package entities

type Model_providerslot struct {
	Providerslot_slug  string `json:"providerslot_slug"`
	Providerslot_name  string `json:"providerslot_name"`
	Providerslot_image string `json:"providerslot_image"`
	Providerslot_title string `json:"providerslot_title"`
	Providerslot_descp string `json:"providerslot_descp"`
}
type Model_prediksislot struct {
	Prediksislot_name           string `json:"prediksislot_name"`
	Prediksislot_image          string `json:"prediksislot_image"`
	Prediksislot_prediksi       int    `json:"prediksislot_prediksi"`
	Prediksislot_prediksi_class string `json:"prediksislot_prediksi_class"`
}
type Controller_providerslot struct {
	Providerslot_slug string `json:"providerslot_slug" `
}
type Controller_prediksislot struct {
	Providerslot_id    string `json:"providerslot_id" `
	Prediksislot_limit int    `json:"prediksislot_limit" `
}
