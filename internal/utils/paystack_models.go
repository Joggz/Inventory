package utils


type PaystackInitializeResponse struct {
	Status  bool   `json:"status"`
    Message string `json:"message"`
    Data    struct {
        AuthorizationURL string `json:"authorization_url"`
        AccessCode       string `json:"access_code"`
        Reference        string `json:"reference"`
    } `json:"data"`
}