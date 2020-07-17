// Package functions defines sensor data model
// and implements CRUD operations of smart farm.
package functions

import (
	"errors"
	"fmt"
	"time"
)

// [Start cloud_functions_sensor_data_struct]

// SensorData represents a single document of smart farm sensor data collection.
type SensorData struct {
	UUID                   string    `json:"uuid"`
	LiquidTemperature      float64   `json:"liquid_temperature"`
	Temperature            float64   `json:"temperature"`
	Humidity               float64   `json:"humidity"`
	LiquidFlowRate         float64   `json:"liquid_flow_rate"`
	PH                     float64   `json:"ph"`
	ElectricalConductivity float64   `json:"ec"`
	Light                  float64   `json:"light"`
	LiquidLevel            bool      `json:"liquid_level"`
	Valve                  bool      `json:"valve"`
	LED                    bool      `json:"led"`
	Fan                    bool      `json:"fan"`
	UnixTime               int64     `json:"unix_time"`
	LocalTime              time.Time `json:"local_time"`
}

// [End cloud_functions_sensor_data_struct]

// setTime sets the time of sensor data.
func (s *SensorData) setTime() {
	s.LocalTime = time.Now()
	s.UnixTime = s.LocalTime.Unix()
}

// verify verifies that there are any invalid values in sensor data.
func (s SensorData) verify() error {
	var msg string
	if s.PH < 0 || s.PH > 14 {
		msg += fmt.Sprintf("Invalid value in pH: %f\n", s.PH)
	}
	if s.ElectricalConductivity < 0 || s.ElectricalConductivity > 2 {
		msg += fmt.Sprintf("Invalid value in electrical conductivity: %f\n", s.ElectricalConductivity)
	}
	if s.Light < 0 || s.Light > 100 {
		msg += fmt.Sprintf("Invalid value in light intensity: %f\n", s.Light)
	}
	if msg != "" {
		return errors.New(msg)
	}
	return nil
}
