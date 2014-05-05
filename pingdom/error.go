// Steve Phillips / elimisteve
// 2014.05.05

package pingdom

type PingdomError struct {
	ErrorMessage string `json:"errormessage"`
	StatusCode   int    `json:"statuscode"`
	StatusDesc   string `json:"statusdesc"`
}
