## DIRENV

A Cli tool for loading env variables. Langauge independent!

```bash
#homebrew
brew install direnv
#nixos
nix-shell -p direnv

# then add a shell hook for direnv
# for fish
echo "direnv hook fish | source" >> ~/.config/fish/config.fish
```

Write all your env variables into `.envrc`

```.envrc
export ADDR=":3000"
```
