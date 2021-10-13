    public class ResourceToolkit : IResourceToolkit
    {
        private readonly Application _app;

        /// <summary>
        /// Initializes a new instance of the <see cref="ResourceToolkit"/> class.
        /// </summary>
        public ResourceToolkit()
        {
            _app = Application.Current;
        }

        /// <inheritdoc/>
        public string GetLocaleString(LanguageNames languageName)
            => ResourceLoader.GetForCurrentView().GetString(languageName.ToString());

        /// <inheritdoc/>
        public T GetResource<T>(string resourceName)
            => (T)_app.Resources[resourceName];
    }