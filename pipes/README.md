# Logging through pipes

This functionality is needed when a parent process starts a child process and both their logs have to be collected in the parent process.

## The parent process:

```
part, _ := pipes.NewParentPart(marshalizer)
profileReader, logsWriter := part.GetChildPipes()

command = exec.Command("child.bin")
childStdout, _:= command.StdoutPipe()
childStderr, _ := command.StderrPipe()
command.ExtraFiles = []*os.File{
		...,
		profileReader,
		logsWriter,
}

_ = part.StartLoop(childStdout, childStderr)
```

`StartLoop` will continuously read log lines from the child  (pipe `logsWriter`) on a separate goroutine. Child's `stdout` and `stderr` are also captured. `Stdout` will be logged with `trace` level, while `stderr` with `error` level.


Furthermore, the parent part forwards log profile changes to the child process (through pipe `profileReader`).

Note that the parent process is responsible to call `logger.NotifyProfileChange()` when it applies a new log profile (whether by sole choice or when instructed by a logviewer).

## The child process

```
profileReader := os.NewFile(42, "/proc/self/fd/42")
logsWriter := os.NewFile(43, "/proc/self/fd/43")
part, _ := pipes.NewChildPart(profileReader, logsWriter, marshalizer)
_ = part.StartLoop()
```

The child process has to aquire the provided pipes, create its part of the logging dialogue and then call `StartLoop`.
The child part is automatically registered as observer to the global default `LogOutputSubject`, which means that it gets notified on each log write from any of the loggers in the process. When notified, the child part simply forwards the message (the serialized log line) to its parent, through pipe `logsWriter`. 

Furthermore, the child part listens for eventual log profile changes on the pipe `profileReader`. Any profile change is applied immediately.
