// Package functions defines sensor data model
// and implements CRUD operations of smart farm.
package functions

// [Start cloud_functions_dependencies]
import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// [End cloud_functions_dependencies]

// [Start cloud_functions_get]

// Get brings the sensor data records for the last week from Firestore with given uuid.
// exported to https://asia-northeast1-superfarmers.cloudfunctions.net/Get
func Get(writer http.ResponseWriter, req *http.Request) {
	ctx := context.Background()

	// create new firestore client and close it when query is done
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		fmt.Fprintf(writer, "firestore.NewClient: %v", err)
		return
	}
	defer client.Close()

	// parse uuid from request body
	var request struct {
		UUID string `json:"uuid"`
	}
	if err = json.NewDecoder(req.Body).Decode(&request); err != nil {
		fmt.Fprintf(writer, "json.Decode: %v", err)
		return
	}

	// create response struct
	var resp struct {
		records []SensorData
	}

	now := time.Now().Unix()
	const weekTime = 7 * 24 * 60 * 60

	// create temporal pointer to store value of document
	sensorData := new(SensorData)

	// iterate records for the last week of the device, add to response
	cursor := client.Collection("sensor_data").Where("UUID", "==", request.UUID).Where("UnixTime", ">=", now-weekTime).Documents(ctx)
	for {
		doc, err := cursor.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Fprintf(writer, "firestore.Documents.Next: %v", err)
		}
		if err = doc.DataTo(sensorData); err != nil {
			fmt.Fprintf(writer, "firestore.DataTo: %v", err)
		}
		resp.records = append(resp.records, *sensorData)
	}

	// free temporal pointer
	sensorData = nil

	// notify that it's JSON response, send response with encoded records.
	writer.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(writer).Encode(resp); err != nil {
		fmt.Fprintf(writer, "json.Encode: %v", err)
	}
}

// [End cloud_functions_get]
