public sealed class PerformanceCounter : Component, ISupportInitialize
{
    public PerformanceCounter()
    {
        InitializeComponent();
    }

    public void BeginInit()
    {
        throw new NotImplementedException();
    }

    public void EndInit()
    {
        throw new NotImplementedException();
    }

    public PerformanceCounter(string categoryName, string counterName)
    {
    }

    public PerformanceCounter(string categoryName, string counterName, string instanceName)
    {
    }

    public PerformanceCounter(string categoryName, string counterName, string instanceName, bool readOnly)
    {
    }

    public PerformanceCounter(string categoryName, string counterName, string instanceName, string machineName)
    {
    }

    public string CategoryName
    {
        get
        {
            return null;
        }
    }

    public string CounterName
    {
        get
        {
            return null;
        }
    }

    public string InstanceName
    {
        get
        {
            return null;
        }
    }

    public string MachineName
    {
        get
        {
            return null;
        }
    }

    public bool ReadOnly
    {
        get
        {
            return false;
        }
    }

    public PerformanceCounterType CounterType
    {
        get
        {
            return PerformanceCounterType.NumberOfItems32;
        }
    }

    pub PerformanceCounterInstanceLifetime InstanceLifetime
    {
        get
        {
            return PerformanceCounterInstanceLifetime.Global;
        }
    }

    public long RawValue
    {
        get
        {
            return 0;
        }
        set
        {
        }
    }

    public void Close()
    {
    }

    public static void CloseSharedResources()
    {
    }

    public long Increment()
    {

    }
}