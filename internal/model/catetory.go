package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CategoryCreateUpdateBase 创建/修改商品分类基类
type CategoryCreateUpdateBase struct {
	ParentId uint
	PicUrl   string
	Name     string
	Sort     uint8
	Level    uint8
}

// CategoryCreateInput 创建商品分类
type CategoryCreateInput struct {
	CategoryCreateUpdateBase
}

// CategoryCreateOutput 创建商品分类返回结果
type CategoryCreateOutput struct {
	CategoryId int `json:"category_id"`
}
type CategoryDeleteReq struct {
	g.Meta `path:"/category/delete" method:"delete" tags:"商品分类" summary:"删除商品分类接口"`
	Id     uint `v:"min:1#请选择需要删除的商品分类" dc:"商品分类id"`
}
type CategoryDeleteRes struct{}

// CategoryUpdateInput 修改商品分类
type CategoryUpdateInput struct {
	CategoryCreateUpdateBase
	Id uint
}

// CategoryGetListInput 获取内容列表
type CategoryGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)

}

// CategoryGetListOutput 查询列表结果
type CategoryGetListOutput struct {
	List  []CategoryGetListOutputItem `json:"list" description:"列表"`
	Page  int                         `json:"page" description:"分页码"`
	Size  int                         `json:"size" description:"分页数量"`
	Total int64                       `json:"total" description:"数据总数"`
}

// CategorySearchInput 搜索列表
type CategorySearchInput struct {
	Key        string // 关键字
	CategoryId uint   // 栏目ID
	Page       int    // 分页号码
	Size       int    // 分页数量，最大50
	Sort       int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// CategorySearchOutput 搜索列表结果
type CategorySearchOutput struct {
	List  []CategorySearchOutputItem `json:"list"`  // 列表
	Stats map[string]int             `json:"stats"` // 搜索统计
	Page  int                        `json:"page"`  // 分页码
	Size  int                        `json:"size"`  // 分页数量
	Total int                        `json:"total"` // 数据总数
}

type CategoryGetListOutputItem struct {
	//Category *CategoryListItem `json:"category"`
	Id        uint        `json:"id"` // 自增ID
	ParentId  uint        `json:"parent_id"`
	PicUrl    string      `json:"pic_url"`
	Name      string      `json:"link"`
	Sort      uint        `json:"sort"`       // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
	Level     uint        `json:"level"`      //级别
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}

type CategorySearchOutputItem struct {
	CategoryGetListOutputItem
}

//// CategoryListItem 主要用于列表展示
//type CategoryListItem struct {
//	Id        uint        `json:"id"` // 自增ID
//	PicUrl    string      `json:"pic_url"`
//	Link      string      `json:"link"`
//	Sort      uint        `json:"sort"`       // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
//	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
//	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
//}
