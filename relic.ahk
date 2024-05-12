#Requires AutoHotkey v2.0

#SingleInstance Force

; https://stackoverflow.com/questions/43298908/how-to-add-administrator-privileges-to-autohotkey-script
if not (A_IsAdmin or RegExMatch(DllCall("GetCommandLine", "str"), " /restart(?!\S)"))
{
  try
  {
    if A_IsCompiled
      Run '*RunAs "' A_ScriptFullPath '" /restart'
    else
      Run '*RunAs "' A_AhkPath '" /restart "' A_ScriptFullPath '"'
  }
  ExitApp
}

; SetMouseDelay 300

click_with_sleep(x, y){
  click(x, y)
  sleep(300)
}

click_lock(){
  click_with_sleep(3623, 541) ; 锁定按钮
}

!r::
    {
        Reload
    }


!s::
    {
        while (true){
          j := 4
          while (j >= 0){
              i := 0
              while (i < 9){
                  pos_x := 360 + 260 * i
                  pos_y := 508 + 300 * j
                  click_with_sleep(pos_x, pos_y)
                  click_lock()
                  i := i + 1
              }
              j := j - 1
          }

          click_with_sleep(360, 508)
          Send "{WheelDown 25}"
          sleep(300)
        }
    }

!d:: 
    {
    j := 0
    while (j < 5){
        i := 0
        while (i < 9){
            pos_x := 360 + 260 * i
            pos_y := 508 + 300 * j
            click_with_sleep(pos_x, pos_y)
            click_lock()
            i := i + 1
        }
        j := j + 1
    }
  }