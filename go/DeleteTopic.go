package main

import "fmt"

func DeleteTopic(top Data_Topic) (int64, error) {
	result, err := db.Exec("DELETE FROM data_topic WHERE ID =  (?)", top.ID)
	fmt.Println(result)
	if err != nil {
		return 0, fmt.Errorf("DeleteTopic: %v", err)
	}
	id, err := result.LastInsertId()
	fmt.Println(id)
	if err != nil {
		return 0, fmt.Errorf("DeleteTopic: %v", err)
	}
	return id, nil
}
