@REM npm install @openzeppelin/contracts
@REM npm install @truffle/hdwallet-provider
@REM npm install truffle-privatekey-provider
@REM npm install dotenv

@REM truffle compile

@REM Collect imports for validate on etherscan.io by single file:
@REM npm run build-contracts

@REM Run new console window and network
start cmd.exe @cmd /k "ganache-cli --networkId 4447"

@REM Deploy and Run local tests:

call truffle migrate --f 1 --to 1 --network development

echo Deployment successful

@REM This project tests:
call truffle test ./test/testing.js --compile-none --migrations_directory migrations_null

call echo Tests finished

pause
