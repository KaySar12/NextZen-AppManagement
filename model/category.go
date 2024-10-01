/*@Author: link a624669980@163.com
 *@Date: 2022-05-16 17:37:08
 *@LastEditors: LinkLeong
 *@LastEditTime: 2022-07-13 10:46:38
 *@FilePath: /CasaOS/model/category.go
 *@Description:
 */
package model

type ServerCategoryList struct {
	Item []Category `json:"item"`
}
type Category struct {
	ID uint `gorm:"column:id;primary_key" json:"id"`
	//CreatedAt time.Time `json:"created_at"`
	//
	//UpdatedAt time.Time `json:"updated_at"`
	Font  string `json:"font"` // @tiger - If this is related to the front end, it should not belong to the scope of the back end, but the front end is defined
	Name  string `json:"name"`
	Count uint   `json:"count"` // @tiger - count belongs to dynamic information, and should be placed in one exit (reason for another annotation about static/dynamic exit)
}
