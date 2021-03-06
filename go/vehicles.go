/*
 * star world API
 *
 * the api to query the information about *Star War* you can check all the Star Wars data you've ever wanted Planets Spaceships Vehicles People Films and Species From all SEVEN Star Wars films 
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

// the response type of vehicles
type Vehicles struct {

	// 记录被创建的时间 
	Created time.Time `json:"created,omitempty"`

	// 记录被编辑的时间 
	Edited time.Time `json:"edited,omitempty"`

	// 当前资源的URL 
	Url string `json:"url,omitempty"`

	Name string `json:"name,omitempty"`

	Model string `json:"model,omitempty"`

	Manufacturer string `json:"manufacturer,omitempty"`

	CostInCredits string `json:"cost_in_credits,omitempty"`

	Length string `json:"length,omitempty"`

	MaxAtmospheringSpeed string `json:"max_atmosphering_speed,omitempty"`

	Crew string `json:"crew,omitempty"`

	Passengers string `json:"passengers,omitempty"`

	CargoCapacity string `json:"cargo_capacity,omitempty"`

	Consumables string `json:"consumables,omitempty"`

	VehicleClass string `json:"vehicle_class,omitempty"`

	Pilots []string `json:"pilots,omitempty"`

	Films []string `json:"films,omitempty"`
}
