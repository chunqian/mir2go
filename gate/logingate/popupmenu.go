// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package main

import "github.com/ying32/govcl/vcl"

type TActiveListPopupMenu struct {
	*vcl.TPopupMenu

	PopMenuRefList     *vcl.TMenuItem
	PopMenuSort        *vcl.TMenuItem
	PopMenuAddTempList *vcl.TMenuItem
	PopMenuBlockList   *vcl.TMenuItem
	PopMenuKick        *vcl.TMenuItem
}

type TBlockListPopupMenu struct {
	*vcl.TPopupMenu

	PopMenuRefList     *vcl.TMenuItem
	PopMenuSort        *vcl.TMenuItem
	PopMenuAdd         *vcl.TMenuItem
	PopMenuAddTempList *vcl.TMenuItem
	PopMenuDelete      *vcl.TMenuItem
}

type TTempBlockListPopupMenu struct {
	*vcl.TPopupMenu

	PopMenuRefList   *vcl.TMenuItem
	PopMenuSort      *vcl.TMenuItem
	PopMenuAdd       *vcl.TMenuItem
	PopMenuBlockList *vcl.TMenuItem
	PopMenuDelete    *vcl.TMenuItem
}

func NewActiveListPopupMenu(sender vcl.IComponent) *TActiveListPopupMenu {
	sf := &TActiveListPopupMenu{
		TPopupMenu: vcl.NewPopupMenu(sender),
	}

	sf.PopMenuRefList = vcl.NewMenuItem(sf)
	sf.PopMenuRefList.SetCaption("刷新(&R)")

	sf.PopMenuSort = vcl.NewMenuItem(sf)
	sf.PopMenuSort.SetCaption("排序(&S)")

	sf.PopMenuAddTempList = vcl.NewMenuItem(sf)
	sf.PopMenuAddTempList.SetCaption("加入动态过滤列表(&A)")

	sf.PopMenuBlockList = vcl.NewMenuItem(sf)
	sf.PopMenuBlockList.SetCaption("加入永久过滤列表(&D)")

	sf.PopMenuKick = vcl.NewMenuItem(sf)
	sf.PopMenuKick.SetCaption("踢除下线(&K)")

	sf.Items().Add(sf.PopMenuRefList)
	sf.Items().Add(sf.PopMenuSort)
	sf.Items().Add(sf.PopMenuAddTempList)
	sf.Items().Add(sf.PopMenuBlockList)
	sf.Items().Add(sf.PopMenuKick)

	return sf
}

func NewBlockListPopupMenu(sender vcl.IComponent) *TBlockListPopupMenu {
	sf := &TBlockListPopupMenu{
		TPopupMenu: vcl.NewPopupMenu(sender),
	}

	sf.PopMenuRefList = vcl.NewMenuItem(sf)
	sf.PopMenuRefList.SetCaption("刷新(&R)")

	sf.PopMenuSort = vcl.NewMenuItem(sf)
	sf.PopMenuSort.SetCaption("排序(&S)")

	sf.PopMenuAdd = vcl.NewMenuItem(sf)
	sf.PopMenuAdd.SetCaption("增加(&A)")

	sf.PopMenuAddTempList = vcl.NewMenuItem(sf)
	sf.PopMenuAddTempList.SetCaption("加入动态过滤列表(&A)")

	sf.PopMenuDelete = vcl.NewMenuItem(sf)
	sf.PopMenuDelete.SetCaption("删除(&D)")

	sf.Items().Add(sf.PopMenuRefList)
	sf.Items().Add(sf.PopMenuSort)
	sf.Items().Add(sf.PopMenuAdd)
	sf.Items().Add(sf.PopMenuAddTempList)
	sf.Items().Add(sf.PopMenuDelete)

	return sf
}

func NewTempBlockListPopupMenu(sender vcl.IComponent) *TTempBlockListPopupMenu {
	sf := &TTempBlockListPopupMenu{
		TPopupMenu: vcl.NewPopupMenu(sender),
	}

	sf.PopMenuRefList = vcl.NewMenuItem(sf)
	sf.PopMenuRefList.SetCaption("刷新(&R)")

	sf.PopMenuSort = vcl.NewMenuItem(sf)
	sf.PopMenuSort.SetCaption("排序(&S)")

	sf.PopMenuAdd = vcl.NewMenuItem(sf)
	sf.PopMenuAdd.SetCaption("增加(&A)")

	sf.PopMenuBlockList = vcl.NewMenuItem(sf)
	sf.PopMenuBlockList.SetCaption("加入永久过滤列表(&D)")

	sf.PopMenuDelete = vcl.NewMenuItem(sf)
	sf.PopMenuDelete.SetCaption("删除(&D)")

	sf.Items().Add(sf.PopMenuRefList)
	sf.Items().Add(sf.PopMenuSort)
	sf.Items().Add(sf.PopMenuAdd)
	sf.Items().Add(sf.PopMenuBlockList)
	sf.Items().Add(sf.PopMenuDelete)

	return sf
}
