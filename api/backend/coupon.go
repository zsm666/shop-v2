package backend

import "github.com/gogf/gf/v2/frame/g"

type CouponReq struct {
	g.Meta `path:"/coupon/add" tags:"Coupon" method:"post" summary:"添加优惠券"`
	CouponCommonAddUpdate
}

type CouponCommonAddUpdate struct {
	Name       string `json:"name" v:"required#名称不能为空"`
	Price      int    `json:"price" v:"require#优惠券必填" dc:"优惠券金额"`
	GoodsIds   string `json:"goods_ids"  dc:"可用商品ID，逗号分割"`
	CategoryId uint   `json:"category_id"  dc:"可用商品分类"`
}

type CouponRes struct {
	CouponId int `json:"coupon_id"`
}
type CouponDeleteReq struct {
	g.Meta `path:"/coupon/delete" method:"delete" tags:"商品分类" summary:"删除商品分类接口"`
	Id     uint `v:"min:1#请选择需要删除的商品分类" dc:"商品分类id"`
}
type CouponDeleteRes struct{}

type CouponUpdateReq struct {
	g.Meta `path:"/coupon/update/" method:"post" tags:"商品分类" summary:"商品分类接口"`
	Id     uint `json:"id"      v:"min:1#请选择需要修改的商品分类" dc:"商品分类Id"`
	CouponCommonAddUpdate
}
type CouponUpdateRes struct {
	Id uint `json:"id"`
}

type CouponGetListCommonReq struct {
	g.Meta `path:"/coupon/list" method:"get" tags:"商品分类" summary:"商品分类列表"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}
type CouponGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  uint        `json:"page" description:"分页码"`
	Size  uint        `json:"size" description:"分页数量"`
	Total int64       `json:"total" description:"数据总数"`
}
type CouponGetListAllCommonReq struct {
	g.Meta `path:"/coupon/list/all" method:"get" tags:"商品分类" summary:"商品分类全部列表"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
}
type CouponGetListAllCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Total int64       `json:"total" description:"数据总数"`
}
