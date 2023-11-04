{
  description = "Dev env for web servers";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-parts.url = "github:hercules-ci/flake-parts";
    templ.url = "github:a-h/templ";
  };

  outputs = inputs@{nixpkgs, flake-parts, templ, ...}:
    let
      buildGoProject = pkgs: with pkgs; buildGoModule {
        # Specify additional attributes for your Go project build
        pname = "your-project-name";
        version = "your-project-version";
        src = ./.;  # assuming your Go project's source code is in the same directory as this flake file
        vendorSha256 = null;  # set this to the correct hash for your project's vendor directory
        nativeBuildInputs = with pkgs; [
          go_1_21
          go-tools
        ];
      };
    in
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
              air
              hurl
              turso-cli
              vscode-langservers-extracted
              # (templ.packages.${system}.templ)
            ];
          };
        };
      };
}
