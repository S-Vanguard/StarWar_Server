/*
 * star world API
 *
 * the api to query the information about *Star War* you can check all the Star Wars data you've ever wanted Planets Spaceships Vehicles People Films and Species From all SEVEN Star Wars films
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package apiType

type UserSignIn struct {

	// 尝试登陆的用户名
	Username string `json:"username"`

	// 尝试登陆的用户密码
	Password string `json:"password"`
}
