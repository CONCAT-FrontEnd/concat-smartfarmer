// Package functions defines sensor data model
// and implements CRUD operations of smart farm.
package functions

// [Start cloud_functions_dependencies]
import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"cloud.google.com/go/firestore"
)

// [End cloud_functions_dependencies]

const projectID = "superfarmers"

// [Start cloud_functions_insert]

// Insert stores a sensor data into Firestore.
// exported to https://asia-northeast1-superfarmers.cloudfunctions.net/Insert
func Insert(writer http.ResponseWriter, req *http.Request) {
	ctx := context.Background()

	// create new firestore client and close it when query is done.
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		fmt.Fprintf(writer, "firestore.NewClient: %v", err)
		return
	}
	defer client.Close()

	// create new sensor data and parse json in request body.
	var sensorData SensorData
	if err = json.NewDecoder(req.Body).Decode(&sensorData); err != nil {
		fmt.Fprintf(writer, "json.Decode: %v", err)
		return
	}
	// update creation time of sensor data
	sensorData.setTime()

	// check whether sensor data is valid or not
	if err = sensorData.verify(); err != nil {
		fmt.Fprintf(writer, "verification failed: %v", err)
	}

	// store into collection
	if _, _, err = client.Collection("sensor_data").Add(ctx, sensorData); err != nil {
		fmt.Fprintf(writer, "firestore.Add: %v", err)
		return
	}
}

// [End cloud_functions_insert]
