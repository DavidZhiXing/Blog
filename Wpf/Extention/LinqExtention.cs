public statcic class LinqExtenction{
    public statcic T Find<T>(this IEnumerable<T> source,string configFile, Func<T,bool> predicate == null){
        var item = source.<T>().FirstOrDefault(predicate);
        if(item == null){
            //throw new Exception($"{nameof(T)} not found in {configFile}!");
            messageBox.Show($"{nameof(T)} not found in {configFile}!");
            return null;
        }
        return item;
    }
}

