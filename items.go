package gotodoist

type Items struct {
	sync_object *SyncObject
}

func (items Items) Add(names []string, project_id int64) (map[string]interface{}, error) {
	commands := []Command{}
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
	}
	response, err := items.sync_object.callWriteApi(commands)
	response_map := apiResponseToMap(response)
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

func (items Items) GetAll() (map[string]interface{}, error) {
	resource_types := []string{"items"}
	response, err := items.sync_object.callReadApi(resource_types)
	response_map := apiResponseToMap(response)
	defer response.Body.Close()
	return response_map, err
}
