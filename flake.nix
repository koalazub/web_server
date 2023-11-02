{
  description = "Dev env for web servers";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-parts.url = "github:hercules-ci/flake-parts";
  };
  outputs = inputs@{nixpkgs, flake-parts, ...}: 
    flake-parts.lib.mkFlake { inherit inputs; }{
      systems = [
        "x86_64-linux"
        "x86_64-darwin"
        "aarch64-linux"
        "aarch64-darwin"
      ];
      perSystem = {pkgs, system, self', inputs', config, ... }: {
        devShells.default = pkgs.mkShell {
          nativeBuildInputs = with pkgs; [go_1_21 gopls go-tools gotools hurl];
        };
      };
    };
}
