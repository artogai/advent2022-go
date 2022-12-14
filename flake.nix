{
  description = "Go env flake";

  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils, ... } @ inputs:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        devShell = pkgs.mkShell {
          nativeBuildInputs = with pkgs; [            
            go
            gopls
            delve
            go-outline
            gomodifytags
            impl
            gotests
            golangci-lint
            go-tools #staticcheck
          ];
          
          GOPATH = "/home/artem/Projects/go";
        };
      });
}
