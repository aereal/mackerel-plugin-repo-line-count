# mackerel-plugin-repo-line-count

## Usage

```sh
mackerel-plugin-repo-line-count --name METRIC_REPO_NAME --path REPO_PATH --pattern SEARCH_PATTERN
```

### Example

```
$ mackerel-plugin-repo-line-count --name dotfiles --path /path/to/aereal/dotfiles --pattern 'TODO'
repo_lines.TODO..vim/colors/xoria256.vim        1       1448017735
repo_lines.TODO..zsh.d/functions/_pkgutil       3       1448017735
repo_lines.TODO..zsh.d/functions/_prove 1       1448017735
```

![image](https://cloud.githubusercontent.com/assets/87649/11298461/721bf9c6-8fc2-11e5-956c-6443c5e76972.png)
