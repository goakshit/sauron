## Overview
A simpl buy now, pay later implementation.

#### Build
```
git clone git@github.com:goakshitsauron.gi
cd sauron

-- Database
Create database 'sdb'
Run the sql commands in ./build/scripts/db/init.sql  
--

go install
```

#### Commands supported
```
sauron new user (args)['name', 'email', 'credit limit']
sauron new merchant (args)['name', 'email', 'discount percentage']
sauron new txn (args)['user name', 'merchant name', 'amount']
sauron update user (args)['user name', 'new credit limit']
sauron update merchant (args)['user name', 'new discount percentage']
sauron payback (args)['user name', 'amount']
sauron report total-dues
sauron report users-at-credit-limit
sauron report dues (args)['user name']
sauron report discount (args)['merchant name']
```

#### Tests
Mocks are added for service & db layer. </br>To exec: `go test ./...`
