git flow init -d
git flow feature start <feature_name>
git flow feature finish <feature_name>
git flow feature publish <feature_name>
git flow feature publish --push <feature_name>
git flow feature publish --sign <feature_name>

git add .
git commit -m "Add <feature_name>"
git reset --hard HEAD~1
git rebase -i <feature_name>

dotnet new console -o <feature_name>
dotnet add <feature_name>
dotnet restore <feature_name>
dotnet build <feature_name>
dotnet pack <feature_name>
dotnet publish <feature_name>

echo "git flow feature publish --push <feature_name>"
del <feature_name>
rar.exe a <feature_name>.rar <feature_name>
zip -r <feature_name>.zip <feature_name>
mkdir <feature_name>
cd  <feature_name>

git fetch --purge origin
