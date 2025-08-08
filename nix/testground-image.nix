{
  dockerTools,
  runCommandLocal,
  chainlet-matrix,
  benchmark-testcase,
}:
let
  patched-chainletd = chainlet-matrix.chainletd.overrideAttrs (oldAttrs: {
    patches = oldAttrs.patches or [ ] ++ [ ./testground-chainletd.patch ];
  });
in
let
  tmpDir = runCommandLocal "tmp" { } ''
    mkdir -p $out/tmp/
  '';
in
dockerTools.buildLayeredImage {
  name = "chainlet-testground";
  created = "now";
  contents = [
    benchmark-testcase
    patched-chainletd
    tmpDir
  ];
  config = {
    Expose = [
      9090
      26657
      26656
      1317
      26658
      26660
      26659
      30000
    ];
    Cmd = [ "/bin/stateless-testcase" ];
    Env = [ "PYTHONUNBUFFERED=1" ];
  };
}
