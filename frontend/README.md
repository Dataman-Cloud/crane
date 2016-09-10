# frontend

## how to start the frontend independently, given that these is a running backend?

### Option 1: nginx in docker

```bash
CRANE_IP=X.X.X.X ./bin/start.sh
```

then, visit http://localhost

### Option 2: gulp serve

1. install the package dependence
  ```bash
  npm install && bower install
  ```
2. edit the config file `conf.js`, make `SAMPLES_URL = '/stack_samples/'`
3. start the serve
  ```bash
  gulp serve
  ```
4. visit http://localhost:5000
