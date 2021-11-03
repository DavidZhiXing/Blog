public class ServiceProvider:IServiceProvider
{
    private readonly IServiceCollection _services;
    public ServiceProvider(IServiceCollection services)
    {
        _services = services;
    }
    public object GetService(Type serviceType)
    {
        var service = _services.FirstOrDefault(d => d.ServiceType == serviceType);
        if (service == null)
        {
            return null;
        }
        return service.ImplementationInstance;
    }
}