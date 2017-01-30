package gotodoist

type Projects struct {
	sync_object *SyncObject
}

func (projects Projects) Add(names []string) (map[string]interface{}, error) {
	commands := []Command{}
	name_temp_ids := map[string]string{}
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
		name_temp_ids[name] = temp_id
	}
	response, err := projects.sync_object.callWriteApi(commands)
	response_map := apiResponseToMap(response)
	body := response_map["body"].(map[string]interface{})
	id_mapping := body["temp_id_mapping"].(map[string]interface{})
	name_ids := map[string]int64{}
	for k, v := range name_temp_ids {
		name_ids[k] = int64(id_mapping[v].(float64))
	}
	response_map["name_ids"] = name_ids
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

func (projects Projects) Share(email string, project_id int64) (map[string]interface{}, error) {
	commands := []Command{}
	uuid, _ := newUUID()
	temp_id, _ := newUUID()
	commands = append(commands, Command{
		Type:   "share_project",
		UUID:   uuid,
		TempID: temp_id,
		Args: map[string]interface{}{
			"email":      email,
			"project_id": project_id,
		},
	})
	response, err := projects.sync_object.callWriteApi(commands)
	response_map := apiResponseToMap(response)
	defer response.Body.Close()
	return response_map, err
}

func (projects Projects) QueueShare(email string, project_id int64) {
	commands := []Command{}
	uuid, _ := newUUID()
	temp_id, _ := newUUID()
	commands = append(commands, Command{
		Type:   "share_project",
		UUID:   uuid,
		TempID: temp_id,
		Args: map[string]interface{}{
			"email":      email,
			"project_id": project_id,
		},
	})
	projects.sync_object.queueCommands(commands)
}

func (projects Projects) Unshare(email string, project_id int64) (map[string]interface{}, error) {
	commands := []Command{}
	uuid, _ := newUUID()
	temp_id, _ := newUUID()
	commands = append(commands, Command{
		Type:   "delete_collaborator",
		UUID:   uuid,
		TempID: temp_id,
		Args: map[string]interface{}{
			"email":      email,
			"project_id": project_id,
		},
	})
	response, err := projects.sync_object.callWriteApi(commands)
	response_map := apiResponseToMap(response)
	defer response.Body.Close()
	return response_map, err
}

func (projects Projects) QueueUnshare(email string, project_id int64) {
	commands := []Command{}
	uuid, _ := newUUID()
	temp_id, _ := newUUID()
	commands = append(commands, Command{
		Type:   "delete_collaborator",
		UUID:   uuid,
		TempID: temp_id,
		Args: map[string]interface{}{
			"email":      email,
			"project_id": project_id,
		},
	})
	projects.sync_object.queueCommands(commands)
}

func (projects Projects) AddNote(content string, project_id int64) (map[string]interface{}, error) {
	commands := []Command{}
	uuid, _ := newUUID()
	temp_id, _ := newUUID()
	commands = append(commands, Command{
		Type:   "note_add",
		UUID:   uuid,
		TempID: temp_id,
		Args: map[string]interface{}{
			"content":    content,
			"project_id": project_id,
		},
	})
	response, err := projects.sync_object.callWriteApi(commands)
	response_map := apiResponseToMap(response)
	defer response.Body.Close()
	return response_map, err
}

func (projects Projects) QueueAddNote(content string, project_id int64) {
	commands := []Command{}
	uuid, _ := newUUID()
	temp_id, _ := newUUID()
	commands = append(commands, Command{
		Type:   "note_add",
		UUID:   uuid,
		TempID: temp_id,
		Args: map[string]interface{}{
			"content":    content,
			"project_id": project_id,
		},
	})
	projects.sync_object.queueCommands(commands)
}

func (projects Projects) GetAll() (map[string]interface{}, error) {
	resource_types := []string{"projects"}
	response, err := projects.sync_object.callReadApi(resource_types)
	response_map := apiResponseToMap(response)
	response.Body.Close()
	return response_map, err
}
