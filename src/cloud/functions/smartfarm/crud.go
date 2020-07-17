// Package smartfarm defines sensor data model of smart farm
// and implements cloud functions for CRUD operations.
package smartfarm

// [Start smart_farm_dependencies]
import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"google.golang.org/api/iterator"

	"cloud.google.com/go/firestore"
)

// [End smart_farm_dependencies]

const projectID = "superfarmers"

// [Start smart_farm_insert]

// Insert inserts sensor data into Firestore.
func Insert(writer http.ResponseWriter, req *http.Request) {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		fmt.Fprintf(writer, "firestore.NewClient: %v", err)
		return
	}
	defer client.Close()

	var sensorData SensorData
	if err = json.NewDecoder(req.Body).Decode(&sensorData); err != nil {
		fmt.Fprintf(writer, "json.Decode: %v", err)
		return
	}
	// update creation time of sensor data
	sensorData.setTime()

	// check whether sensor data is usual or not
	if err = sensorData.verify(); err != nil {
		fmt.Fprintf(writer, "verification failed: %v", err)
	}

	// store into collection
	if _, _, err = client.Collection("sensor_data").Add(ctx, sensorData); err != nil {
		fmt.Fprintf(writer, "firestore.Add: %v", err)
		return
	}
}

// [End smart_farm_insert]

// [Start smart_farm_get]

// Get gets documents from Firestore with given query.
func Get(writer http.ResponseWriter, req *http.Request) {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		fmt.Fprintf(writer, "firestore.NewClient: %v", err)
		return
	}
	defer client.Close()

	var request struct {
		UUID string `json:"uuid"`
	}
	if err = json.NewDecoder(req.Body).Decode(&request); err != nil {
		fmt.Fprintf(writer, "json.Decode: %v", err)
		return
	}

	now := time.Now().Unix()

	cursor := client.Collection("sensor_data").Where("uuid", "==", request.UUID).Where("unix_time", ">=", now-7*24*60*60).Documents(ctx)
	for {
		doc, err := cursor.Next()
		if err == iterator.Done {
			break
		}
		// TODO: send json data with documents
		_ = doc
	}
}

// [End smart_farm_get]
