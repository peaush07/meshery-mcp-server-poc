Output demo.gif

Set FontSize 15
Set Width 1200
Set Height 700
Set Padding 20
Set Theme "TokyoNight"
Set CornerRadius 10

Type "cd ~/meshery-mcp-server-poc" Enter
Sleep 1s

Type "make test" Enter
Sleep 3.5s

Type "make bench" Enter
Sleep 4.5s

Type "make build" Enter
Sleep 2s

Type "./bin/meshery-mcp-server -transport=stdio -help" Enter
Sleep 3s
