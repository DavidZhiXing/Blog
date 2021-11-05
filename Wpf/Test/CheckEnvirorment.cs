public class CheckEnvirorment
{
    public static void CheckDotNetVersion()
    {
        if (Environment.Version.Major < 4)
        {
            throw new Exception("This sample requires .NET Framework 4.0 or higher.");
        }
    }

    public static void InstallDotNetRumtime()
    {
        if (Environment.Version.Major < 4)
        {
            // Download and install .NET Framework 4

            // Download .NET Framework 4 from http://www.microsoft.com/download/en/details.aspx?id=17851
            // Install .NET Framework 4
            var path = @"C:\Program Files (x86)\Reference Assemblies\Microsoft\Framework\v4.0";
            if (!Directory.Exists(path))
            {
                //下载并安装.net runtime 4.7.2

            }
        }
    }

    public static void InstallRuntimeBackward(){

    }

    public void Config(){
        services.AddLogging(logging =>
        {
            logging.AddConsole();
            logging.AddDebug();
            logging.AddLog4Net();
        });
    }

}
  