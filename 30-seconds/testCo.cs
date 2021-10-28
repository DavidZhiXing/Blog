public class TestCo
{
    public static void Main()
    {
        var co = new Co();
        co.Start(() =>
        {
            Console.WriteLine("Hello");
            co.Stop();
        });
    }

    public static void readxml(string xml)
    {
        var doc = new XmlDocument();
        doc.LoadXml(xml);
        var root = doc.DocumentElement;
        var nodes = root.SelectNodes("/root/node");
        foreach (XmlNode node in nodes)
        {
            var name = node.Attributes["name"].Value;
            var value = node.Attributes["value"].Value;
            Console.WriteLine("{0} = {1}", name, value);
        }
    }

    public static void ObjectToXmlFile(object obj, string fileName)
    {
        var doc = new XmlDocument();
        var root = doc.CreateElement("root");
        doc.AppendChild(root);
        var type = obj.GetType();
        var properties = type.GetProperties();
        foreach (var property in properties)
        {
            var node = doc.CreateElement("node");
            node.SetAttribute("name", property.Name);
            node.SetAttribute("value", property.GetValue(obj).ToString());
            root.AppendChild(node);
        }
        doc.Save(fileName);
    }

    // find all files in a directory
    public static IEnumerable<string> FindFiles(string path)
    {
        var files = Directory.GetFiles(path);
        foreach (var file in files)
        {
            yield return file;
        }
        var dirs = Directory.GetDirectories(path);
        foreach (var dir in dirs)
        {
            foreach (var file in FindFiles(dir))
            {
                yield return file;
            }
        }
    }

    /// create a tcp server
    public static void TcpServer()
    {
        var server = new TcpListener(IPAddress.Any, 80);
        server.Start();
        while (true)
        {
            var client = server.AcceptTcpClient();
            var stream = client.GetStream();
            var buffer = new byte[1024];
            var bytesRead = stream.Read(buffer, 0, 1024);
            var data = Encoding.ASCII.GetString(buffer, 0, bytesRead);
            Console.WriteLine(data);
            var response = "HTTP/1.1 200 OK\r\n" +
                           "Content-Type: text/html\r\n" +
                           "Content-Length: 11\r\n\r\n" +
                           "Hello World";
            var responseBytes = Encoding.ASCII.GetBytes(response);
            stream.Write(responseBytes, 0, responseBytes.Length);
        }
    }

    /// create a tcp server via async
    public static void TcpServerAsync()
    {
        var server = new TcpListener(IPAddress.Any, 80);
        server.Start();
        while (true)
        {
            server.BeginAcceptTcpClient(AcceptTcpClient, server);
        }
    }

    /// create class to handle tcp client   
    public static void AcceptTcpClient(IAsyncResult ar)
    {
        var server = ar.AsyncState as TcpListener;
        var client = server.EndAcceptTcpClient(ar);
        var stream = client.GetStream();
        var buffer = new byte[1024];
        var bytesRead = stream.Read(buffer, 0, 1024);
        var data = Encoding.ASCII.GetString(buffer, 0, bytesRead);
        Console.WriteLine(data);
        var response = "HTTP/1.1 200 OK\r\n" +
                       "Content-Type: text/html\r\n" +
                       "Content-Length: 11\r\n\r\n" +
                       "Hello World";
        var responseBytes = Encoding.ASCII.GetBytes(response);
        stream.Write(responseBytes, 0, responseBytes.Length);
        client.Close();
    }


}