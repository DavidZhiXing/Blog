public class Log4NetAdapterTest
{
    private static readonly ILog Log = loggerFactory.CreateLogger("Test");

    [Fact]
    public void Log4NetAdapterTest_Logs()
    {
        Log.Information("Log4NetAdapterTest_Logs");
        Log.Warning("Log4NetAdapterTest_Logs");
        Log.Error("Log4NetAdapterTest_Logs");
        Log.Fatal("Log4NetAdapterTest_Logs");
    }

}