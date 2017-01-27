package gotodoist

type Projects struct {
	sync_object *SyncObject
}

func (projects Projects) Add(names []string) (map[string]interface{}, error) {
	commands := []Command{}
	for _, name := range names {
		uuid, _ := newUUID()
		temp_id, _ := newUUID()
		commands = append(commands, Command{
			Type:   "project_add",
			UUID:   uuid,
			TempID: temp_id,
			Args: map[string]interface{}{
				"name": name,
			},
		})
	}
	response, err := projects.sync_object.callWriteApi(commands)
	response_map := apiResponseToMap(response)
	defer response.Body.Close()
	return response_map, err
}

func (projects *Projects) QueueAdd(names []string) {
	commands := []Command{}
	for _, name := range names {
		uuid, _ := newUUID()
		temp_id, _ := newUUID()
		commands = append(commands, Command{
			Type:   "project_add",
			UUID:   uuid,
			TempID: temp_id,
			Args: map[string]interface{}{
				"name": name,
			},
		})
	}
	projects.sync_object.queueCommands(commands)
}

func (projects Projects) Delete(project_ids []int64) (map[string]interface{}, error) {
	uuid, _ := newUUID()
	commands := []Command{
		{
			Type: "project_delete",
			UUID: uuid,
			Args: map[string]interface{}{
				"content": project_ids,
				"ids":     project_ids,
			},
		},
	}
	response, err := projects.sync_object.callWriteApi(commands)
	response_map := apiResponseToMap(response)
	defer response.Body.Close()
	return response_map, err
}

func (projects Projects) QueueDelete(project_ids []int64) {
	uuid, _ := newUUID()
	commands := []Command{
		{
			Type: "project_delete",
			UUID: uuid,
			Args: map[string]interface{}{
				"content": project_ids,
				"ids":     project_ids,
			},
		},
	}
	projects.sync_object.queueCommands(commands)
}

func (projects Projects) GetAll() (map[string]interface{}, error) {
	resource_types := []string{"projects"}
	response, err := projects.sync_object.callReadApi(resource_types)
	response_map := apiResponseToMap(response)
	response.Body.Close()
	return response_map, err
}
