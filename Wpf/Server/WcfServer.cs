// caculate the time of last call method
public class WcfServer
{
    public static void Main()
    {
        ServiceHost host = new ServiceHost(typeof(Calculator));
        host.Open();
        Console.WriteLine("Service started at {0}", DateTime.Now);
        Console.WriteLine("Press <ENTER> to terminate service.");
        Console.ReadLine();
        host.Close();
    }

    // caculate the time of method excute
    public static void CaculateTime(string methodName, string serviceName)
    {
        Stopwatch stopwatch = new Stopwatch();
        stopwatch.Start();
        // call the method
        // ...
        InvokeMethod(methodName, serviceName);
        stopwatch.Stop();
        Console.WriteLine("{0} {1} {2}", serviceName, methodName, stopwatch.Elapsed);
    }

    public static void AddTwoNumbers()
    {
        CaculateTime("AddTwoNumbers", "WcfServer");
        // call CaculateTime 10 times
        for (int i = 0; i < 10; i++)
        {
            CaculateTime("AddTwoNumbers", "WcfServer");
        }
    }

        public static void AddTwoNumbers(int a, int b)
    {
        Console.WriteLine("{0} {1}", a, b);
    }

    public static void InvokeMethod(string methodName, string serviceName)
    {
        Type type = Type.GetType(serviceName);
        MethodInfo method = type.GetMethod(methodName);
        method.Invoke(null, null);
    }

}