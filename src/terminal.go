package terminal
import(
	"os"
	"syscall"
)

var (
	Stdout =getStdout()
	Stderr =getStderr()
)

func getStdout() writer{

	//todo: check the platform and return win32/linux writer.
	return createWin32Writer(os.Stdout,uintptr(syscall.Stdout))
}

func getStderr() writer{
	return createWin32Writer(os.Stderr,uintptr(syscall.Stderr))
}