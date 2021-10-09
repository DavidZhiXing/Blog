    public static class UtilsTests
    {
        [SkippableFact]
        public static void TestIsWinPEHost()
        {
            Skip.IfNot(Platform.IsWindows);
            Assert.False(Utils.IsWinPEHost());
        }

        
        [Fact]
        public static void TestHistoryStack()
        {
            var historyStack = new HistoryStack<string>(20);
            Assert.Equal(0, historyStack.UndoCount);
            Assert.Equal(0, historyStack.RedoCount);

            historyStack.Push("first item");
            historyStack.Push("second item");
            Assert.Equal(2, historyStack.UndoCount);
            Assert.Equal(0, historyStack.RedoCount);

            Assert.Equal("second item", historyStack.Undo("second item"));
            Assert.Equal("first item", historyStack.Undo("first item"));
            Assert.Equal(0, historyStack.UndoCount);
            Assert.Equal(2, historyStack.RedoCount);

            Assert.Equal("first item", historyStack.Redo("first item"));
            Assert.Equal(1, historyStack.UndoCount);
            Assert.Equal(1, historyStack.RedoCount);

            // Pushing a new item should invalidate the RedoCount
            historyStack.Push("third item");
            Assert.Equal(2, historyStack.UndoCount);
            Assert.Equal(0, historyStack.RedoCount);

            // Check for the correct exception when the Redo/Undo stack is empty.
            Assert.Throws<InvalidOperationException>(() => historyStack.Redo("bar"));
            historyStack.Undo("third item");
            historyStack.Undo("first item");
            Assert.Equal(0, historyStack.UndoCount);
            Assert.Throws<InvalidOperationException>(() => historyStack.Undo("foo"));
        }

        [Fact]
        public static void TestConvertToJsonBasic()
        {
            var context = new JsonObject.ConvertToJsonContext(maxDepth: 1, enumsAsStrings: false, compressOutput: true);
            string expected = "{\"name\":\"req\",\"type\":\"http\"}";
            OrderedDictionary hash = new OrderedDictionary {
                {"name", "req"},
                {"type", "http"}
            };
            string json = JsonObject.ConvertToJson(hash, in context);
            Assert.Equal(expected, json);

            hash.Add("self", hash);
            json = JsonObject.ConvertToJson(hash, context);
            expected = "{\"name\":\"req\",\"type\":\"http\",\"self\":{\"name\":\"req\",\"type\":\"http\",\"self\":\"System.Collections.Specialized.OrderedDictionary\"}}";
            Assert.Equal(expected, json);
        }
    }