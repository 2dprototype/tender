import "dll"

kernel32 := dll.new("kernel32")
user32 := dll.new("user32.dll")
GetConsoleWindow := kernel32.proc("GetConsoleWindow")
ShowWindowAsync := user32.proc("ShowWindowAsync")

export {
	hide : fn(){
		return ShowWindowAsync.call(GetConsoleWindow.call(), 0)
	},
	show : fn(){
		return ShowWindowAsync.call(GetConsoleWindow.call(), 4)
	}
}
