public class PlinqTest
{
    public static void TestPlinq(){
        var numbers = Enumerable.Range(0, 100);
        var parallelResult = numbers.AsParallel().AsOrdered()
            .Where(i => i % 2 == 0).AsSequential();
        foreach (int i in parallelResult.Take(10))
        {
            Console.WriteLine(i);
        }
    }

    public static void TestParallelForEach(){
        var numbers = Enumerable.Range(0, 100);
        Parallel.ForEach(numbers, (number) =>
        {
            Console.WriteLine(number);
        });
    }

    public static void TestParallelFor(){
        var numbers = Enumerable.Range(0, 100);
        Parallel.For(0, 100, (int i) =>
        {
            Console.WriteLine(i);
        });
    }

    public static void TestParallelInvoke(){
        Parallel.Invoke(() =>
        {
            Console.WriteLine("1");
        }, () =>
        {
            Console.WriteLine("2");
        });

    }

    public static void TestParallelWithLock(){
        var numbers = Enumerable.Range(0, 100);
        Parallel.ForEach(numbers, () =>
        {
            int i = 0;
            lock (i)
            {
                Console.WriteLine(i);
            }
        });
    }

    public static void TestPlinqWithAsyncSelectFilter(){
        var numbers = Enumerable.Range(0, 100);
        List<int> result = Task.WhenAll(number.Select(async number =>
        {
            await Task.Delay(100);
            return number;
        }).Where(number => number % 2 == 0)).Result.ToList();
    }

    public static void TestDepthClone(){
        using MemoryStream ms = new MemoryStream();
        BinaryFormatter bf = new BinaryFormatter();
        bf.Serialize(ms, new object());
        ms.Position = 0;
        object obj = bf.Deserialize(ms);
    }

    [TestMethod]
    public static void TestCompareAndExchange(){

        // given
        int i = 0;

        // when
        int j = Interlocked.CompareExchange(ref i, 1, 0);

        // then
        Assert.AreEqual(0, j);
        Assert.AreEqual(1, i);
        
    }
}
