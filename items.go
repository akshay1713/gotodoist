package gotodoist

type Items struct {
	sync_object *SyncObject
}

func (items Items) Add(names []string, project_id int64) (map[string]interface{}, error) {
	commands := []Command{}
	name_temp_ids := map[string]string{}
	for _, name := range names {
		uuid, _ := newUUID()
		temp_id, _ := newUUID()
		commands = append(commands, Command{
			Type:   "item_add",
			UUID:   uuid,
			TempID: temp_id,
			Args: map[string]interface{}{
				"content":    name,
				"project_id": project_id,
			},
		})
		name_temp_ids[name] = temp_id
	}
	response, err := items.sync_object.callWriteApi(commands)
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

func (items *Items) QueueAdd(names []string) {
	commands := []Command{}
	for _, name := range names {
		uuid, _ := newUUID()
		temp_id, _ := newUUID()
		commands = append(commands, Command{
			Type:   "item_add",
			UUID:   uuid,
			TempID: temp_id,
			Args: map[string]interface{}{
				"name": name,
			},
		})
	}
	items.sync_object.queueCommands(commands)
}

func (items Items) Delete(item_ids []int64) (map[string]interface{}, error) {
	uuid, _ := newUUID()
	commands := []Command{
		{
			Type: "item_delete",
			UUID: uuid,
			Args: map[string]interface{}{
				"content": item_ids,
				"ids":     item_ids,
			},
		},
	}
	response, err := items.sync_object.callWriteApi(commands)
	response_map := apiResponseToMap(response)
	defer response.Body.Close()
	return response_map, err
}

func (items Items) QueueDelete(item_ids []int64) {
	uuid, _ := newUUID()
	commands := []Command{
		{
			Type: "item_delete",
			UUID: uuid,
			Args: map[string]interface{}{
				"content": item_ids,
				"ids":     item_ids,
			},
		},
	}
	items.sync_object.queueCommands(commands)
}

func (items Items) Complete(item_ids []int64) (map[string]interface{}, error) {
	uuid, _ := newUUID()
	commands := []Command{
		{
			Type: "item_complete",
			UUID: uuid,
			Args: map[string]interface{}{
				"ids": item_ids,
			},
		},
	}
	response, err := items.sync_object.callWriteApi(commands)
	response_map := apiResponseToMap(response)
	defer response.Body.Close()
	return response_map, err
}

func (items Items) QueueComplete(item_ids []int64) {
	uuid, _ := newUUID()
	commands := []Command{
		{
			Type: "item_complete",
			UUID: uuid,
			Args: map[string]interface{}{
				"ids": item_ids,
			},
		},
	}
	items.sync_object.queueCommands(commands)
}

func (items Items) Uncomplete(item_ids []int64) (map[string]interface{}, error) {
	uuid, _ := newUUID()
	commands := []Command{
		{
			Type: "item_uncomplete",
			UUID: uuid,
			Args: map[string]interface{}{
				"ids": item_ids,
			},
		},
	}
	response, err := items.sync_object.callWriteApi(commands)
	response_map := apiResponseToMap(response)
	defer response.Body.Close()
	return response_map, err
}

func (items Items) QueueUncomplete(item_ids []int64) {
	uuid, _ := newUUID()
	commands := []Command{
		{
			Type: "item_uncomplete",
			UUID: uuid,
			Args: map[string]interface{}{
				"ids": item_ids,
			},
		},
	}
	items.sync_object.queueCommands(commands)
}

func (items Items) Close(item_id int64) (map[string]interface{}, error) {
	uuid, _ := newUUID()
	commands := []Command{
		{
			Type: "item_close",
			UUID: uuid,
			Args: map[string]interface{}{
				"ids": item_id,
			},
		},
	}
	response, err := items.sync_object.callWriteApi(commands)
	response_map := apiResponseToMap(response)
	defer response.Body.Close()
	return response_map, err
}

func (items Items) QueueClose(item_id int64) {
	uuid, _ := newUUID()
	commands := []Command{
		{
			Type: "item_close",
			UUID: uuid,
			Args: map[string]interface{}{
				"ids": item_id,
			},
		},
	}
	items.sync_object.queueCommands(commands)
}

func (items Items) AddNote(content string, item_id int64) (map[string]interface{}, error) {
	commands := []Command{}
	uuid, _ := newUUID()
	temp_id, _ := newUUID()
	commands = append(commands, Command{
		Type:   "note_add",
		UUID:   uuid,
		TempID: temp_id,
		Args: map[string]interface{}{
			"content": content,
			"item_id": item_id,
		},
	})
	response, err := items.sync_object.callWriteApi(commands)
	response_map := apiResponseToMap(response)
	defer response.Body.Close()
	return response_map, err
}

func (items Items) QueueAddNote(content string, item_id int64) {
	commands := []Command{}
	uuid, _ := newUUID()
	temp_id, _ := newUUID()
	commands = append(commands, Command{
		Type:   "note_add",
		UUID:   uuid,
		TempID: temp_id,
		Args: map[string]interface{}{
			"content": content,
			"item_id": item_id,
		},
	})
	items.sync_object.queueCommands(commands)
}

func (items Items) GetAll() (map[string]interface{}, error) {
	resource_types := []string{"items"}
	response, err := items.sync_object.callReadApi(resource_types)
	response_map := apiResponseToMap(response)
	defer response.Body.Close()
	return response_map, err
}
