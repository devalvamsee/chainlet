{
  pkgs ? import ../../nix { },
}:
let
  chainletd = (pkgs.callPackage ../../. { });
in
chainletd.overrideAttrs (oldAttrs: {
  patches = oldAttrs.patches or [ ] ++ [ ./broken-chainletd.patch ];
})
