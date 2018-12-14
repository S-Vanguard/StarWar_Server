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

// type of planets
type Planets struct {

	// 记录被创建的时间 
	Created time.Time `json:"created,omitempty"`

	// 记录被编辑的时间 
	Edited time.Time `json:"edited,omitempty"`

	// 当前资源的URL 
	Url string `json:"url,omitempty"`

	// the name of planet
	Name string `json:"name,omitempty"`

	// the rotation period of the planet
	RotationPeriod string `json:"rotation_period,omitempty"`

	// the orbital period of the planet
	OrbitalPeriod string `json:"orbital_period,omitempty"`

	// the diameter of the planet
	Diameter string `json:"diameter,omitempty"`

	// the climate of the planet
	Climate string `json:"climate,omitempty"`

	// the gravity of the planet
	Gravity string `json:"gravity,omitempty"`

	// the terrain of the planet
	Terrain string `json:"terrain,omitempty"`

	// the surface water of the planet
	SurfaceWater string `json:"surface_water,omitempty"`

	// the population of the planet
	Population string `json:"population,omitempty"`

	Films []string `json:"films,omitempty"`
}
