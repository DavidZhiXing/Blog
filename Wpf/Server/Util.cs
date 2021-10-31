public static class Util
{
    public static string GetAssemblyPath()
    {
        return Path.GetDirectoryName(Assembly.GetExecutingAssembly().Location);
    }

    public static T DeepCopy<T>(this T obj)
    {
        using (MemoryStream stream = new MemoryStream())
        {
            BinaryFormatter formatter = new BinaryFormatter();
            formatter.Serialize(stream, obj);
            stream.Position = 0;
            return (T)formatter.Deserialize(stream);
        }
    }

    public static void TestDeepCopy(){
        var a = new A();
        var b = a.DeepCopy();
        b.Name = "b";
        Console.WriteLine(a.Name);
    }

    public void Retry

    public class A
    {
        public string Name { get; set; }
    }

    public static long CaculateMethodCallTime(this Action action)
    {
        var stopwatch = new Stopwatch();
        stopwatch.Start();
        action();
        stopwatch.Stop();
        return stopwatch.ElapsedMilliseconds;
    }

    public static void LoadAssembly(string path)
    {
        var assembly = Assembly.LoadFile(path);
        var types = assembly.GetTypes();
        foreach (var type in types)
        {
            if (type.IsClass)
            {
                var obj = Activator.CreateInstance(type);
                Console.WriteLine(obj);
            }
        }
    }

    public static IList<Assembly> LoadPluginAssemblies(string path)
    {
        var assemblies = new List<Assembly>();
        var files = Directory.GetFiles(path, "*.dll");
        foreach (var file in files)
        {
            var assembly = Assembly.LoadFile(file);
            assemblies.Add(assembly);
        }
        return assemblies;
    }

    public IList<IPlugin> GetPlugins(IList<Assembly> assemblies)
    {
        var plugins = new List<IPlugin>();
        foreach (var assembly in assemblies)
        {
            var types = assembly.GetTypes();
            foreach (var type in types)
            {
                if (type.IsClass)
                {
                    var obj = Activator.CreateInstance(type);
                    if (obj is IPlugin)
                    {
                        plugins.Add(obj as IPlugin);
                    }
                }
            }
        }
        return plugins;
    }

    public void RunPlugins(IList<IPlugin> plugins)
    {
        foreach (var plugin in plugins)
        {
            try
            {
                plugin.Run();
            }
            catch (Exception ex)
            {
                Console.WriteLine(ex.Message);
            }
        }
    }

    public interface IPlugin
    {
        void Config();

        void Start();

        void Stop();
        void Run();
    }
}

[MyPlugin("test")]
public class Plugin : IPlugin
{
    public void Config()
    {
        Console.WriteLine("Config");
    }

    public void Start()
    {
        Console.WriteLine("Start");
    }

    public void Stop()
    {
        Console.WriteLine("Stop");
    }


    public void Run()
    {
        Console.WriteLine("Run");
    }
}

public class MyPluginAttribute : Attribute
{
    public string Name { get; set; }

    public MyPluginAttribute(string name)
    {
        Name = name;
    }
}

public class PluginHost{
    public PluginHost(IPlugin plugin)
    {
        plugin.Config();
        _plugin = plugin;
    }

}

public class PluginHostProvider{
    public PluginHostProvider(IList<IPlugin> plugins)
    {
        _plugins = plugins;
    }

    public PluginHost GetPluginHost(IPlugin plugin)
    {
        return new PluginHost(plugin);
    }

    public void Run()
    {
        foreach (var plugin in _plugins)
        {
            plugin.Run();
        }
    }

    private IList<IPlugin> _plugins;
}
