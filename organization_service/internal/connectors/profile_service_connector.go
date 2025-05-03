package connectors

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func DoesAdminProfileExist(id uuid.UUID) bool {
	url := fmt.Sprintf("http://profileService.api:8080/admin-profile/%s", id)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("no connection to profile Service: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("unreadable body: %s", err)
	}

	var data map[string]interface{}
	jsonErr := json.Unmarshal([]byte(body), &data)
	return jsonErr == nil
}
