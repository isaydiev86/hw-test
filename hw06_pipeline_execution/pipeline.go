package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	if in == nil {
		return nil
	}

	for _, stage := range stages {
		mCh := make(Bi)
		go func(ch In) {
			defer close(mCh)
			for {
				select {
				case item, ok := <-ch:
					if !ok {
						return
					}
					mCh <- item
				case <-done:
					return
				}
			}
		}(in)

		in = stage(mCh)
	}
	return in
}
