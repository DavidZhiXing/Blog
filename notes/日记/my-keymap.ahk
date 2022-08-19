
#NoEnv  ; Recommended for performance and compatibility with future AutoHotkey releases.
; #Warn  ; Enable warnings to assist with detecting common errors.
SendMode Input  ; Recommended for new scripts due to its superior speed and reliability.
SetWorkingDir %A_ScriptDir%  ; Ensures a consistent starting directory.

;=====================================================================o
;                    CapsLock Direction Navigator                    ;|
;-----------------------------------o---------------------------------o
;                      CapsLock + j |  Left                          ;|
;                      CapsLock + k |  Down                          ;|
;                      CapsLock + i |  Up                            ;|
;                      CapsLock + l |  Right                         ;|
;                      Ctrl, Alt Compatible                          ;|
;-----------------------------------o---------------------------------o
CapsLock & j::                                                       ;|
if GetKeyState("control") = 0                                        ;|
{                                                                    ;|
    if GetKeyState("shift") = 0                                        ;|
        Send, {Left}                                                 ;|
    else                                                             ;|
        Send, +{Left}                                                ;|
    return                                                           ;|
}                                                                    ;|
else {                                                               ;|
    if GetKeyState("shift") = 0                                        ;|
        Send, ^{Left}                                                ;|
    else                                                             ;|
        Send, +^{Left}                                               ;|
    return                                                           ;|
}                                                                    ;|
return                                                               ;|
;-----------------------------------o                                ;|
CapsLock & k::                                                       ;|
if GetKeyState("control") = 0                                        ;|
{                                                                    ;|
    if GetKeyState("shift") = 0                                        ;|
        Send, {Down}                                                 ;|
    else                                                             ;|
        Send, +{Down}                                                ;|
    return                                                           ;|
}                                                                    ;|
else {                                                               ;|
    if GetKeyState("shift") = 0                                        ;|
        Send, ^{Down}                                                ;|
    else                                                             ;|
        Send, +^{Down}                                               ;|
    return                                                           ;|
}                                                                    ;|
return                                                               ;|
;-----------------------------------o                                ;|
CapsLock & i::                                                       ;|
if GetKeyState("control") = 0                                        ;|
{                                                                    ;|
    if GetKeyState("shift") = 0                                        ;|
        Send, {Up}                                                   ;|
    else                                                             ;|
        Send, +{Up}                                                  ;|
    return                                                           ;|
}                                                                    ;|
else {                                                               ;|
    if GetKeyState("shift") = 0                                        ;|
        Send, ^{Up}                                                  ;|
    else                                                             ;|
        Send, +^{Up}                                                 ;|
    return                                                           ;|
}                                                                    ;|
return                                                               ;|
;-----------------------------------o                                ;|
CapsLock & l::                                                       ;|
if GetKeyState("control") = 0                                        ;|
{                                                                    ;|
    if GetKeyState("shift") = 0                                        ;|
        Send, {Right}                                                ;|
    else                                                             ;|
        Send, +{Right}                                               ;|
    return                                                           ;|
}                                                                    ;|
else {                                                               ;|
    if GetKeyState("shift") = 0                                        ;|
        Send, ^{Right}                                               ;|
    else                                                             ;|
        Send, +^{Right}                                              ;|
    return                                                           ;|
}                                                                    ;|
return                                                               ;|
;---------------------------------------------------------------------o


;=====================================================================o
;                     CapsLock Home/End Navigator                    ;|
;-----------------------------------o---------------------------------o
;                      CapsLock + u |  Home                          ;|
;                      CapsLock + o |  End                           ;|
;                      Ctrl, Alt Compatible                          ;|
;-----------------------------------o---------------------------------o
CapsLock & u::                                                       ;|
if GetKeyState("control") = 0                                        ;|
{                                                                    ;|
    if GetKeyState("shift") = 0                                        ;|
        Send, {Home}                                                 ;|
    else                                                             ;|
        Send, +{Home}                                                ;|
    return                                                           ;|
}                                                                    ;|
else {                                                               ;|
    if GetKeyState("shift") = 0                                        ;|
        Send, ^{Home}                                                ;|
    else                                                             ;|
        Send, +^{Home}                                               ;|
    return                                                           ;|
}                                                                    ;|
return                                                               ;|
;-----------------------------------o                                ;|
CapsLock & o::                                                       ;|
if GetKeyState("control") = 0                                        ;|
{                                                                    ;|
    if GetKeyState("shift") = 0                                        ;|
        Send, {End}                                                  ;|
    else                                                             ;|
        Send, +{End}                                                 ;|
    return                                                           ;|
}                                                                    ;|
else {                                                               ;|
    if GetKeyState("shift") = 0                                        ;|
        Send, ^{End}                                                 ;|
    else                                                             ;|
        Send, +^{End}                                                ;|
    return                                                           ;|
}                                                                    ;|
return                                                               ;|
;---------------------------------------------------------------------o

