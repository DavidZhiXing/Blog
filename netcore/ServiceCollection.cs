public class ServiceCollection : IserviceCollection
{
    public ServiceCollection()
    {
        services = new List<ServiceDescriptor>();
    }

    public void Add(ServiceDescriptor descriptor)
    {
        services.Add(descriptor);
    }

    public void AddSingleton<T>(T implementation)
    {
        Add(new ServiceDescriptor(typeof(T), implementation, ServiceLifetime.Singleton));
    }

    public void AddSingleton<T>(Func<IServiceProvider, T> implementationFactory)
    {
        Add(new ServiceDescriptor(typeof(T), implementationFactory, ServiceLifetime.Singleton));
    }

    public void AddScoped<T>(T implementation)
    {
        Add(new ServiceDescriptor(typeof(T), implementation, ServiceLifetime.Scoped));
    }

    public void AddScoped<T>(Func<IServiceProvider, T> implementationFactory)
    {
        Add(new ServiceDescriptor(typeof(T), implementationFactory, ServiceLifetime.Scoped));
    }

    public void AddTransient<T>(T implementation)
    {
        Add(new ServiceDescriptor(typeof(T), implementation, ServiceLifetime.Transient));
    }

    public void AddTransient<T>(Func<IServiceProvider, T> implementationFactory)
    {
        Add(new ServiceDescriptor(typeof(T), implementationFactory, ServiceLifetime.Transient));
    }

    public void Add(Type serviceType, Type implementationType)
    {
        Add(new ServiceDescriptor(serviceType, implementationType, ServiceLifetime.Transient));
    }

    public void Add(Type serviceType, Type implementationType, ServiceLifetime lifetime)
    {
        Add(new ServiceDescriptor(serviceType, implementationType, lifetime));
    }

    public void Add(Type serviceType, Func<IServiceProvider, object> implementationFactory)
    {
        Add(new ServiceDescriptor(serviceType, implementationFactory, ServiceLifetime.Transient));
    }

    public void Add(Type serviceType, Func<IServiceProvider, object> implementationFactory, ServiceLifetime lifetime)
    {
        Add(new ServiceDescriptor(serviceType, implementationFactory, lifetime));
    }

    private List<ServiceDescriptor> services;
}