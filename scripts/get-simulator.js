const child_process = require('child_process');
const process = require('process');

const TESTER_RELEASE_URL =
  "https://github.com/massalabs/massa-sc-tester/releases/download";
const VERSION = "v0.3.1";

let fileName;
switch (process.platform) {
  case "win32":
    // TODO
    fileName = "release_windows.zip";
    process.exit(0);
  case "darwin":
    fileName = "release_macos.tar.gz";
    break;
  case "linux":
    fileName = "release_linux.tar.gz";
    break;
  default:
    console.error(`OS not supported`);
    process.exit(1);
}

const url = `${TESTER_RELEASE_URL}/${VERSION}/${fileName}`;
const cmd = `wget -qO- ${url} | tar xz -C ./pkg --strip-components 1`;
child_process.execSync(cmd);

process.exit(0);
