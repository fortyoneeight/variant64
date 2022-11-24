# Variant 64

## Setup
- Setup repo after clone
```
make install
```

## Common Commands
- Get list of commands
```
make help
```

- Run all components in containers
```
make run
```

- Run backend component in container
```
make run-server
```

- Run frontend components in containers
```
make run-proxy
make run-client
```

- Run frontend components manually
```
cd frontend/
npx lerna run start
```

- Stop components in containers
```
make stop
```

- Run tests
```
make test
```
