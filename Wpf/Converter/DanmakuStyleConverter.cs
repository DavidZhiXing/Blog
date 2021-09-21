    /// <summary>
    /// 弹幕样式转换器.
    /// </summary>
    public class DanmakuStyleConverter : IValueConverter
    {
        /// <inheritdoc/>
        public object Convert(object value, Type targetType, object parameter, string language)
        {
            var result = string.Empty;
            var resToolkit = ServiceLocator.Instance.GetService<IResourceToolkit>();
            switch ((DanmakuStyle)value)
            {
                case DanmakuStyle.Stroke:
                    result = resToolkit.GetLocaleString(Models.Enums.LanguageNames.Stroke);
                    break;
                case DanmakuStyle.NoStroke:
                    result = resToolkit.GetLocaleString(Models.Enums.LanguageNames.NoStroke);
                    break;
                case DanmakuStyle.Shadow:
                    result = resToolkit.GetLocaleString(Models.Enums.LanguageNames.Shadow);
                    break;
            }

            return result;
        }

        /// <inheritdoc/>
        public object ConvertBack(object value, Type targetType, object parameter, string language) => throw new NotImplementedException();
    }