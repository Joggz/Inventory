package utils
import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateInvoiceRef(variantID int64) string {
	rand.Seed(time.Now().UnixNano())

	random := rand.Intn(900000) + 100000 // 6-digit random
	return fmt.Sprintf("inv-%d-%d", variantID, random)
}
