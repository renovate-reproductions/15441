# Renovate gomod manager issues

This repo demonstrates an issue with the gomod manager of Renovate.
The problem is that `go mod tidy` does not work when set as `postUpdateOptions` in the renovate.json (see [log.txt](log.txt)).

```json
  "postUpdateOptions": [
    "gomodTidy"
  ]
```

The error message is not reproducible in any of the following ways to execute `go get -d -t ./... && go mod tidy`.

1. Run it locally in the repo by checking out the branch created by renovate and running `go get -d -t ./... && go mod tidy` -> no issues
2. Running the following docker command also works (as shown in the renovate bot's error message, are any environment settings overwritten?)

   ```shell
   docker run --rm --name=renovate_go --label=renovate_child -v $(pwd):"/mnt/renovate/gh/brumhard/renovate-testing2" -w "/mnt/renovate/gh/brumhard/renovate-testing2" docker.io/renovate/go:1.18.1 bash -l -c "go get -d -t ./... && go mod tidy && go mod tidy"
   ```

3. Setting `go mod tidy` as a `postUpgradeTasks` also works when using the self hosted runner (see [config.js](config.js)). To confirm this you can remove the `postUpdateOptions` in the [renovate.json](renovate.json) and run the self hosted runner again.