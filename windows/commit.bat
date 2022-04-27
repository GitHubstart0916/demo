call base\set_path.bat
:: set source_path and target_path

cd source_path

git commit -m "%1"


cd target_path
git commit -m "%1"
