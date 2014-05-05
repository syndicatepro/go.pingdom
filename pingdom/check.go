// Steve Phillips / elimisteve
// 2014.05.05

package pingdom

const (
	baseApiUrl = "https://api.pingdom.com"

	urlChecks = baseApiUrl + "/api/2.0/checks"
)

type checksResponse struct {
	Checks []*Check      `json:"checks"`
	Error  *PingdomError `json:"error"`
}

type Check struct {
	Hostname         string `json:"hostname"`
	ID               int    `json:"id"`
	LastErrorTime    int    `json:"lasterrortime"`
	LastResponseTime int    `json:"lastresponsetime"`
	LastTestTime     int    `json:"lasttesttime"`
	Name             string `json:"name"`
	Resolution       int    `json:"resolution"`
	Status           string `json:"status"`
	Type             string `json:"type"`
}

func (c *Check) IsDown() bool {
	return c.Status == "down"
}
