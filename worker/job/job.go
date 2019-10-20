package job

type Job interface {
	Process()
	WaitCancel()
	HasCancel() bool
}
