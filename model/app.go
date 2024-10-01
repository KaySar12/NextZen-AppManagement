package model

import (
	"time"
)

type ServerAppListCollection struct {
	List      []ServerAppList `json:"list"`
	Recommend []ServerAppList `json:"recommend"`
	Community []ServerAppList `json:"community"`
}

type StateEnum int

const (
	StateEnumNotInstalled StateEnum = iota
	StateEnumInstalled
)

// @tiger - 对于用于出参的数据结构，静态信息（例如 title）和
//
// Dynamic information (such as state, query_count) should be divided into different data structures
//
// This advantage is
// 1 -When obtaining dynamic information multiple times, it can reduce the complexity of the reference, because the static information is only obtained once
// 2 -In the future iterative iteration, the maintenance cost can be reduced (all fields are expanded at a level of maintenance costs slightly higher)
//
// In addition, some targeted fields, such as docker -related, can be saved with MAP.
// In the future, add polymorphic apps in the future, such as SNAP, no need to maintain multiple structures, or a unnecessary field preserved in structure
type ServerAppList struct {
	ID             uint      `gorm:"column:id;primary_key" json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Tagline        string    `json:"tagline"`
	Tags           Strings   `gorm:"type:json" json:"tags"`
	Icon           string    `json:"icon"`
	ScreenshotLink Strings   `gorm:"type:json" json:"screenshot_link"`
	Category       string    `json:"category"`
	CategoryID     int       `json:"category_id"`
	CategoryFont   string    `json:"category_font"`
	PortMap        string    `json:"port_map"`
	ImageVersion   string    `json:"image_version"`
	Tip            string    `json:"tip"`
	Envs           EnvArray  `json:"envs"`
	Ports          PortArray `json:"ports"`
	Volumes        PathArray `json:"volumes"`
	Devices        PathArray `json:"devices"`
	NetworkModel   string    `json:"network_model"`
	Image          string    `json:"image"`
	Index          string    `json:"index"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	State          StateEnum `json:"state"`
	Author         string    `json:"author"`
	MinMemory      int       `json:"min_memory"`
	MinDisk        int       `json:"min_disk"`
	Thumbnail      string    `json:"thumbnail"`
	Healthy        string    `json:"healthy"`
	Plugins        Strings   `json:"plugins"`
	Origin         string    `json:"origin"`
	Type           int       `json:"type"`
	QueryCount     int       `json:"query_count"`
	Developer      string    `json:"developer"`
	HostName       string    `json:"host_name"`
	Privileged     bool      `json:"privileged"`
	CapAdd         Strings   `json:"cap_add"`
	Cmd            Strings   `json:"cmd"`
	Architectures  Strings   `json:"architectures"`
	LatestDigest   Strings   `json:"latest_digests"`
}

type MyAppList struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Icon           string `json:"icon"`
	State          string `json:"state"`
	CustomID       string `gorm:"column:custom_id;primary_key" json:"custom_id"`
	Index          string `json:"index"`
	Port           string `json:"port"`
	Slogan         string `json:"slogan"`
	Type           string `json:"type"`
	Image          string `json:"image"`
	Volumes        string `json:"volumes"`
	Latest         bool   `json:"latest"`
	Host           string `json:"host"`
	Protocol       string `json:"protocol"`
	Created        int64  `json:"created"`
	AppStoreID     uint   `json:"appstore_id"`
	IsUncontrolled bool   `json:"is_uncontrolled"`
}

type Ports struct {
	ContainerPort uint   `json:"container_port"`
	CommendPort   int    `json:"commend_port"`
	Desc          string `json:"desc"`
	Type          int    `json:"type"` //  1:必选 2:可选 3:默认值不必显示 4:系统处理  5:container内容也可编辑
}

type Volume struct {
	ContainerPath string `json:"container_path"`
	Path          string `json:"path"`
	Desc          string `json:"desc"`
	Type          int    `json:"type"` //  1:必选 2:可选 3:默认值不必显示 4:系统处理   5:container内容也可编辑
}

type Envs struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Desc  string `json:"desc"`
	Type  int    `json:"type"` //  1:必选 2:可选 3:默认值不必显示 4:系统处理 5:container内容也可编辑
}

type Devices struct {
	ContainerPath string `json:"container_path"`
	Path          string `json:"path"`
	Desc          string `json:"desc"`
	Type          int    `json:"type"` //  1:必选 2:可选 3:默认值不必显示 4:系统处理 5:container内容也可编辑
}

type Strings []string

type MapStrings []map[string]string
