#NoEnv  ; Recommended for performance and compatibility with future AutoHotkey releases.
; #Warn  ; Enable warnings to assist with detecting common errors.
SendMode Input  ; Recommended for new scripts due to its superior speed and reliability.
SetWorkingDir %A_ScriptDir%  ; Ensures a consistent starting directory.

; SetCapsLockState, AlwaysOff

CapsLock & i::Send, {blind}{Up}
CapsLock & j::Send, {blind}{Left}
CapsLock & l::Send, {blind}{Right}
CapsLock & k::Send, {blind}{Down}
CapsLock & u::Send, ^{left}
CapsLock & o::Send, ^{right}
CapsLock & h::Send, {blind}{home}
CapsLock & y::Send, {blind}{PgUp}
CapsLock & n::Send, {blind}{PgDn}
CapsLock & q::Send, !{F4}

CapsLock & p::Send, {blind}{Insert}
CapsLock & f::Send, {blind}{delete}
CapsLock & d::Send, {blind}{Backspace}

CapsLock & `;::
	 SendInput, {End}
return

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
CapsLock & Up::    MouseMove, 0, -10, 0, R                           ;|
CapsLock & Down::  MouseMove, 0, 10, 0, R                            ;|
CapsLock & Left::  MouseMove, -10, 0, 0, R                           ;|
CapsLock & Right:: MouseMove, 10, 0, 0, R                            ;|
;-----------------------------------o                                ;|
CapsLock & Enter::                                                   ;|
SendEvent {Blind}{LButton down}                                      ;|
KeyWait Enter                                                        ;|
SendEvent {Blind}{LButton up}                                        ;|
return                                                               ;|
;---------------------------------------------------------------------o



;CAPSLOCK & w::LWin

; Speed
CapsLock & 8::Send, {Up 5}
CapsLock & m::Send, {blind}^{Left 6}
CapsLock & .::Send, {blind}^{Right 6}
CapsLock & ,::Send, {Down 5}

; since Alt is in position of Apple Cmd key these are usually:
!w::Send, ^w ; close window
!r::Send, ^r ; refresh
!c::Send, ^c ; copy
!x::Send, ^x ; cut
!v::Send, ^v ; paste
!s::Send, ^s ; save
!l::Send, ^l ; focus address bar
!z::Send, ^z ; undo
!+z::Send, ^+z ; redo
; move wordly
!u::Send, ^{Left}
!o::Send, ^{Right}
!+u::Send, ^+{Left}
!+o::Send, ^+{Right}
; move to ends
!h::Send, {Home}
!;::Send, {End}
!+h::Send, +{Home}
!+;::Send, +{End}

; Keep window open
CapsLock & +::Winset, Alwaysontop, , A

; Change Case
CapsLock & 9::goSub, set_upper_case
CapsLock & 0::goSub, set_lower_case
CapsLock & -::goSub, set_title_case

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