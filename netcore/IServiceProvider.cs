public interface IServiceProvider
{
    object GetService(Type serviceType);
}