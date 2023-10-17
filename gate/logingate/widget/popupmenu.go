// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package widget

import "github.com/ying32/govcl/vcl"

type TActiveListPopupMenu struct {
	*vcl.TPopupMenu

	popMenuRefList     *vcl.TMenuItem
	popMenuSort        *vcl.TMenuItem
	popMenuAddTempList *vcl.TMenuItem
	popMenuBlockList   *vcl.TMenuItem
	popMenuKick        *vcl.TMenuItem
}

type TBlockListPopupMenu struct {
	*vcl.TPopupMenu

	popMenuRefList     *vcl.TMenuItem
	popMenuSort        *vcl.TMenuItem
	popMenuAdd         *vcl.TMenuItem
	popMenuAddTempList *vcl.TMenuItem
	popMenuDelete      *vcl.TMenuItem
}

type TTempBlockListPopupMenu struct {
	*vcl.TPopupMenu

	popMenuRefList   *vcl.TMenuItem
	popMenuSort      *vcl.TMenuItem
	popMenuAdd       *vcl.TMenuItem
	popMenuBlockList *vcl.TMenuItem
	popMenuDelete    *vcl.TMenuItem
}

func NewActiveListPopupMenu(sender vcl.IComponent) *TActiveListPopupMenu {
	sf := &TActiveListPopupMenu{
		TPopupMenu: vcl.NewPopupMenu(sender),
	}

	sf.popMenuRefList = vcl.NewMenuItem(sf)
	sf.popMenuRefList.SetCaption("刷新(&R)")

	sf.popMenuSort = vcl.NewMenuItem(sf)
	sf.popMenuSort.SetCaption("排序(&S)")

	sf.popMenuAddTempList = vcl.NewMenuItem(sf)
	sf.popMenuAddTempList.SetCaption("加入动态过滤列表(&A)")

	sf.popMenuBlockList = vcl.NewMenuItem(sf)
	sf.popMenuBlockList.SetCaption("加入永久过滤列表(&D)")

	sf.popMenuKick = vcl.NewMenuItem(sf)
	sf.popMenuKick.SetCaption("踢除下线(&K)")

	sf.Items().Add(sf.popMenuRefList)
	sf.Items().Add(sf.popMenuSort)
	sf.Items().Add(sf.popMenuAddTempList)
	sf.Items().Add(sf.popMenuBlockList)
	sf.Items().Add(sf.popMenuKick)

	return sf
}

func NewBlockListPopupMenu(sender vcl.IComponent) *TBlockListPopupMenu {
	sf := &TBlockListPopupMenu{
		TPopupMenu: vcl.NewPopupMenu(sender),
	}

	sf.popMenuRefList = vcl.NewMenuItem(sf)
	sf.popMenuRefList.SetCaption("刷新(&R)")

	sf.popMenuSort = vcl.NewMenuItem(sf)
	sf.popMenuSort.SetCaption("排序(&S)")

	sf.popMenuAdd = vcl.NewMenuItem(sf)
	sf.popMenuAdd.SetCaption("增加(&A)")

	sf.popMenuAddTempList = vcl.NewMenuItem(sf)
	sf.popMenuAddTempList.SetCaption("加入动态过滤列表(&A)")

	sf.popMenuDelete = vcl.NewMenuItem(sf)
	sf.popMenuDelete.SetCaption("删除(&D)")

	sf.Items().Add(sf.popMenuRefList)
	sf.Items().Add(sf.popMenuSort)
	sf.Items().Add(sf.popMenuAdd)
	sf.Items().Add(sf.popMenuAddTempList)
	sf.Items().Add(sf.popMenuDelete)

	return sf
}

func NewTempBlockListPopupMenu(sender vcl.IComponent) *TTempBlockListPopupMenu {
	sf := &TTempBlockListPopupMenu{
		TPopupMenu: vcl.NewPopupMenu(sender),
	}

	sf.popMenuRefList = vcl.NewMenuItem(sf)
	sf.popMenuRefList.SetCaption("刷新(&R)")

	sf.popMenuSort = vcl.NewMenuItem(sf)
	sf.popMenuSort.SetCaption("排序(&S)")

	sf.popMenuAdd = vcl.NewMenuItem(sf)
	sf.popMenuAdd.SetCaption("增加(&A)")

	sf.popMenuBlockList = vcl.NewMenuItem(sf)
	sf.popMenuBlockList.SetCaption("加入永久过滤列表(&D)")

	sf.popMenuDelete = vcl.NewMenuItem(sf)
	sf.popMenuDelete.SetCaption("删除(&D)")

	sf.Items().Add(sf.popMenuRefList)
	sf.Items().Add(sf.popMenuSort)
	sf.Items().Add(sf.popMenuAdd)
	sf.Items().Add(sf.popMenuBlockList)
	sf.Items().Add(sf.popMenuDelete)

	return sf
}
