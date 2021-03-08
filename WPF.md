**Mvvm**

示例程序

![Figure 1 Workspaces screenshot](https://docs.microsoft.com/en-us/archive/msdn-magazine/2009/february/images/dd419663.fig01_l.gif(en-us,msdn.10).gif)

![New Customer Data Entry Form](https://docs.microsoft.com/en-us/archive/msdn-magazine/2009/february/images/dd419663.fig02_l.gif(en-us,msdn.10).gif)

ViewModelBase 复用性非常高，基本上任何ViewModel都要实现

``` C#
public event PropertyChangedEventHandler PropertyChanged;
protected virtual void OnPropertyChanged(string propertyName)
{
    this.VerifyPropertyName(propertyName);
    PropertyChangedEventHandler handler = this.PropertyChanged;
    if (handler != null)
    {
        var e = new PropertyChangedEventArgs(propertyName); 
        handler(this, e);
    }
}
```

RelayCommand 比DeleteGateCommand更加高效

``` C#
public class RelayCommand : ICommand
{
    #region Fields 
    readonly Action<object> _execute;
    readonly Predicate<object> _canExecute;
    #endregion // Fields 
    #region Constructors 
    public RelayCommand(Action<object> execute) : this(execute, null) { }
    public RelayCommand(Action<object> execute, Predicate<object> canExecute)
    {
        if (execute == null)
            throw new ArgumentNullException("execute");
        _execute = execute; _canExecute = canExecute;
    }
    #endregion // Constructors 
    #region ICommand Members 
    [DebuggerStepThrough]
    public bool CanExecute(object parameter)
    {
        return _canExecute == null ? true : _canExecute(parameter);
    }
    public event EventHandler CanExecuteChanged
    {
        add { CommandManager.RequerySuggested += value; }
        remove { CommandManager.RequerySuggested -= value; }
    }
    public void Execute(object parameter) { _execute(parameter); }
    #endregion // ICommand Members 
}
```



``` C#
RelayCommand _saveCommand; public ICommand SaveCommand
{
    get
    {
        if (_saveCommand == null) {
            _saveCommand = new RelayCommand(param => this.Save(), 
                param => this.CanSave);
        }
        return _saveCommand;
    }
}
```

![Inheritance Hierarchy](https://docs.microsoft.com/en-us/archive/msdn-magazine/2009/february/images/dd419663.fig04_l.gif(en-us,msdn.10).gif)

Command ViewModel

``` c#
public class CommandViewModel : ViewModelBase
{
    public CommandViewModel(string displayName, ICommand command)
    {
        if (command == null)
            throw new ArgumentNullException("command");
        base.DisplayName = displayName;
        this.Command = command;
    }
    public ICommand Command { get; private set; }
}
```

Render Command List

``` xml
<!-- In MainWindowResources.xaml -->
<!-- This template explains how to render the list of commands 
on the left side in the main window (the 'Control Panel' area). -->
<DataTemplate x:Key="CommandsTemplate">
    <ItemsControl ItemsSource="{Binding Path=Commands}">
        <ItemsControl.ItemTemplate>
            <DataTemplate>
                <TextBlock Margin="2,6"> 
                    <Hyperlink Command="{Binding Path=Command}"> 
                    <TextBlock Text="{Binding Path=DisplayName}" /> 
                    </Hyperlink> 
                </TextBlock>
            </DataTemplate>
        </ItemsControl.ItemTemplate>
    </ItemsControl>
</DataTemplate>
```

MainViewModel

``` C#
// In App.xaml.cs 
protected override void OnStartup(StartupEventArgs e)
{
    base.OnStartup(e); MainWindow window = new MainWindow();
    // Create the ViewModel to which 
    // the main window binds. 
    string path = "Data/customers.xml";
    var viewModel = new MainWindowViewModel(path);
    // When the ViewModel asks to be closed, 
    // close the window. 
    viewModel.RequestClose += delegate { window.Close(); };
    // Allow all controls in the window to 
    // bind to the ViewModel by setting the 
    // DataContext, which propagates down 
    // the element tree. 
    window.DataContext = viewModel;
    window.Show();
}
```

``` xml
<!-- In MainWindow.xaml -->
<Menu>
    <MenuItem Header="_File">
        <MenuItem Header="_Exit" 
                    Command="{Binding Path=CloseCommand}" />
    </MenuItem>
    <MenuItem Header="_Edit" />
    <MenuItem Header="_Options" />
    <MenuItem Header="_Help" />
</Menu>
```

DataTemplate

``` xml
<DataTemplate x:Key="ClosableTabItemTemplate">
    <DockPanel Width="120">
        <Button Command="{Binding Path=CloseCommand}" 
                Content="X" 
                DockPanel.Dock="Right" 
                Width="16" 
                Height="16" />
        <ContentPresenter 
            Content="{Binding Path=DisplayName}" />
    </DockPanel>
</DataTemplate>
```



Observable Collection-->Removing Workspace from the UI

``` c#
// In MainWindowViewModel.cs 
ObservableCollection<WorkspaceViewModel> _workspaces;
public ObservableCollection<WorkspaceViewModel> Workspaces
{
    get
    {
        if (_workspaces == null)
        {
            _workspaces = new ObservableCollection<WorkspaceViewModel>();
            _workspaces.CollectionChanged += this.OnWorkspacesChanged;
        }
        return _workspaces;
    }
}
void OnWorkspacesChanged(object sender, NotifyCollectionChangedEventArgs e)
{
    if (e.NewItems != null && e.NewItems.Count != 0)
        foreach (WorkspaceViewModel workspace in e.NewItems)
            workspace.RequestClose += this.OnWorkspaceRequestClose;
    if (e.OldItems != null && e.OldItems.Count != 0)
        foreach (WorkspaceViewModel workspace in e.OldItems)
            workspace.RequestClose -= this.OnWorkspaceRequestClose;
}
void OnWorkspaceRequestClose(object sender, EventArgs e)
{
    this.Workspaces.Remove(sender as WorkspaceViewModel);
}
```

Test Method

``` C#
// In MainWindowViewModelTests.cs 
[TestMethod]
public void TestCloseAllCustomersWorkspace()
{
    // Create the MainWindowViewModel, but not the MainWindow. 
    MainWindowViewModel target = new MainWindowViewModel(Constants.CUSTOMER_DATA_FILE);
    Assert.AreEqual(0, target.Workspaces.Count, "Workspaces isn't empty.");
    // Find the command that opens the "All Customers" workspace. 
    CommandViewModel commandVM = target.Commands.First(cvm => cvm.DisplayName == "View all customers");
    // Open the "All Customers" workspace. 
    commandVM.Command.Execute(null); Assert.AreEqual(1, target.Workspaces.Count, "Did not create viewmodel.");
    // Ensure the correct type of workspace was created. 
    var allCustomersVM = target.Workspaces[0] as AllCustomersViewModel;
    Assert.IsNotNull(allCustomersVM, "Wrong viewmodel type created.");
    // Tell the "All Customers" workspace to close. 
    allCustomersVM.CloseCommand.Execute(null);
    Assert.AreEqual(0, target.Workspaces.Count, "Did not close viewmodel.");
}
```

Applying a view to viewmodel, 通过DateTemplate来为ViewModel绑定一个View

``` xaml
<!-- This resource dictionary is used by the MainWindow. -->
<ResourceDictionary xmlns="http://schemas.microsoft.com/winfx/2006/xaml/presentation" 
                    xmlns:x="http://schemas.microsoft.com/winfx/2006/xaml" 
                    xmlns:vm="clr-namespace:DemoApp.ViewModel" 
                    xmlns:vw="clr-namespace:DemoApp.View" >
    <!-- This template applies an AllCustomersView to an instance 
    of the AllCustomersViewModel class shown in the main window. -->
    <DataTemplate DataType="{x:Type vm:AllCustomersViewModel}">
        <vw:AllCustomersView />
    </DataTemplate>
    <!-- This template applies a CustomerView to an instance of 
    the CustomerViewModel class shown in the main window. -->
    <DataTemplate DataType="{x:Type vm:CustomerViewModel}">
        <vw:CustomerView />
    </DataTemplate>
    <!-- Other resources omitted for clarity... -->
</ResourceDictionary>
```

