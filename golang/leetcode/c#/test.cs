public class test{
    public static int add(int a, int b)
    {
        return a + b;
    }

    public static int concurrent_add(int a, int b)
    {
        lock (typeof(test))
        {
            return a + b;
        }
    }

    public static List<int> GetShalowCopy(List<int> list)
    {
        return new List<int>(list);
    }

    [TestMethod]
    public void TestMethod1()
    {
        var list = new List<int> { 1, 2, 3 };
        var list2 = GetShalowCopy(list);
        list.Add(4);
        Assert.AreEqual(list2.Count, 3);
        Assert.AreEqual(ReferenceEquals(list, list2), false);
        Assert.AreEqual(list, list2);
    }
}