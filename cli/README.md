Rolex Client
============

### login

```bash
./rolex-cli.py login -e admin@admin.com -padminadmin
```

### about me

```bash
./rolex-cli.py aboutme
```

### my groups

```
./rolex-cli.py mygroups -u1
```

### create stack

#### sample 1

```bash
./rolex-cli.py create_stack -G1 -ncli_created -f samples/test.json
```

#### sample 2

```bash
./rolex-cli.py create_stack -G1 -ncli2_created -f samples/double_services.json
```

### list stack

```bash
./rolex-cli.py list_stack
```

### list stack services

```bash
./rolex-cli.py list_stack_services -n cli_created
```

