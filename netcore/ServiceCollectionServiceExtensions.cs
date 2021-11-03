public static class ServiceCollectionServiceExtensions
{
    public static IServiceCollection AddServiceCollection(this IServiceCollection services)
    {
        return services;
    }

    public static IServiceCollection AddSingleton(this IServiceCollection services, Type serviceType, Type implementationType)
    {
        services.AddSingleton(serviceType, implementationType);
        return services;
    }
}