CapsLock & h::Send, {blind}{home}
CapsLock & y::Send, {blind}{PgUp}
CapsLock & n::Send, {blind}{PgDn}
CapsLock & q::Send, !{F4}


CapsLock & s::
	 SendInput, {End}
	 SendInput, {Enter}
return

;=====================================================================o
;                     CapsLock Mouse Controller                      ;|
;-----------------------------------o---------------------------------o
;                   CapsLock + Up   |  Mouse Up                      ;|
;                   CapsLock + Down |  Mouse Down                    ;|
;                   CapsLock + Left |  Mouse Left                    ;|
;                  CapsLock + Right |  Mouse Right                   ;|
;    CapsLock + Enter(Push Release) |  Mouse Left Push(Release)      ;|
;-----------------------------------o---------------------------------o
CapsLock & Up::    
if GetKeyState("control") = 0                                        ;|
{                                                                    ;|
    if GetKeyState("shift") = 0                                        ;|
        MouseMove, 0, -10, 0, R                                                ;|
    else                                                             ;|
        MouseMove, 0, -100, 0, R                                               ;|
    return                                                           ;|
}                                                                    ;|
else {                                                               ;|
    if GetKeyState("shift") = 0                                        ;|
        MouseMove, 0, -200, 0, R                                               ;|
    else                                                             ;|
        MouseMove, 0, -300, 0, R                                              ;|
    return                                                           ;|
}                                                                    ;|
return 

                           ;|
CapsLock & Down::  
if GetKeyState("control") = 0                                        ;|
{                                                                    ;|
    if GetKeyState("shift") = 0                                        ;|
        MouseMove, 0, 10, 0, R                                                ;|
    else                                                             ;|
        MouseMove, 0, 100, 0, R                                               ;|
    return                                                           ;|
}                                                                    ;|
else {                                                               ;|
    if GetKeyState("shift") = 0                                        ;|
        MouseMove, 0, 200, 0, R                                               ;|
    else                                                             ;|
        MouseMove, 0, 300, 0, R                                              ;|
    return                                                           ;|
}                                                                    ;|
return 

                            ;|
CapsLock & Left::
if GetKeyState("control") = 0                                        ;|
{                                                                    ;|
    if GetKeyState("shift") = 0                                        ;|
        MouseMove, -10, 0, 0, R                                                ;|
    else                                                             ;|
        MouseMove, -100, 0, 0, R                                               ;|
    return                                                           ;|
}                                                                    ;|
else {                                                               ;|
    if GetKeyState("shift") = 0                                        ;|
        MouseMove, -200, 0, 0, R                                               ;|
    else                                                             ;|
        MouseMove, -300, 0, 0, R                                              ;|
    return                                                           ;|
}                                                                    ;|
return 

                             ;|
CapsLock & Right::
if GetKeyState("control") = 0                                        ;|
{                                                                    ;|
    if GetKeyState("shift") = 0                                        ;|
        MouseMove, 10, 0, 0, R                                                 ;|
    else                                                             ;|
        MouseMove, 100, 0, 0, R                                                ;|
    return                                                           ;|
}                                                                    ;|
else {                                                               ;|
    if GetKeyState("shift") = 0                                        ;|
        MouseMove, 200, 0, 0, R                                                ;|
    else                                                             ;|
        MouseMove, 300, 0, 0, R                                               ;|
    return                                                           ;|
}                                                                    ;|
return 

                            ;|
;-----------------------------------o                                ;|
CapsLock & Enter::
if GetKeyState("Control") = 0
{                                                   ;|
    SendEvent {Blind}{LButton down}                                      ;|
    KeyWait Enter                                                        ;|
    SendEvent {Blind}{LButton up}                                        ;|
    return
}   
else
{
    SendEvent {Blind}{RButton down}                                      ;|
    KeyWait Enter                                                        ;|
    SendEvent {Blind}{RButton up}                                        ;|
    return
}                                                            ;|
;---------------------------------------------------------------------o

