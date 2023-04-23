package dto

import (
	"encoding/json"
	"fmt"
)

func mapToStruct(src map[string]interface{}, dst interface{}) error {
	tmp, err := json.Marshal(src)
	if err != nil {
		return fmt.Errorf("json marshal src error: %w", err)
	}

	if err := json.Unmarshal(tmp, dst); err != nil {
		return fmt.Errorf("json unmarshal dst error: %w", err)
	}

	return nil
}
