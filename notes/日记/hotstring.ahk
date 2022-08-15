;sends the db_home variable
:::db::
SendRaw, `%DB_HOME`%\
Return

;Sends the BOX_HOME variable
:::bx::
SendRaw, `%BOX_HOME`%\
Return

;time and date stamp
::/now::
FormatTime, Time, ,yyyy-MM-dd--HH-mm
SendInput, %Time%
Return

;time stamp
::ttt::
FormatTime, Time, ,HH:mm
SendInput, %Time%
Return

;date stamp
::ddd::
FormatTime, Time, ,yyyy-MM-dd
SendInput, %Time%
Return

;MLA Header
:::mla::
FormatTime, Time, ,yyyy-MM-dd
SendInput, Nate Mara
Send, {ENTER}
SendInput, Prof. Amir Hassan
Send, {ENTER}
SendInput, ENG 112
Send, {ENTER}
SendInput, %Time%
Return

;Lorem Ipsum
:::Lorem::
SendInput, Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
Return

;Code header
:::cse::
FormatTime, Time, ,yyyy-MM-dd
InputBox, title, Program, Enter the title of this program:
InputBox, does, Program, What does the program do?
SendInput, /*
Send, {ENTER}
SendInput, Nate Mara
Send, {ENTER}
Send, %Time%
Send, {ENTER}
Send, {ENTER}
Send, "%title%"
Send, {ENTER}
Send, {ENTER}
Send, This program %does%
Send, {ENTER}
Send, {ENTER}
SendInput, */
Return

;Contact Info
:::college::
Send, https://
Return

;defines function for sending unicode charaters
SendU(argString)
{
  oldClipboard = %clipboard%
  Clipboard = %argString%
  Send ^v
  Clipboard = %oldClipboard%
}
