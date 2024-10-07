package main

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "bytes"
    // "strconv"
    //"time"
)

func main() {

    // accessKey := "your_access_key"
    // secretKey := "your_secret_key"
        accessKey := "KtPoJRfVIW7dfgdUlv_c2YHy9okYqxmDMXpBbB2cFns"
        secretKey := "__OZxDPKHkRa3HPJK3GxGOvm6BAn_n_prymYUGNjjfg68WuKnAPZxuC9dWNzMHoarluSj_XMXWtRFDp14bLmyQ"

    // GET
    // url := "http://10.159.62.166:8009/api/v1/microservice/get-crop/"

    // POST
    // url := "http://10.159.62.166:8009/api/v3/microservice/create-and-send-alert/"

    // requestData := map[string]interface{}{
    //     "mode": "fcm",
    //     "plotalert": []map[string]interface{}{
    //         {
    //             "alert_type":    "irrigation",
    //             "plot_id":      "6ee5a2cb-d7b3-457c-8ea9-795757b7575f",
    //             "severity":      "High",
    //             "date_of_alert": "2024-10-01",
    //             "creation_mode": "Auto",
    //             "meta": map[string]interface{}{
    //                 "irrigation_time": 11,
    //                 "soil_moisture":  2.2,
    //             },
    //         },
    //     },
    // }


    // PUT
    url := "http://10.159.62.166:8009/api/v1/microservice/update-latest-sensor-data/"
    requestData := map[string]interface{}{
        "device_imei":          "1234",
        "device_health":        "Offline",
        "device_health_reason": "SI_DEVICE_DATA_LESS_THAN_THREE",
        "module":               "dataengine",
        "data": []map[string]interface{}{
            {
                "sensor":        "tm",
                "sensor_health": "Offline",
                "error_msg":     "SI_SENSOR_DATA_UNAVAIL",
            },
            {
                "sensor":        "hm",
                "sensor_health": "Offline",
                "error_msg":     "SI_SENSOR_DATA_UNAVAIL",
            },
            {
                "sensor":        "pp",
                "sensor_health": "Offline",
                "error_msg":     "SI_SENSOR_DATA_UNAVAIL",
            },
            {
                "sensor":        "sm",
                "sensor_health": "Offline",
                "error_msg":     "SI_SENSOR_DATA_UNAVAIL",
            },
            {
                "sensor":        "st",
                "sensor_health": "Offline",
                "error_msg":     "SI_SENSOR_DATA_UNAVAIL",
            },
            {
                "sensor":        "sc",
                "sensor_health": "Offline",
                "error_msg":     "SI_SENSOR_DATA_UNAVAIL",
            },
            {
                "sensor":        "lt",
                "sensor_health": "Offline",
                "error_msg":     "SI_SENSOR_DATA_UNAVAIL",
            },
            {
                "sensor":        "lw",
                "sensor_health": "Offline",
                "error_msg":     "SI_SENSOR_DATA_UNAVAIL",
            },
            {
                "sensor":        "bl",
                "sensor_health": "Offline",
                "error_msg":     "SI_SENSOR_DATA_UNAVAIL",
            },
            {
                "sensor":        "pv",
                "sensor_health": "Offline",
                "error_msg":     "SI_SENSOR_DATA_UNAVAIL",
            },
        },
    }

    // For POST/PUT
    // requestData = map[string]interface{}{"key1": "value1", "key2": "value2"}
    // url = url
    // sendRequest("POST", url, accessKey, secretKey, requestData)
    // sendRequest("GET", url, accessKey, secretKey, nil)
    sendRequest("PUT", url, accessKey, secretKey, requestData)
    // sendRequest("DELETE", url, accessKey, secretKey, nil)
}

func sendRequest(method, url, accessKey, secretKey string, requestData interface{}) {

    //timestamp := time.Now().Unix()

    var jsonData []byte
    var err error

    if requestData != nil {
        jsonData, err = json.Marshal(requestData)
        if err != nil {
            fmt.Println("Error marshaling request JSON:", err)
            return
        }
    }

    // Create the full URL with the timestamp as a query parameter
    //fullURL := fmt.Sprintf("%s?timestamp=%d", url, timestamp)
    fullURL := url + string(jsonData)
    fmt.Println("\nFullUrl=", fullURL)

    // Create HMAC SHA256 signature
    mac := hmac.New(sha256.New, []byte(secretKey))
    mac.Write([]byte(fullURL))
    // if requestData != nil {
    //     mac.Write(jsonData)
    // }
    signature := fmt.Sprintf("%x", mac.Sum(nil))
        fmt.Println("\nSignature=", signature)

    client := &http.Client{}

    // Create a new HTTP request
    req, err := http.NewRequest(method, url, bytes.NewReader(jsonData))
    if err != nil {
        fmt.Println("Error creating request:", err)
        return
    }

    // Set headers
    req.Header.Set("Access-Key", accessKey)
    req.Header.Set("Signature", signature)
    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error making request:", err)
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return
    }

    fmt.Println("Response Body:", string(body))
}

