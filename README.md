# mackerel-plugin-repo-line-count

## Usage

```sh
mackerel-plugin-repo-line-count --name METRIC_REPO_NAME --path REPO_PATH --pattern SEARCH_PATTERN
```

### Example

```
$ mackerel-plugin-repo-line-count --name mackerel-agent --path /path/to/mackerel-agent --pattern 'TODO'
repo_lines.TODO.command/command.go      1       1448017565
repo_lines.TODO.mackerel/host.go        1       1448017565
repo_lines.TODO.main.go 1       1448017565
repo_lines.TODO.metrics/windows/interface_test.go       1       1448017565
repo_lines.TODO.spec/freebsd/interface.go       1       1448017565
repo_lines.TODO.spec/netbsd/interface.go        1       1448017565
```
