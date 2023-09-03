
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
  Dialogs,
  ExtCtrls;

type
  TFrmLogData = class(TForm)
    protected
      Timer1: TTimer;
      procedure Timer1Timer(Sender: TObject);

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

constructor TFrmLogData.Create(TheOwner: TComponent);
begin
  inherited CreateNew(TheOwner, 1);
  Caption := '日志服务器';
  Width := 200;
  Height := 75;
  Left := 200;
  Top := 200;

  // Timer1
  Timer1 := TTimer.Create(Self);
  Timer1.Interval := 3000;
  Timer1.Enabled := True;
  Timer1.OnTimer := @Timer1Timer;

  Constraints.MaxWidth:= 500;

  OnCloseQuery := @FormCloseQuery;
end;

destructor TFrmLogData.Destroy();
begin
  inherited;
end;

procedure TFrmLogData.FormCloseQuery(Sender: TObject; var CanClose: Boolean);
var
  mr: Integer;
begin
  if m_boRemoteClose then exit;
  mr := MessageDlg( '提示信息',
                    '是否确认退出服务器?',
                    mtConfirmation, [mbYes, mbNo], 0);

  if mr = mrYes then
  begin
    // 确认退出操作
  end
  else begin
    CanClose := False;
  end;
end;

procedure TFrmLogData.Timer1Timer(Sender: TObject);
begin
  // 
end;

end.
