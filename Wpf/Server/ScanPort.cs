public class ScanPort
{
    public static void Main()
    {
        int port = int.Parse(Console.ReadLine());
        TcpClient client = new TcpClient();
        try
        {
            client.Connect("localhost", port);
            Console.WriteLine("Port {0} is open", port);
        }
        catch (Exception)
        {
            Console.WriteLine("Port {0} is closed", port);
        }
        finally
        {
            client.Close();
        }
    }

    public static void ScanPortRange()
    {
        int startPort = int.Parse(Console.ReadLine());
        int endPort = int.Parse(Console.ReadLine());
        TcpClient client = new TcpClient();
        for (int i = startPort; i <= endPort; i++)
        {
            try
            {
                client.Connect("localhost", i);
                Console.WriteLine("Port {0} is open", i);
            }
            catch (Exception)
            {
                Console.WriteLine("Port {0} is closed", i);
            }
        }
    }

    public static void ScanPortRangeWithThreads()
    {
        int startPort = int.Parse(Console.ReadLine());
        int endPort = int.Parse(Console.ReadLine());

        Thread[] threads = new Thread[endPort - startPort + 1];
        for (int i = 0; i < threads.Length; i++)
        {
            int currentPort = startPort + i;
            threads[i] = new Thread(() =>
            {
                TcpClient client = new TcpClient();
                try
                {
                    client.Connect("localhost", currentPort);
                    Console.WriteLine("Port {0} is open", currentPort);
                }
                catch (Exception)
                {
                    Console.WriteLine("Port {0} is closed", currentPort);
                }
                finally
                {
                    client.Close();
                }
            });
        }
        
        for (int i = 0; i < threads.Length; i++)
        {
            threads[i].Start();
        }
    }

    public static async void ScanPortRangeWithAsync()
    {
        int startPort = int.Parse(Console.ReadLine());
        int endPort = int.Parse(Console.ReadLine());

        for (int i = startPort; i <= endPort; i++)
        {
            TcpClient client = new TcpClient();
            try
            {
                await client.ConnectAsync("localhost", i);
                Console.WriteLine("Port {0} is open", i);
            }
            catch (Exception)
            {
                Console.WriteLine("Port {0} is closed", i);
            }
            finally
            {
                client.Close();
            }
        }
    }

    public async void ScanPortRangeByArp()
    {
        int startPort = int.Parse(Console.ReadLine());
        int endPort = int.Parse(Console.ReadLine());

        for (int i = startPort; i <= endPort; i++)
        {
            try
            {
                await PingAsync("localhost", i);
                Console.WriteLine("Port {0} is open", i);
            }
            catch (Exception)
            {
                Console.WriteLine("Port {0} is closed", i);
            }
        }
    }

    public static async Task PingAsync(string host, int port)
    {
        using (TcpClient client = new TcpClient())
        {
            await client.ConnectAsync(host, port);
        }
    }

    public async IEnumerable<string> ExportProperties()
    {
        var properties = typeof(ScanPort).GetProperties();
        foreach (var property in properties)
        {
            yield return property.Name;
        }
    }

    public async IEnumerable<string, string> ExportPropertiesWithValues()
    {
        var properties = typeof(ScanPort).GetProperties();
        foreach (var property in properties)
        {
            yield return property.Name;
            yield return property.GetValue(this).ToString();
        }
    }

    public async IEnumerable<Tuple<string, string, string>> ExportPropertiesWithValuesAndComments(){
        var properties = typeof(ScanPort).GetProperties();
        foreach (var property in properties)
        {
            yield return Tuple.Create(property.Name, property.GetValue(this).ToString(), property.GetCustomAttribute<DescriptionAttribute>().Description);
        }
        
    }

    public void LinqToXml()
    {
        var properties = typeof(ScanPort).GetProperties();
        XElement root = new XElement("ScanPort");
        foreach (var property in properties)
        {
            root.Add(new XElement(property.Name, property.GetValue(this)));
        }
        root.Save("ScanPort.xml");
    }

}