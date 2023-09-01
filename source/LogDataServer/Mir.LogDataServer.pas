
// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

unit Mir.LogDataServer;

{$Mode objfpc}
{$H+}

// ******************** interface ********************
interface

uses
  SysUtils,
  Classes,
  Forms,
  Buttons,
  Controls,
  Dialogs;

type
  TFrmLogData = class(TForm)
    public
      constructor Create(TheOwner: TComponent); override;
      destructor Destroy(); override;
      procedure FormCloseQuery(Sender: TObject; var CanClose: Boolean);

    private
      LogMsgList: TStringList;
      m_boRemoteClose: Boolean;
  end;

var
  FrmLogData: TFrmLogData;

// ******************** implementation ********************
implementation

// 
constructor TFrmLogData.Create(TheOwner: TComponent);
begin
  inherited CreateNew(TheOwner, 1);
  Caption := 'LogDataServer';
  Width := 200;
  Height := 75;
  Left := 200;
  Top := 200;

  Constraints.MaxWidth:= 500;

  OnCloseQuery := @FormCloseQuery;
end;

// 
destructor TFrmLogData.Destroy();
begin
  inherited;
end;

procedure TFrmLogData.FormCloseQuery(Sender: TObject; var CanClose: Boolean);
var
  mr: Integer;
begin
  if m_boRemoteClose then exit;
  mr := MessageDlg('提示信息', '是否确认退出服务器?', mtConfirmation, [mbYes, mbNo], 0);

  if mr = mrYes then
  begin
    // 确认退出操作
  end
  else begin
    CanClose := False;
  end;
end;

end.
