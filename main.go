package encryption

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "bytes"
    "time"
)

// SendRequest sends an HTTP request (GET, POST, PUT, DELETE) to the specified URL.
func SendRequest(method, url, accessKey, secretKey string, requestData interface{}) (string, error) {
    var jsonData []byte
    var err error

    if requestData != nil {
        jsonData, err = json.Marshal(requestData)
        if err != nil {
            return "", fmt.Errorf("error marshaling request JSON: %w", err)
        }
    }

    // Create HMAC SHA256 signature
    mac := hmac.New(sha256.New, []byte(secretKey))
    mac.Write([]byte(url))
    if requestData != nil {
        mac.Write(jsonData)
    }
    signature := fmt.Sprintf("%x", mac.Sum(nil))

    client := &http.Client{
	Timeout: 10 * time.Second,
    }

    // Create a new HTTP request
    req, err := http.NewRequest(method, url, bytes.NewReader(jsonData))
    if err != nil {
        return "", fmt.Errorf("error creating request: %w", err)
    }

    // Set headers
    req.Header.Set("Access-Key", accessKey)
    req.Header.Set("Signature", signature)
    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        return "", fmt.Errorf("error making request: %w", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", fmt.Errorf("error reading response body: %w", err)
    }

    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("error response: %s", string(body))
    }

    return string(body), nil
}


