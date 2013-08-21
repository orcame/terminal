package terminal
import(
	"os"
	"syscall"
)

var (
	Stdout =getStdout()
	Stderr =getStderr()
)

func getStdout() *writer{
	return createWriter(os.Stdout,uintptr(syscall.Stdout))
}

func getStderr() *writer{
	return createWriter(os.Stderr,uintptr(syscall.Stderr))
}