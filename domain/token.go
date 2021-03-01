package domain

const HMAC_SAMPLE_SECRET = "hmacSampleSecret"

type Claims struct {
	CustomerId string   `json:"customer_id"`
	Accounts   []string `json:"accounts"`
	Username   string   `json:"username"`
	Expiry     int64    `json:"expiry"`
	Role       string   `json:"role"`
}

func (c Claims) IsUserRole() bool{
	if c.Role == "username"{
		return true
	}
	return false
}

func BuildClaimsFromJwtMapClaims(mapClaims jwt.MapClaims) (*Claims, error){

}