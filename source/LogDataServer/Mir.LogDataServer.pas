
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
  ExtCtrls,
  StdCtrls,
  IniFiles;

type
  TFrmLogData = class(TForm)
    protected
      Label1: TLabel;
      Timer1: TTimer;
      Memo1: TMemo;
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

var
  sBaseDir: string = '.\LogBase';
  sServerName: string = '热血传奇';
  sCaption: string = '引擎日志服务器';

  nServerPort: Integer = 10000;

// ******************** implementation ********************
implementation

constructor TFrmLogData.Create(TheOwner: TComponent);
var
  Conf: TIniFile;
begin
  inherited CreateNew(TheOwner, 1);
  Caption := '日志服务器';
  Width := 329;
  Height := 121;
  Left := 782;
  Top := 338;

  // Label1
  Label1 := TLabel.Create(Self);
  Label1.Parent := Self;
  Label1.Caption := '当前日志文件:';
  Label1.Top := 9;
  Label1.Left := 9;
  Label1.Height := 13;
  Label1.Width := 85;

  // Timer1
  Timer1 := TTimer.Create(Self);
  Timer1.Interval := 3000;
  Timer1.Enabled := True;
  Timer1.OnTimer := @Timer1Timer;

  // Memo1
  Memo1 := TMemo.Create(Self);
  Memo1.Parent := Self;
  Memo1.Top := 30;
  Memo1.Left := 11;
  Memo1.Height := 75;
  Memo1.Width := 303;
  Memo1.ReadOnly := True;

  Constraints.MaxWidth:= 500;

  OnCloseQuery := @FormCloseQuery;

  Conf := TIniFile.Create('.\logdata.ini');
  if Conf <> nil then begin
    sBaseDir := Conf.ReadString('Setup', 'BaseDir', sBaseDir);
    sServerName := Conf.ReadString('Setup', 'Caption', sServerName);
    sServerName := Conf.ReadString('Setup', 'ServerName', sServerName);
    nServerPort := Conf.ReadInteger('Setup', 'Port', nServerPort);
    Conf.Free;
  end;
  Caption := sCaption + ' - ' + sServerName;

  Memo1.Text := sBaseDir;
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
