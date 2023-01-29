### A Good Example with Attribute
``` C#
using System;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.Filters;

namespace WebApp
{
    public class LoggingActionFilter : ActionFilterAttribute
    {
        public override void OnActionExecuting(ActionExecutingContext context)
        {
            Console.WriteLine($"Action '{context.ActionDescriptor.DisplayName}' is executing");
        }

        public override void OnActionExecuted(ActionExecutedContext context)
        {
            Console.WriteLine($"Action '{context.ActionDescriptor.DisplayName}' has executed");
        }
    }

    [LoggingActionFilter]
    public class HomeController : Controller
    {
        public IActionResult Index()
        {
            return View();
        }
    }
}
```

### Examples of Interceptor
- LoggingInterceptor
- CachingInterceptor
- PerformanceInterceptor
- AuthenticationInterceptor
- RetryInterceptor
- ExceptionInterceptor