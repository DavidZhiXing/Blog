    /// <summary>
    /// 剧集封面转换器（降低分辨率）.
    /// </summary>
    public class EpisodeCoverConverter : IValueConverter
    {
        /// <inheritdoc/>
        public object Convert(object value, Type targetType, object parameter, string language)
        {
            if (value is string coverUrl)
            {
                var url = coverUrl + "@200w_180h_1c_100q.jpg";
                return new BitmapImage(new Uri(url));
            }

            return value;
        }

        /// <inheritdoc/>
        public object ConvertBack(object value, Type targetType, object parameter, string language) => throw new NotImplementedException();
    }