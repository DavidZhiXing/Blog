public void ThrottledLoop(Action action, int cpuPercentageLimit) {
    Stopwatch stopwatch = new Stopwatch();

    while(true) {
        stopwatch.Reset();
        stopwatch.Start();

        long actionStart = stopwatch.ElapsedTicks;
        action.Invoke();
        long actionEnd = stopwatch.ElapsedTicks;
        long actionDuration = actionEnd - actionStart;

        long relativeWaitTime = (int)(
            (1/(double)cpuPercentageLimit) * actionDuration);

        Thread.Sleep((int)((relativeWaitTime / (double)Stopwatch.Frequency) * 1000));
    }
}


public void ThrottledLoop(Action action, int cpuPercentageLimit) {
    Stopwatch stopwatch = new Stopwatch();

    while(true){
        stopwatch.Reset();
        stopwatch.Start();

        long actionStart = stopwatch.ElapsedTicks;
        action.Invoke();
        long actionEnd = stopwatch.ElapsedTicks;
        long actionDuration = actionEnd - actionStart;

        long relativeWaitTime = (int)(
            (1/(double)cpuPercentageLimit) * actionDuation;
        )

        Thread.Sleep((int)((relativeWaitTime / (double)Stopwatch.Frequency)*1000));
    }
}

public void Test(){
    
}
