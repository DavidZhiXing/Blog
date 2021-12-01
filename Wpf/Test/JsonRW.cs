public patial Utill{

    public static string ToJson(this object obj)
    {
        return JsonConvert.SerializeObject(obj);
    }

    public static T FromJson<T>(this string json)
    {
        return JsonConvert.DeserializeObject<T>(json);
    }

    public static double[] DeepCopyArray(this double[] source)
    {
        double[] result = new double[source.Length];
        Array.Copy(source, result, source.Length);
        return result;
    }
    
}