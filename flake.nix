{
  description = "Tidy Bookmark";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    nixpkgs,
    flake-utils,
    ...
  }:
    flake-utils.lib.eachDefaultSystem (
      system: let
        pkgs = nixpkgs.legacyPackages.${system};

        version = "0.0.0";

        devTools = with pkgs; [
          gh
          go
          gopls
          gotools
          golangci-lint
          lefthook
        ];

        commonBuildArgs = {
          pname = "tidy-bookmark";
          inherit version;
          src = ./.;
          vendorHash = "sha256-7K17JaXFsjf163g5PXCb5ng2gYdotnZ2IDKk8KFjNj0=";
          subPackages = ["."];
          doCheck = false;
          ldflags = [
            "-s"
            "-w"
          ];
          meta.mainProgram = "tidy-bookmark";
        };

        mkCheck = {
          name,
          nativeBuildInputs ? [],
          command,
        }:
          pkgs.buildGoModule (
            commonBuildArgs
            // {
              pname = "tidy-bookmark-${name}";
              inherit nativeBuildInputs;
              buildPhase = ''
                runHook preBuild
                export HOME="$TMPDIR"
                export GOCACHE="$TMPDIR/go-cache"
                export GOLANGCI_LINT_CACHE="$TMPDIR/golangci-lint"
                ${command}
                runHook postBuild
              '';
              installPhase = ''
                runHook preInstall
                mkdir -p "$out"
                runHook postInstall
              '';
            }
          );

        defaultPackage = pkgs.buildGoModule commonBuildArgs;

        buildReleaseArtifacts = pkgs.writeShellApplication {
          name = "build-release-artifacts";
          runtimeInputs = with pkgs; [
            coreutils
            go
            gnutar
            gzip
          ];
          text = ''
            output_dir="''${1:-dist}"
            work_dir="$(mktemp -d)"

            export HOME="$work_dir/home"
            export GOCACHE="$work_dir/go-cache"

            mkdir -p "$output_dir"
            mkdir -p "$HOME" "$GOCACHE"

            for target in "linux amd64" "linux arm64" "darwin amd64" "darwin arm64"; do
              read -r goos goarch <<<"$target"
              staging_dir="$(mktemp -d)"
              archive_path="$output_dir/tidy-bookmark_''${goos}_''${goarch}.tar.gz"

              GOOS="$goos" GOARCH="$goarch" CGO_ENABLED=0 \
                go build -trimpath -ldflags="-s -w" -o "$staging_dir/tidy-bookmark" .

              tar -C "$staging_dir" -czf "$archive_path" tidy-bookmark
              rm -rf "$staging_dir"
            done

            sha256sum "$output_dir"/tidy-bookmark_*.tar.gz > "$output_dir/tidy-bookmark_checksums.txt"
            chmod -R u+w "$work_dir" 2>/dev/null || true
            rm -rf "$work_dir" 2>/dev/null || true
          '';
        };
      in {
        devShells.default = pkgs.mkShell {packages = devTools;};

        packages = {
          default = defaultPackage;
          build-release-artifacts = buildReleaseArtifacts;
        };

        checks = {
          fmt = mkCheck {
            name = "fmt";
            nativeBuildInputs = [
              pkgs.findutils
              pkgs.gotools
            ];
            command = ''
              unformatted="$(
                find . -path ./vendor -prune -o -name '*.go' -print0 | xargs -0 gofmt -l
              )"
              if [ -n "$unformatted" ]; then
                echo "Files need gofmt:"
                echo "$unformatted"
                exit 1
              fi

              missingImports="$(
                find . -path ./vendor -prune -o -name '*.go' -print0 | xargs -0 goimports -l
              )"
              if [ -n "$missingImports" ]; then
                echo "Files need goimports:"
                echo "$missingImports"
                exit 1
              fi
            '';
          };

          lint = mkCheck {
            name = "lint";
            nativeBuildInputs = [pkgs.golangci-lint];
            command = ''
              golangci-lint run ./...
            '';
          };

          test = mkCheck {
            name = "test";
            command = ''
              go test ./...
            '';
          };

          build = defaultPackage;
        };
      }
    );
}
