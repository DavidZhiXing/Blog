public class Log4NetProvider
{
    private IDictionary<string, ILogger> _loggers = new Dictionary<string, ILogger>();
    
    public ILogger CreateLogger(string name)
    {
        if (_loggers.ContainsKey(name))
        {
            return _loggers[name];
        }
        else
        {
            ILogger logger = new Log4NetAdapter(name);
            _loggers.Add(name, logger);
            return logger;
        }
    }

    public void Dispose()
    {
        _loggers.Clear();
    }
}