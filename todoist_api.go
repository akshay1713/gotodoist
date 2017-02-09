//Package gotodoist is a library for interacting with the Todoist api
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

//TodoistAPI holds everything needed to interact with this library.
//It is the gateway to this library. All external calls go through a variable of this type.
type TodoistAPI struct {
	Projects    Projects
	Items       Items
	sync_object *SyncObject
}

//SyncObject is used for directly interacting with the Sync API of Todoist.
type SyncObject struct {
	token       string
	url         string
	sync_token  string
	write_queue []Command
}

//Command is a type which is used for creating the parameters to be sent to the API.
type Command struct {
	Type   string                 `json:"type"`
	UUID   string                 `json:"uuid"`
	TempID string                 `json:"temp_id"`
	Args   map[string]interface{} `json:"args"`
}

//callWriteApi is used while writing resources.
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

//callReadApi is used while reading resources
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

//queueCommands queues the given commands to the write_queue
func (sync *SyncObject) queueCommands(commands []Command) {
	current_commands := sync.write_queue
	updated_commands := append(current_commands, commands...)
	sync.write_queue = updated_commands
}

//Commit executes all the commands currently queued in the write_queue type of SyncObject.
//This is used to take advantage of the batched api calls feature provided by the Todoist API.
//Multiple commands can be executed in a single HTTP call
func (todoist_api *TodoistAPI) Commit() (map[string]interface{}, error) {
	commands := todoist_api.sync_object.write_queue
	response, err := todoist_api.sync_object.callWriteApi(commands)
	if err == nil {
		commands = []Command{}
		todoist_api.sync_object.write_queue = commands
	}
	response_map := apiResponseToMap(response)
	return response_map, err
}

//InitTodoistAPI creates and returns an instance of a TodoistAPI struct.
//This instance will be used for interacting with the API.
//Takes a user auth token string as a parameter.
func InitTodoistAPI(api_token string) TodoistAPI {
	commands := []Command{}
	sync_object := SyncObject{api_token, "https://todoist.com/API/v7/sync", "*", commands}
	projects := Projects{&sync_object}
	items := Items{&sync_object}
	todoist_api := TodoistAPI{projects, items, &sync_object}
	return todoist_api
}

//apiResponseToMap converts the response from TodoistAPI to a map containing the keys
//body(body of response) and status(http status of response)
func apiResponseToMap(response *http.Response) map[string]interface{} {
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
	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
