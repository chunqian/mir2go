
// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

unit Mir.LogDataServer;

{$Mode objfpc}
{$H+}

// ******************** interface ********************
interface

uses SysUtils, Classes, Forms, Buttons;

type
  TFrmLogData = class(TForm)
  public
    constructor Create(Sender: TComponent); override;
  end;

var
  FrmLogData: TFrmLogData;

// ******************** implementation ********************
implementation

// 
constructor TFrmLogData.Create(Sender: TComponent);
begin
  inherited CreateNew(Sender, 1);
  Caption := 'LogDataServer';
  Width := 200;
  Height := 75;
  Left := 200;
  Top := 200;

  Self.Constraints.MaxWidth:= 500; 
end;

end.
