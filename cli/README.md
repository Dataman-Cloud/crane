Crane Client
============

### login

```bash
./crane-cli.py login -e admin@admin.com -padminadmin
```

### about me

```bash
./crane-cli.py aboutme
```

### my groups

```
./crane-cli.py mygroups -u1
```

### create stack

#### sample 1

```
./crane-cli.py create_stack -G1 -ncli -f ../frontend/stack_samples/wordpress.json
```

#### sample 2

```
./crane-cli.py create_stack -G1 -ncli_mysql_web -f ../frontend/stack_samples/mysql_display.json
```

#### sample 3

```
./crane-cli.py create_stack -G1 -ncli_2048 -f ../frontend/stack_samples/2048.json
```

### list stack

```bash
./crane-cli.py list_stack
```

### list stack services

```bash
./crane-cli.py list_stack_services -n cli_created
```

### scale given service amounts

```bash
./crane-cli.py scale_service -n cli_created -s 6prqzc47jiohi6e4iwst9fwdw -a 2
```
