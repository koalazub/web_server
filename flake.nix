{
  description = "Dev env for web servers";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-parts.url = "github:hercules-ci/flake-parts";
    templ.url = "github:a-h/templ";
  };
  outputs = inputs@{nixpkgs, flake-parts, templ, ...}: 
    flake-parts.lib.mkFlake { inherit inputs; }{
      systems = [
        "x86_64-linux"
        "x86_64-darwin"
        "aarch64-linux"
        "aarch64-darwin"
      ];
      perSystem = {pkgs, system, self', inputs', config, ... }: {
        devShells.default = pkgs.mkShell {
          nativeBuildInputs = with pkgs; [
            go_1_21 
            gopls 
            go-tools 
            gotools 
            hurl 
            (templ.packages.${system}.templ)
          ];
        };
      };
    };
}
