package jiralert

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

// Grafana API Server
type Grafana struct {
	URL   string
	Token string
}

// Dashboard from Grafana
type Dashboard struct {
	Dashboard interface{} `json:"dashboard"`
}

// Snapshot from Grafana
type Snapshot struct {
	DeleteKey string `json:"deleteKey"`
	DeleteURL string `json:"deleteUrl"`
	Key       string `json:"key"`
	URL       string `json:"url"`
}

func (g *Grafana) dashboard(uid string) (*Dashboard, error) {
	req, err := http.NewRequest("GET", g.URL+"/api/dashboards/uid/"+uid, nil)
	req.Header.Add("Authorization", "Bearer "+g.Token)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New("HTTP Error " + string(res.StatusCode))
	}

	dashboard := new(Dashboard)

	err = json.NewDecoder(res.Body).Decode(dashboard)

	return dashboard, nil
}
func (g *Grafana) snapshot(dashboard *Dashboard) (*Snapshot, error) {
	buffer, err := json.Marshal(dashboard)

	req, err := http.NewRequest("POST", g.URL+"/api/snapshots", bytes.NewBuffer(buffer))

	req.Header.Add("Authorization", "Bearer "+g.Token)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	snapshot := new(Snapshot)

	json.NewDecoder(res.Body).Decode(snapshot)

	return snapshot, nil
}
