^#up::MouseMove, 0, -10, 0, R  ; Win+Ctrl+UpArrow hotkey => Move cursor upward
^#Down::MouseMove, 0, 10, 0, R  ; Win+Ctrl+DownArrow => Move cursor downward
^#Left::MouseMove, -10, 0, 0, R  ; Win+Ctrl+LeftArrow => Move cursor to the left
^#Right::MouseMove, 10, 0, 0, R  ; Win+Ctrl+RightArrow => Move cursor to the right

*<#RCtrl::  ; LeftWin + RightControl => Left-click (hold down Control/Shift to Control-Click or Shift-Click).
SendEvent {Blind}{LButton down}
KeyWait RCtrl  ; Prevents keyboard auto-repeat from repeating the mouse click.
SendEvent {Blind}{LButton up}
return