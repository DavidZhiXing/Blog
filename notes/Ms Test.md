`DataRowAttribute`

``` c#
[TestMethod]
  [DataRow(1,2)]
  [DataRow(2,2)]
  public void TestSomeNumbers (int x, int y)
  {
    Assert.AreEqual(x,y)
  }
```

Gives output:
Test Failed - `TestSomeNumbers (1,2)`
Test Passed - `TestSomeNumbers (2,2)`

You can also specify the test display name for each data row:

``` c#
[TestMethod]
  [DataRow(1,2, DisplayName ="Sequential numbers")]
  [DataRow(2,2, DisplayName ="Equal numbers")]
  public void TestSomeNumbers (int x, int y)
  {
    Assert.AreEqual(x,y)       
  }
```

Gives output:
Test Failed - Sequential numbers
Test Passed - Equal numbers

`Assert.AreEqual`

``` C#
public static void AreEqual (double expected, double actual, double delta, string message, params object[] parameters);
```

