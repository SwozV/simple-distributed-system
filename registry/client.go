package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// 为 register-service 发送 post 请求
func RegisterService(r Registration) error {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err := enc.Encode(r)
	if err != nil {
		return err
	}

	res, err := http.Post(ServiceURL, "application/json", buf)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to register service. Register service"+
			"responed with code %v", res.StatusCode)
	}

	return nil
}