;=====================================================================o
;                            CapsLock Editor                         ;|
;-----------------------------------o---------------------------------o
;                     CapsLock + z  |  Ctrl + z (Cancel)             ;|
;                     CapsLock + x  |  Ctrl + x (Cut)                ;|
;                     CapsLock + c  |  Ctrl + c (Copy)               ;|
;                     CapsLock + v  |  Ctrl + z (Paste)              ;|
;                     CapsLock + a  |  Ctrl + a (Select All)         ;|
;                     CapsLock + y  |  Ctrl + z (Yeild)              ;|
;                     CapsLock + w  |  Ctrl + Right(Move as [vim: w]);|
;                     CapsLock + b  |  Ctrl + Left (Move as [vim: b]);|
;-----------------------------------o---------------------------------o
CapsLock & z:: Send, ^z                                              ;|
CapsLock & x:: Send, ^x                                              ;|
CapsLock & c:: Send, ^c                                              ;|
CapsLock & v:: Send, ^v                                              ;|
CapsLock & a:: Send, ^a                                              ;|
CapsLock & r:: Send, ^y                                              ;|
CapsLock & w:: Send, ^{Right}                                        ;|
CapsLock & b:: Send, ^{Left}                                         ;|
;---------------------------------------------------------------------o


;=====================================================================o
;                       CapsLock Media Controller                    ;|
;-----------------------------------o---------------------------------o
;                    CapsLock + F1  |  Volume_Mute                   ;|
;                    CapsLock + F2  |  Volume_Down                   ;|
;                    CapsLock + F3  |  Volume_Up                     ;|
;                    CapsLock + F3  |  Media_Play_Pause              ;|
;                    CapsLock + F5  |  Media_Next                    ;|
;                    CapsLock + F6  |  Media_Stop                    ;|
;-----------------------------------o---------------------------------o
CapsLock & F1:: Send, {Volume_Mute}                                  ;|
CapsLock & F2:: Send, {Volume_Down}                                  ;|
CapsLock & F3:: Send, {Volume_Up}                                    ;|
CapsLock & F4:: Send, {Media_Play_Pause}                             ;|
CapsLock & F5:: Send, {Media_Next}                                   ;|
CapsLock & F6:: Send, {Media_Stop}                                   ;|
;---------------------------------------------------------------------o


;=====================================================================o
;                      CapsLock Window Controller                    ;|
;-----------------------------------o---------------------------------o
;                     CapsLock + s  |  Ctrl + s (Save)        ;|
;                     CapsLock + q  |  Ctrl + W   (Close Tag)        ;|
;   (Disabled)  Alt + CapsLock + s  |  AltTab     (Switch Windows)   ;|
;               Alt + CapsLock + q  |  Ctrl + Tab (Close Windows)    ;|
;                     CapsLock + g  |  AppsKey    (Menu Key)         ;|
;-----------------------------------o---------------------------------o
;CapsLock & s::Send, ^s                                           ;|
;-----------------------------------o                                ;|
;CapsLock & q::                                                       ;|
;if GetKeyState("alt") = 0                                            ;|
;{                                                                    ;|
;    Send, ^w                                                         ;|
;}                                                                    ;|
;else {                                                               ;|
;    Send, !{F4}                                                      ;|
;    return                                                           ;|
;}                                                                    ;|
;return                                                               ;|
;-----------------------------------o                                ;|
CapsLock & g:: Send, {AppsKey}                                       ;|
;---------------------------------------------------------------------o


;=====================================================================o
;                        CapsLock Self Defined Area                  ;|
;-----------------------------------o---------------------------------o
;                     CapsLock + d  |  Alt + d(Dictionary)           ;|
;                     CapsLock + f  |  Alt + f(Search via Everything);|
;                     CapsLock + e  |  Open Search Engine            ;|
;                     CapsLock + r  |  Open Shell                    ;|
;                     CapsLock + t  |  Open Text Editor              ;|
;-----------------------------------o---------------------------------o
; CapsLock & d:: Send, !d                                              ;|
; CapsLock & f:: Send, !f                                              ;|
CapsLock & e:: Run http://cn.bing.com/                               ;|
CapsLock & t:: Run Powershell                                        ;|
;CapsLock & t:: Run C:\Program Files (x86)\Notepad++\notepad++.exe    ;|
;---------------------------------------------------------------------o


