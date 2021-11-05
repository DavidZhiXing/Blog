public class Log4NetAdapter : ILogger
{
    private readonly ILog _log;

    public Log4NetAdapter(string logname)
    {
        _log = LogManager.GetLogger(logname);
    }

    public IDisposable BeginScope<TState>(TState state)
    {
        return _log.BeginScope(state);
    }

    public bool IsEnabled(LogLevel logLevel)
    {
        switch (logLevel)
        {
            case LogLevel.Critical:
                return _log.IsFatalEnabled;
            case LogLevel.Debug:
            case LogLevel.Trace:
                return _log.IsDebugEnabled;
            case LogLevel.Error:
                return _log.IsErrorEnabled;
            case LogLevel.Information:
                return _log.IsInfoEnabled;
            case LogLevel.Warning:
                return _log.IsWarnEnabled;
            default:
                throw new ArgumentOutOfRangeException(nameof(logLevel));
        }
    }

    public void Log<TState>(LogLevel logLevel, EventId eventId, TState state, Exception exception, Func<TState, Exception, string> formatter)
    {
        if (!IsEnabled(logLevel))
        {
            return;
        }

        if (formatter != null)
        {
            message = formatter(state, exception);
        }
        else
        {
            message = LogFormatter.Formatter(state, exception);
        }

        switch (logLevel)
        {
            case LogLevel.Critical:
                _log.Fatal(state, exception);
                break;
            case LogLevel.Debug:
            case LogLevel.Trace:
                _log.Debug(state, exception);
                break;      
            case LogLevel.Error:
                _log.Error(state, exception);
                break;
            case LogLevel.Information:
                _log.Info(state, exception);
                break;
            case LogLevel.Warning:
                _log.Warn(state, exception);
                break;
            default:
                _log.Warn($"Encountered unknown log level {logLevel}, writing out as Info.");
                _log.Info(state, exception);
                break;
        }
}

public interface ILogger
{

    IDisposable BeginScope<TState>(TState state);

    bool IsEnabled(LogLevel logLevel);

    void Log<TState>(LogLevel logLevel, EventId eventId, TState state, Exception exception, Func<TState, Exception, string> formatter);
}

