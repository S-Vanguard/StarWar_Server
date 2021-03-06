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

// the response type of films
type Films struct {

	// 记录被创建的时间
	Created time.Time `json:"created,omitempty"`

	// 记录被编辑的时间
	Edited time.Time `json:"edited,omitempty"`

	// 当前资源的URL
	Url string `json:"url,omitempty"`

	Title string `json:"title,omitempty"`

	EpisodeId int `json:"episode_id,omitempty"`

	OpeningCrawl string `json:"opening_crawl,omitempty"`

	Director string `json:"director,omitempty"`

	Producer string `json:"producer,omitempty"`

	ReleaseDate string `json:"release_date,omitempty"`

	Characters []string `json:"characters,omitempty"`

	Planets []string `json:"planets,omitempty"`

	Starships []string `json:"starships,omitempty"`

	Vehicles []string `json:"vehicles,omitempty"`

	Species []string `json:"species,omitempty"`
}
