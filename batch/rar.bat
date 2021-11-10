mkdir Secure
dotNET_Reactor.exe -project ..\..\..\AGV.nrproj -mapping_file 1 -obfuscation 1 -q -file AGV.exe
dotNET_Reactor.exe -mapping_file 1 -obfuscation 1 -q -file AGV.Device.dll -targetfile Secure -satellite_assemblies "AGV.Commom.dll/AGV.Function.dll/AGVMap.dll"
del /a /f /s *.rar
copy AGVTest.exe  Secure\AGVTest.exe
cd Secure
del /a /f /s *.nrmap
del /a /f /s *.rar
del /a /f /s *.pdb
set WinRarDir=%ProgramFiles%\WinRAR
cd ..\
@echo off
For /f "tokens=1-3 delims=/ " %%a in ('date /t') do (set mydate=%%a%%b%%c)
For /f "tokens=1-2 delims=/:" %%a in ('time /t') do (set mytime=%%a%%b)
echo %mydate%_%mytime%
rar a -r test%mydate%_%mytime%.rar Secure\*
pause