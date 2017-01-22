package gotodoist

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type TodoistAPI struct {
	Projects Projects
	Items    Items
}

type SyncObject struct {
	token      string
	url        string
	sync_token string
}

type Command struct {
	Type   string                 `json:"type"`
	UUID   string                 `json:"uuid"`
	TempID string                 `json:"temp_id"`
	Args   map[string]interface{} `json:"args"`
}

func (sync SyncObject) callWriteApi(commands []Command) (*http.Response, error) {
	sync_api_url := sync.url + "?token=" + sync.token
	buf := new(bytes.Buffer)
	e := json.NewEncoder(buf)
	e.Encode(commands)
	jsonStr := buf.String()
	sync_obj := url.Values{"commands": {jsonStr}}
	resp_new, err_new := http.PostForm(sync_api_url, sync_obj)
	return resp_new, err_new
}

func (sync SyncObject) callReadApi(resource_types []string) (*http.Response, error) {
	sync_api_url := sync.url + "?token=" + sync.token + "&sync_token=" + sync.sync_token
	buf := new(bytes.Buffer)
	e := json.NewEncoder(buf)
	e.Encode(resource_types)
	jsonStr := buf.String()
	sync_obj := url.Values{"resource_types": {jsonStr}}
	resp_new, err_new := http.PostForm(sync_api_url, sync_obj)
	return resp_new, err_new
}

func InitTodoistAPI(api_token string) TodoistAPI {
	sync_object := SyncObject{api_token, "https://todoist.com/API/v7/sync", "*"}
	projects := Projects{&sync_object}
	items := Items{&sync_object}
	todoist_api := TodoistAPI{projects, items}
	return todoist_api
}

func apiResponseToMap(response *http.Response) map[string]interface{} {
	fmt.Println(response)
	response_body, _ := ioutil.ReadAll(response.Body)
	response_body_map := make(map[string]interface{})
	json.Unmarshal(response_body, &response_body_map)
	response_and_status := make(map[string]interface{})
	response_and_status["body"] = response_body_map
	response_and_status["status"] = response.Status
	return response_and_status
}

func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
