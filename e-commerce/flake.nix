{
  description = "Go Development Environment";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    utils.url = "github:numtide/flake-utils"; #This is a helper library. Nix can run on Mac (Intel/M1) and Linux. This utility helps us write the config once and have it work on all those systems automatically.
  };

  outputs = { self, nixpkgs, utils }:
    utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
        go-migrate-pg = pkgs.go-migrate.overrideAttrs(oldAttrs: {
           tags = ["postgres"];
          });
      in
      {
        devShells.default = pkgs.mkShell {
  
          buildInputs = with pkgs; [
            go
            go-migrate-pg
            air
            gopls
          ];

          shellHook = ''
            echo "🐹 Gopher Environment Active"
            echo "Go version: $(go version)"
          '';
        };
      });
}