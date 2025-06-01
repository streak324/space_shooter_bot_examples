@echo off
echo Building c-bot...
cd c-bot
call build.bat
cd ..

echo Building example-bot-go...
cd example-bot-go
call build.bat
cd ..

echo All builds complete.