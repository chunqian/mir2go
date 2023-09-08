
// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

program main;

{$Mode objfpc}
{$H+}

uses Interfaces, Forms, Mir.LogDataServer;

begin
  Application.Initialize;
  Application.CreateForm(TFrmLogData, FrmLogData);
  Application.Run;
end.
