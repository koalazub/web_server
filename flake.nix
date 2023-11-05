{
  description = "Dev env for web servers";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-parts.url = "github:hercules-ci/flake-parts";
    templ.url = "github:a-h/templ";
  };

  outputs = inputs@{nixpkgs, flake-parts, templ, ...}: flake-parts.lib.mkFlake { inherit inputs; }{
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

          packages.staging = pkgs.buildGoModule rec {
            pname = "web_server";
            version = "0.5.0";
            src = ./.;
            vendorSha256 = null;
            fullName = "${pname}-${version}";

            nativeBuildInputs = with pkgs; [  
              go_1_21
              go-tools
            ];  
            preBuild = ''
              go mod vendor
              go mod tidy 
            '';
            testPhase = ''
              go test ./...
            '';
            buildPhase = ''
              go build -o results ./ 
            ''; 
            
          };
        };
      };
}
