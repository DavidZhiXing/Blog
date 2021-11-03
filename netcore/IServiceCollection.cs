public interface IServiceCollection
{
    void AddSingleton<TService, TImplementation>()
        where TService : class
        where TImplementation : class, TService;

    void AddSingleton<TService>()
        where TService : class;

    void AddSingleton<TService>(TService instance)
        where TService : class;

    void AddScoped<TService, TImplementation>()
        where TService : class
        where TImplementation : class, TService;

    void AddScoped<TService>()
        where TService : class;

    void AddScoped<TService>(TService instance)
        where TService : class;

    void AddTransient<TService, TImplementation>()
        where TService : class
        where TImplementation : class, TService;

    void AddTransient<TService>()
        where TService : class;

    void AddTransient<TService>(TService instance)
        where TService : class;

    void AddTransient<TService, TImplementation>(Func<IServiceProvider, TImplementation> factory)
        where TService : class
        where TImplementation : class, TService;

}