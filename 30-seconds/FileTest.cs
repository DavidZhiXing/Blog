[TestMethod]
public class FileTest
{
    [Fact]
    public void FileExists()
    {
        // arrange
        string fileName = "test.txt";
        string fileContent = "Hello World!";
        File.WriteAllText(fileName, fileContent);

        // act
        bool exists = File.Exists(fileName);

        // assert
        Assert.True(exists);
    }

    [Fact]
    public void FileDoesNotExist()
    {
        // arrange
        string fileName = "test.txt";

        // act
        bool exists = File.Exists(fileName);

        // assert
        Assert.False(exists);
    }

    [Fact]
    public void FileRead()
    {
        // arrange
        string fileName = "test.txt";
        string fileContent = "Hello World!";

        File.WriteAllText(fileName, fileContent);
        // act
        string content = File.ReadAllText(fileName);

        // assert
        Assert.Equal(fileContent, content);
    }

    [Fact]
    public void FileWrite()
    {
        // arrange
        string fileName = "test.txt";
        string fileContent = "Hello World!";

        // act
        File.WriteAllText(fileName, fileContent);
        string content = File.ReadAllText(fileName);

        // assert
        Assert.Equal(fileContent, content);
    }
}

// test if file is readable
// test if file is writable
// test if file is deletable
