- SetUp()
- TestXXXX
- TearDown()

### example:

**SetUp**

``` cpp
   void SetUp()
   {
      DirManager::SetTempDir("/tmp/sequence-test-dir");
      mDirManager = new DirManager;

      mSequence = new Sequence(mDirManager, floatSample);

      mMemorySequence.clear();
   }
```

**TearDown**

``` cpp
   void TearDown()
   {
      delete mSequence;
      delete mDirManager;
      mMemorySequence.clear();
   }
```

**Tester**

``` cpp
int main()
{
   SequenceTest tester;

   tester.SetUp();
   tester.TestReferencing();
   tester.TearDown();

   tester.SetUp();
   tester.TestSetGarbageInput();
   tester.TearDown();

   return 0;
}
```