;=====================================================================o
;                        CapsLock Char Mapping                       ;|
;-----------------------------------o---------------------------------o
;                     CapsLock + ;  |  Enter (Cancel)                ;|
;                     CapsLock + '  |  =                             ;|
;                     CapsLock + [  |  Back         (VSCODE)  ;|
;                     CapsLock + ]  |  Goto Define  (VSCODE)  ;|
;                     CapsLock + /  |  Comment      (VSCODE)  ;|
;                     CapsLock + \  |  Uncomment    (VSCODE)  ;|
;                     CapsLock + 1  |  Previous Editor(VSCODE)  ;|
;                     CapsLock + 2  |  Next Editor     (VSCODE)  ;|
;                     CapsLock + 3  |  Step Over    (VSCODE)  ;|
;                     CapsLock + 4  |  Step In      (VSCODE)  ;|
;                     CapsLock + 5  |  Stop Debuging(VSCODE)  ;|
;                     CapsLock + 6  |  Shift + 6     ^               ;|
;                     CapsLock + 7  |  Shift + 7     &               ;|
;                     CapsLock + 8  |  Shift + 8     *               ;|
;                     CapsLock + 9  |  Shift + 9     (               ;|
;                     CapsLock + 0  |  Shift + 0     )               ;|
;-----------------------------------o---------------------------------o
CapsLock & `;:: Send, {Enter}                                        ;|
CapsLock & ':: Send, =                                               ;|
CapsLock & [:: Send, !{left}                                              ;|
CapsLock & ]:: Send, {F12}                                           ;|
;-----------------------------------o                                ;|
CapsLock & /::                                                       ;|
Send, ^e                                                             ;|
Send, c                                                              ;|
return                                                               ;|
;-----------------------------------o                                ;|
CapsLock & \::                                                       ;|
Send, ^e                                                             ;|
Send, u                                                              ;|
return                                                               ;|
;-----------------------------------o                                ;|
CapsLock & 1:: Send,^{PgUp}                                            ;|
CapsLock & 2:: Send,^{PgDn}                                             ;|
CapsLock & 3:: Send,{F10}                                            ;|
CapsLock & 4:: Send,{F11}                                            ;|
CapsLock & 5:: Send,+{F5}                                            ;|
;-----------------------------------o                                ;|
CapsLock & 6:: Send,+6                                               ;|
CapsLock & 7:: Send,+7                                               ;|
CapsLock & 8:: Send,+8                                               ;|
CapsLock & 9:: Send,+9                                               ;|
CapsLock & 0:: Send,+0                                               ;|
;---------------------------------------------------------------------o

;=====================================================================o
;                       CapsLock Switcher:                           ;|
;---------------------------------o-----------------------------------o
;                    CapsLock + ` | {CapsLock}                       ;|
;---------------------------------o-----------------------------------o
CapsLock & `::                                                       ;|
GetKeyState, CapsLockState, CapsLock, T                              ;|
if CapsLockState = D                                                 ;|
    SetCapsLockState, AlwaysOff                                      ;|
else                                                                 ;|
    SetCapsLockState, AlwaysOn                                       ;|
KeyWait, ``                                                          ;|
return                                                               ;|
;---------------------------------------------------------------------o


;=====================================================================o
;                         CapsLock Escaper:                          ;|
;----------------------------------o----------------------------------o
;                        CapsLock  |  {ESC}                          ;|
;----------------------------------o----------------------------------o
CapsLock::Send, {ESC}                                                ;|
;---------------------------------------------------------------------o


;CAPSLOCK & w::LWin

; Speed
; CapsLock & 8::Send, {Up 5}
CapsLock & m::Send, {blind}^{Left 6}
CapsLock & .::Send, {blind}^{Right 6}
CapsLock & ,::Send, {Down 5}

; since Alt is in position of Apple Cmd key these are usually:
;#IfWinNoExist, Emacs
;{
    ; !w::Send, ^w ; close window
    ; !r::Send, ^r ; refresh
    ; !c::Send, ^c ; copy
    ; !x::Send, ^x ; cut
    ; !v::Send, ^v ; paste
    ; !s::Send, ^s ; save
    ; !l::Send, ^l ; focus address bar
    ; !z::Send, ^z ; undo
    ; !+z::Send, ^+z ; redo
    ; ; move wordly
    ; !u::Send, ^{Left}
    ; !o::Send, ^{Right}
    ; !+u::Send, ^+{Left}
    ; !+o::Send, ^+{Right}
    ; ; move to ends
    ; !h::Send, {Home}
    ; !;::Send, {End}
    ; !+h::Send, +{Home}
    ; !+;::Send, +{End}
; }

; Keep window open
CapsLock & +::Winset, Alwaysontop, , A

; Change Case
;CapsLock & 9::goSub, set_upper_case
;CapsLock & 0::goSub, set_lower_case
;CapsLock & -::goSub, set_title_case

set_upper_case:
set_lower_case:
set_title_case:
revert_clipboard := clipboardAll
clipboard =
send ^{c}
clipWait, 0.3

if (a_thisLabel = "set_upper_case")
    stringUpper, clipboard, clipboard
else if (a_thisLabel = "set_lower_case")
    stringLower, clipboard, clipboard
else if (a_thisLabel = "set_title_case")
    stringLower, clipboard, clipboard, T

send ^{v}
sleep 50
clipboard := revert_clipboard
return