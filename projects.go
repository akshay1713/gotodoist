package gotodoist

//Projects is a struct on which all the project related api calls are made
type Projects struct {
	sync_object *SyncObject
}

//Add adds new projects.
//Takes a slice of names as parameter.
//Returns a map with the following keys-
//body[body of the response], status[http status of the response],
//name_ids[a map of the project names to the project ids]
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

//QueueAdd queues project add command for the given names.
//Project names queued will be added when the Commit function is called.
//Takes a slice of names as parameter
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

//Delete deletes existing projects.
//Takes a slice of project ids as parameter.
//Returns a map with the following keys-
//body[body of the response], status[http status of the response],
func (projects Projects) Delete(project_ids []int64) (map[string]interface{}, error) {
	uuid, _ := newUUID()
	commands := []Command{
		{
			Type: "project_delete",
			UUID: uuid,
			Args: map[string]interface{}{
				"ids": project_ids,
			},
		},
	}
	response, err := projects.sync_object.callWriteApi(commands)
	response_map := apiResponseToMap(response)
	defer response.Body.Close()
	return response_map, err
}

//QueueDelete queues delete command for the given project id.
//Projects will be deleted when the Commit function is called.
//Takes a slice of project ids as parameter.
func (projects Projects) QueueDelete(project_ids []int64) {
	uuid, _ := newUUID()
	commands := []Command{
		{
			Type: "project_delete",
			UUID: uuid,
			Args: map[string]interface{}{
				"ids": project_ids,
			},
		},
	}
	projects.sync_object.queueCommands(commands)
}

//Share shares the given project with the given email (Adds a collaborator to the project).
//Takes an email and a project id (int64) as parameter.
//Returns a map with the following keys-
//body[body of the response], status[http status of the response],
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

//QueueShare queues share project command for the given project id and email.
//Project will be shared when the Commit function is called.
//Takes an email and a project id (int64) as parameter.
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

//Unshare unshared a project with an email (Removes a collaborator from a project).
//Takes an email and a project id (int64) as parameter.
//Returns a map with the following keys-
//body[body of the response], status[http status of the response],
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

//QueueUnshare queues a project unshare command with the given email and project id.
//Project will be unshared when the Commit function is called.
//Takes an email and project_id (int64) as parameter.
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

//AddNote adds a project note.
//Takes a string of note content and project id as parameters.
//Returns a map with the following keys-
//body[body of the response], status[http status of the response],
//note_id[id of the newly created note]
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
	body := response_map["body"].(map[string]interface{})
	id_mapping := body["temp_id_mapping"].(map[string]interface{})
	response_map["note_id"] = id_mapping[temp_id]
	defer response.Body.Close()
	return response_map, err
}

//QueueAddNote queues an add note command.
//Takes a string of note content and project id as parameter.
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

//GetAll returns all the existing projects.
//Returns a map with the following keys-
//body[body of the response], status[http status of the response],
func (projects Projects) GetAll() (map[string]interface{}, error) {
	resource_types := []string{"projects"}
	response, err := projects.sync_object.callReadApi(resource_types)
	response_map := apiResponseToMap(response)
	response.Body.Close()
	return response_map, err
}
