package services

type ProductInfo struct {
	Name string
	Description string
	Variants []ProductVariant
}


type ProductVariant struct {
	Name	string
	SKU   string
}

