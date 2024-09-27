package service

import (
	"fiber-api/app/model"
	"fiber-api/pkg/db"
)

// 菜单数据服务
type Menu struct {
	Id        int    `json:"id"`
	ParentId  int    `json:"parentId"`
	Name      string `json:"name"`
	Type      int    `json:"type"`
	Sort      int    `json:"sort"`
	Path      string `json:"path"`
	Component string `json:"component"`
	Icon      string `json:"icon"`
	Redirect  string `json:"redirect"`
	Status    int    `json:"status"`
}

type MenuTree struct {
	*Menu
	Children []*MenuTree `json:"children"`
}

// 创建菜单
func (*Menu) Create(menu *Menu) error {
	return db.GormClient.Model(&model.Menu{}).Create(&model.Menu{
		ParentId:  menu.ParentId,
		Name:      menu.Name,
		Type:      menu.Type,
		Sort:      menu.Sort,
		Path:      menu.Path,
		Component: menu.Component,
		Icon:      menu.Icon,
		Redirect:  menu.Redirect,
		Status:    menu.Status,
	}).Error
}

// 更新菜单
func (*Menu) Update(menu *Menu) error {
	return db.GormClient.Model(&model.Menu{}).Where("id = ?", menu.Id).Updates(&model.Menu{
		ParentId:  menu.ParentId,
		Name:      menu.Name,
		Type:      menu.Type,
		Sort:      menu.Sort,
		Path:      menu.Path,
		Component: menu.Component,
		Icon:      menu.Icon,
		Redirect:  menu.Redirect,
		Status:    menu.Status,
	}).Error
}

// 删除菜单
func (*Menu) Delete(id int) error {
	return db.GormClient.Model(&model.Menu{}).Where("id = ?", id).Delete(nil).Error
}

// 获取下级菜单
func (t *Menu) ChildrenList(id, status int) []*MenuTree {

	var (
		list []*Menu
		tree []*MenuTree
	)

	db.GormClient.Model(&model.Menu{}).Order("id desc").Where("parent_id = ?", id).Find(&list)
	if len(list) == 0 {
		return nil
	}

	for _, item := range list {
		child := &MenuTree{
			Menu: item,
		}
		// 递归获取子菜单树
		child.Children = t.ChildrenList(item.Id, status)
		tree = append(tree, child)
	}

	return tree
}

// 获取菜单列表
func (*Menu) ListByIds(ids []int) []*Menu {

	var list []*Menu

	query := db.GormClient.Model(&model.Menu{}).Where("status = ?", 1)

	if len(ids) > 0 {
		query.Where("id in ?", ids)
	}

	query.Find(&list)

	return list
}

// 菜单详情
func (*Menu) Detail(id int) *Menu {

	var detail *Menu

	db.GormClient.Model(&model.Menu{}).Where("id = ?", id).Take(&detail)

	return detail
}

// 列表转树形
func (t *Menu) ListToTree(list []*Menu, parentId int) []*MenuTree {

	if len(list) <= 0 {
		return nil
	}

	var tree []*MenuTree

	for _, item := range list {
		if item.ParentId == parentId {
			tree = append(tree, &MenuTree{
				Menu:     item,
				Children: t.ListToTree(list, item.Id),
			})
		}
	}

	return tree
}
